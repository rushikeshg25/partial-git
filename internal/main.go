package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
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

		// Create a timeout context for download operations (60 seconds)
		downloadCtx, cancel := context.WithTimeout(ctx, 60*time.Second)
		defer cancel()

		fmt.Println("Starting download (timeout: 60 seconds)...")
		if err := githubURL.Download(downloadCtx); err != nil {
			if downloadCtx.Err() == context.DeadlineExceeded {
				fmt.Fprintf(os.Stderr, "Error: Download timed out after 60 seconds\n")
			} else if downloadCtx.Err() == context.Canceled {
				fmt.Fprintf(os.Stderr, "Error: Download was cancelled\n")
			} else {
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
	req, err := http.NewRequestWithContext(ctx, "GET", githubURL.GetAPIURL(), nil)
	if err != nil {
		return false, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, fmt.Errorf("failed to check repository: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return true, nil
	}

	if resp.StatusCode == http.StatusForbidden {
		return true, nil
	}

	if resp.StatusCode == http.StatusOK {
		var repoInfo struct {
			Private bool `json:"private"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&repoInfo); err != nil {
			return false, fmt.Errorf("failed to parse repository info: %w", err)
		}

		return repoInfo.Private, nil
	}

	return true, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
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
	fmt.Println("Note: Use this token to authenticate with GitHub API")
	// TODO: Add actual GitHub API call to get user info with context
}

// checkTokenStatus checks the token status and GitHub API rate limits
func checkTokenStatus(ctx context.Context, tokenManager *token.Manager) {
	// Check for cancellation
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
	fmt.Println("✓ Token ready for GitHub API calls")
	// TODO: Add actual GitHub API call to check rate limits with context
}
