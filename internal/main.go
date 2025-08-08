package internal

import (
	"fmt"
	"os"
	"partial-git/internal/token"
)

type Flags struct {
	Set   string
	Auth  bool
	Check bool
	Unset bool
}

func Run(flags Flags, args []string) {
	tokenManager := token.NewManager()

	if flags.Set != "" {
		if err := tokenManager.SetToken(flags.Set); err != nil {
			fmt.Fprintf(os.Stderr, "Error setting GitHub token: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Storage backend: %s\n", tokenManager.GetStorageInfo())
		return
	}

	if flags.Unset {
		if err := tokenManager.DeleteToken(); err != nil {
			fmt.Fprintf(os.Stderr, "Error deleting GitHub token: %v\n", err)
			os.Exit(1)
		}
		return
	}

	// Perform runtime validations for other operations
	if err := validateRuntimeConditions(flags, tokenManager); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Handle auth and check operations
	if flags.Auth {
		showAuthInfo(tokenManager)
		return
	}

	if flags.Check {
		checkTokenStatus(tokenManager)
		return
	}

	// Default behavior - show input summary (for download operations)
	fmt.Println("=== User Input Summary ===")

	// Print command line arguments
	if len(args) > 0 {
		fmt.Printf("Arguments: %v\n", args)
	} else {
		fmt.Println("Arguments: none")
	}

	fmt.Printf("Set flag: %q\n", flags.Set)
	fmt.Printf("Auth flag: %t\n", flags.Auth)
	fmt.Printf("Check flag: %t\n", flags.Check)
	fmt.Printf("Unset flag: %t\n", flags.Unset)

	fmt.Println("========================")
}

// validateRuntimeConditions performs additional runtime validations
func validateRuntimeConditions(flags Flags, tokenManager *token.Manager) error {
	// Check if GitHub token exists for operations that require it
	if flags.Auth || flags.Check {
		if !tokenManager.TokenExists() {
			return fmt.Errorf("GitHub token not found. Use --set to configure it first")
		}
	}

	return nil
}

// showAuthInfo displays information about the authenticated user
func showAuthInfo(tokenManager *token.Manager) {
	storedToken, err := tokenManager.GetToken()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error retrieving token: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("GitHub token found (storage: %s)\n", tokenManager.GetStorageInfo())
	fmt.Printf("Token prefix: %s***\n", storedToken[:8])
	fmt.Println("Note: Use this token to authenticate with GitHub API")
	// TODO: Add actual GitHub API call to get user info
}

// checkTokenStatus checks the token status and GitHub API rate limits
func checkTokenStatus(tokenManager *token.Manager) {
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
	// TODO: Add actual GitHub API call to check rate limits
}
