package repository

import (
	"context"
	"os"

	"partial-git/internal/token"

	"github.com/google/go-github/v57/github"
	"golang.org/x/oauth2"
)

type GitHubClient struct {
	client *github.Client
}

func NewGitHubClient() *GitHubClient {
	var client *github.Client

	authToken := os.Getenv("GITHUB_TOKEN")

	if authToken == "" {
		tokenManager := token.NewManager()
		if tokenManager.TokenExists() {
			if t, err := tokenManager.GetToken(); err == nil {
				authToken = t
			}
		}
	}

	if authToken != "" {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: authToken},
		)
		tc := oauth2.NewClient(context.Background(), ts)
		client = github.NewClient(tc)
	} else {
		client = github.NewClient(nil)
	}

	return &GitHubClient{client: client}
}

func (gc *GitHubClient) GetContents(ctx context.Context, owner, repo, path string, opts *github.RepositoryContentGetOptions) (*github.RepositoryContent, []*github.RepositoryContent, error) {
	fileContent, directoryContent, _, err := gc.client.Repositories.GetContents(ctx, owner, repo, path, opts)
	return fileContent, directoryContent, err
}

func (gc *GitHubClient) GetRepository(ctx context.Context, owner, repo string) (*github.Repository, error) {
	repository, _, err := gc.client.Repositories.Get(ctx, owner, repo)
	return repository, err
}

func (gc *GitHubClient) GetRateLimit(ctx context.Context) (*github.RateLimits, error) {
	rateLimits, _, err := gc.client.RateLimit.Get(ctx)
	return rateLimits, err
}

func (gc *GitHubClient) GetAuthenticatedUser(ctx context.Context) (*github.User, error) {
	user, _, err := gc.client.Users.Get(ctx, "")
	return user, err
}
