package repository

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

// GitHubURL represents a parsed GitHub repository URL
type GitHubURL struct {
	Owner      string
	Repository string
	Path       string // Optional path within the repository
	Branch     string // Optional branch/ref
	RawURL     string // Original URL string
}

// ParseGitHubURL parses and validates a GitHub URL
func ParseGitHubURL(urlStr string) (*GitHubURL, error) {
	if urlStr == "" {
		return nil, fmt.Errorf("URL cannot be empty")
	}

	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return nil, fmt.Errorf("invalid URL format: %w", err)
	}

	// Validate GitHub domain
	if parsedURL.Host != "github.com" && parsedURL.Host != "www.github.com" {
		return nil, fmt.Errorf("URL must be from github.com")
	}

	// Validate scheme
	if parsedURL.Scheme == "" {
		return nil, fmt.Errorf("URL must include scheme (https://)")
	}

	if parsedURL.Scheme != "https" && parsedURL.Scheme != "http" {
		return nil, fmt.Errorf("URL scheme must be http or https")
	}

	// Parse path components
	pathParts := strings.Split(strings.Trim(parsedURL.Path, "/"), "/")
	if len(pathParts) < 2 {
		return nil, fmt.Errorf("GitHub URL must include owner and repository (e.g., https://github.com/owner/repo)")
	}

	// Validate owner and repository names
	if pathParts[0] == "" || pathParts[1] == "" {
		return nil, fmt.Errorf("GitHub URL must include valid owner and repository names")
	}

	// GitHub username/org and repo name validation
	nameRegex := regexp.MustCompile(`^[a-zA-Z0-9]([a-zA-Z0-9\-_])*[a-zA-Z0-9]$|^[a-zA-Z0-9]$`)
	if !nameRegex.MatchString(pathParts[0]) {
		return nil, fmt.Errorf("invalid GitHub owner name: %s", pathParts[0])
	}
	if !nameRegex.MatchString(pathParts[1]) {
		return nil, fmt.Errorf("invalid GitHub repository name: %s", pathParts[1])
	}

	githubURL := &GitHubURL{
		Owner:      pathParts[0],
		Repository: pathParts[1],
		RawURL:     urlStr,
	}

	// Parse additional path components (for file/folder paths)
	if len(pathParts) > 2 {
		// Handle different GitHub URL patterns:
		// /owner/repo/tree/branch/path/to/file
		// /owner/repo/blob/branch/path/to/file
		// /owner/repo/path/to/file (direct path)

		if len(pathParts) > 3 && (pathParts[2] == "tree" || pathParts[2] == "blob") {
			// URL has branch/ref specified
			githubURL.Branch = pathParts[3]
			if len(pathParts) > 4 {
				githubURL.Path = strings.Join(pathParts[4:], "/")
			}
		} else {
			// Direct path without branch specification
			githubURL.Path = strings.Join(pathParts[2:], "/")
		}
	}

	return githubURL, nil
}

// String returns a string representation of the GitHub URL
func (g *GitHubURL) String() string {
	return fmt.Sprintf("%s/%s", g.Owner, g.Repository)
}

// GetAPIURL returns the GitHub API URL for this repository
func (g *GitHubURL) GetAPIURL() string {
	return fmt.Sprintf("https://api.github.com/repos/%s/%s", g.Owner, g.Repository)
}

// GetCloneURL returns the HTTPS clone URL for this repository
func (g *GitHubURL) GetCloneURL() string {
	return fmt.Sprintf("https://github.com/%s/%s.git", g.Owner, g.Repository)
}

func (g *GitHubURL) Download() error {
	return nil
}
