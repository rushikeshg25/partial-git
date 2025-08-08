package repository

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"sync"
	"time"
)

type GitHubURL struct {
	Owner      string
	Repository string
	Path       string
	Branch     string
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

// GetAPIURL returns the GitHub API URL for this repository
func (g *GitHubURL) GetAPIURL() string {
	return fmt.Sprintf("https://api.github.com/repos/%s/%s", g.Owner, g.Repository)
}

func (g *GitHubURL) Download(ctx context.Context) error {
	wg := sync.WaitGroup{}
	errChan := make(chan error, 1)
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	wg.Add(1)
	go g.getRepoContent(ctx, wg, errChan)

	go func() {
		defer func() {
			wg.Wait()
			close(errChan)
		}()
	}()

	select {
	case err := <-errChan:
		if err != nil {
			return fmt.Errorf("failed to download: %w", err)
		}
	case <-ctx.Done():
		if errors.Is(ctx.Err(), context.Canceled) {
			return context.Canceled
		}
		return errors.New("Timeout")
	}

	return nil
}

func (g *GitHubURL) getRepoContent(ctx *context.Context, wg *sync.WaitGroup, erchan chan error) error {
}
