package repository

import (
	"context"
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/google/go-github/v57/github"
)

type GitHubURL struct {
	Owner      string
	Repository string
	Path       string // path within the repository
	Branch     string // branch/ref
	RawURL     string
}

func ParseGitHubURL(urlStr string) (*GitHubURL, error) {
	if urlStr == "" {
		return nil, fmt.Errorf("URL cannot be empty")
	}

	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return nil, fmt.Errorf("invalid URL format: %w", err)
	}

	if parsedURL.Host != "github.com" && parsedURL.Host != "www.github.com" {
		return nil, fmt.Errorf("URL must be from github.com")
	}

	if parsedURL.Scheme == "" {
		return nil, fmt.Errorf("URL must include scheme (https://)")
	}

	if parsedURL.Scheme != "https" && parsedURL.Scheme != "http" {
		return nil, fmt.Errorf("URL scheme must be http or https")
	}

	pathParts := strings.Split(strings.Trim(parsedURL.Path, "/"), "/")
	if len(pathParts) < 2 {
		return nil, fmt.Errorf("GitHub URL must include owner and repository (e.g., https://github.com/owner/repo)")
	}

	if pathParts[0] == "" || pathParts[1] == "" {
		return nil, fmt.Errorf("GitHub URL must include valid owner and repository names")
	}

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

func (g *GitHubURL) String() string {
	return fmt.Sprintf("%s/%s", g.Owner, g.Repository)
}

func (g *GitHubURL) GetAPIURL() string {
	return fmt.Sprintf("https://api.github.com/repos/%s/%s", g.Owner, g.Repository)
}

func (g *GitHubURL) GetCloneURL() string {
	return fmt.Sprintf("https://github.com/%s/%s.git", g.Owner, g.Repository)
}

func (g *GitHubURL) Download(ctx context.Context) error {
	return g.DownloadWithOptions(ctx, false)
}

func (g *GitHubURL) DownloadWithOptions(ctx context.Context, quiet bool) error {
	client := NewGitHubClient()
	downloader := NewDownloaderWithOptions(client, g.Owner, g.Repository, g.Path, g.Branch, quiet)
	return downloader.Download(ctx)
}

func (g *GitHubURL) IsPrivate(ctx context.Context) (bool, error) {
	client := NewGitHubClient()
	repo, err := client.GetRepository(ctx, g.Owner, g.Repository)
	if err != nil {
		// assume it might be private or doesn't exist
		return true, err
	}

	return repo.GetPrivate(), nil
}

func (g *GitHubURL) GetRateLimit(ctx context.Context) (*github.RateLimits, error) {
	client := NewGitHubClient()
	return client.GetRateLimit(ctx)
}

func (g *GitHubURL) GetAuthenticatedUser(ctx context.Context) (*github.User, error) {
	client := NewGitHubClient()
	return client.GetAuthenticatedUser(ctx)
}
