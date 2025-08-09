package internal

import (
	"context"
	"fmt"
	"os"
	"partial-git/internal/repository"
	"partial-git/internal/token"
	"time"
)

type Flags struct {
	Set   string
	Auth  bool
	Check bool
	Unset bool
}

func Run(ctx context.Context, flags Flags, args []string) {
	tokenManager := token.NewManager()

	switch {
	case flags.Set != "":
		if err := tokenManager.SetToken(flags.Set); err != nil {
			fmt.Fprintf(os.Stderr, "Error setting GitHub token: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Storage backend: %s\n", tokenManager.GetStorageInfo())
		return

	case flags.Unset:
		if err := tokenManager.DeleteToken(); err != nil {
			fmt.Fprintf(os.Stderr, "Error deleting GitHub token: %v\n", err)
			os.Exit(1)
		}
		return

	case flags.Auth:
		if err := validateRuntimeConditions(ctx, flags, tokenManager, nil); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		showAuthInfo(ctx, tokenManager)
		return

	case flags.Check:
		if err := validateRuntimeConditions(ctx, flags, tokenManager, nil); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		checkTokenStatus(ctx, tokenManager)
		return

	default:
		if len(args) == 0 {
			fmt.Fprintf(os.Stderr, "Error: GitHub URL is required\n")
			fmt.Fprintf(os.Stderr, "Usage: pgit <github-url>\n")
			fmt.Fprintf(os.Stderr, "   or: pgit --set <token>\n")
			fmt.Fprintf(os.Stderr, "   or: pgit --check\n")
			fmt.Fprintf(os.Stderr, "   or: pgit --auth\n")
			fmt.Fprintf(os.Stderr, "   or: pgit --unset\n")
			os.Exit(1)
		}

		start := time.Now()
		githubURL, err := parseGitHubURL(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing GitHub URL: %v\n", err)
			os.Exit(1)
		}

		if err := validateRuntimeConditions(ctx, flags, tokenManager, githubURL); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Starting download of %s...\n", githubURL.String())
		if githubURL.Path != "" {
			fmt.Printf("Path: %s\n", githubURL.Path)
		}
		if githubURL.Branch != "" {
			fmt.Printf("Branch: %s\n", githubURL.Branch)
		}

		if err := githubURL.Download(ctx); err != nil {
			switch err {
			case repository.ErrTookTooLong:
				fmt.Fprintf(os.Stderr, "Error: Download timed out after 60 seconds\n")
			case context.Canceled:
				fmt.Fprintf(os.Stderr, "Error: Download was cancelled\n")
			default:
				fmt.Fprintf(os.Stderr, "Error downloading repository: %v\n", err)
			}
			os.Exit(1)
		}
		fmt.Println("Download Completed")
		fmt.Println(time.Since(start))
	}
}

func parseGitHubURL(urlStr string) (*repository.GitHubURL, error) {
	githubURL, err := repository.ParseGitHubURL(urlStr)
	if err != nil {
		return nil, fmt.Errorf("invalid GitHub URL '%s': %w", urlStr, err)
	}
	return githubURL, nil
}

func validateRuntimeConditions(ctx context.Context, flags Flags, tokenManager *token.Manager, githubURL *repository.GitHubURL) error {
	if flags.Auth || flags.Check {
		if !tokenManager.TokenExists() {
			return fmt.Errorf("GitHub token not found. Use --set to configure it first")
		}
		return nil
	}

	if githubURL != nil {
		isPrivate, err := isRepositoryPrivate(ctx, githubURL)
		if err != nil {
			fmt.Printf("Warning: Could not determine repository visibility: %v\n", err)
			fmt.Println("Assuming repository might be private...")
			if !tokenManager.TokenExists() {
				return fmt.Errorf("GitHub token not found. Repository might be private. Use --set to configure a token")
			}
		} else if isPrivate {
			if !tokenManager.TokenExists() {
				return fmt.Errorf("GitHub token required for private repository. Use --set to configure a token")
			}
		}
	}

	return nil
}

func isRepositoryPrivate(ctx context.Context, githubURL *repository.GitHubURL) (bool, error) {
	return githubURL.IsPrivate(ctx)
}

func showAuthInfo(ctx context.Context, tokenManager *token.Manager) {
	select {
	case <-ctx.Done():
		fmt.Println("Operation cancelled")
		return
	default:
	}

	storedToken, err := tokenManager.GetToken()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error retrieving token: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("GitHub token found (storage: %s)\n", tokenManager.GetStorageInfo())
	fmt.Printf("Token prefix: %s***\n", storedToken[:8])

	client := repository.NewGitHubClient()

	user, err := client.GetAuthenticatedUser(ctx)
	if err != nil {
		fmt.Printf("Warning: Could not get user info: %v\n", err)
		fmt.Println("Token appears to be valid but user info unavailable")
		return
	}

	fmt.Printf("✓ Authenticated as: %s\n", user.GetLogin())
	if user.GetName() != "" {
		fmt.Printf("✓ Name: %s\n", user.GetName())
	}
	if user.GetEmail() != "" {
		fmt.Printf("✓ Email: %s\n", user.GetEmail())
	}
}

func checkTokenStatus(ctx context.Context, tokenManager *token.Manager) {
	select {
	case <-ctx.Done():
		fmt.Println("Operation cancelled")
		return
	default:
	}

	if !tokenManager.TokenExists() {
		fmt.Println("No GitHub token configured")
		fmt.Println("Use --set to configure a Personal Access Token")
		return
	}

	storedToken, err := tokenManager.GetToken()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error retrieving token: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✓ GitHub token found (storage: %s)\n", tokenManager.GetStorageInfo())
	fmt.Printf("✓ Token format valid (prefix: %s***)\n", storedToken[:8])

	client := repository.NewGitHubClient()

	rateLimits, err := client.GetRateLimit(ctx)
	if err != nil {
		fmt.Printf("Warning: Could not get rate limit info: %v\n", err)
		fmt.Println("✓ Token ready for GitHub API calls")
		return
	}

	core := rateLimits.GetCore()
	fmt.Printf("✓ Token ready for GitHub API calls\n")
	fmt.Printf("✓ Rate limit: %d/%d (resets at %v)\n",
		core.Remaining,
		core.Limit,
		core.Reset.Time.Format("15:04:05"))

	if core.Remaining < 100 {
		fmt.Printf("⚠️  Warning: Low rate limit remaining (%d requests)\n", core.Remaining)
	}
}
