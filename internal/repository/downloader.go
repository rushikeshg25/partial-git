package repository

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/google/go-github/v57/github"
)

var (
	ErrTookTooLong = errors.New("download took too long")
)

type Downloader struct {
	client          *GitHubClient
	httpClient      *http.Client
	owner           string
	repo            string
	basePath        string
	branch          string
	quiet           bool
	downloadedCount int
	totalCount      int
	mu              sync.Mutex
}

func NewDownloader(client *GitHubClient, owner, repo, basePath, branch string) *Downloader {
	return NewDownloaderWithOptions(client, owner, repo, basePath, branch, false)
}

func NewDownloaderWithOptions(client *GitHubClient, owner, repo, basePath, branch string, quiet bool) *Downloader {
	httpClient := &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 10,
			IdleConnTimeout:     90 * time.Second,
		},
	}

	return &Downloader{
		client:     client,
		httpClient: httpClient,
		owner:      owner,
		repo:       repo,
		basePath:   basePath,
		branch:     branch,
	}
}

func (d *Downloader) Download(ctx context.Context) error {
	wg := &sync.WaitGroup{}
	errCh := make(chan error, 1)

	timeoutCtx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	wg.Add(1)
	go d.downloadContents(timeoutCtx, wg, d.basePath, errCh)

	go func() {
		wg.Wait()
		close(errCh)
	}()

	select {
	case err := <-errCh:
		return err
	case <-timeoutCtx.Done():
		if timeoutCtx.Err() == context.DeadlineExceeded {
			return ErrTookTooLong
		}
		return timeoutCtx.Err()
	}
}

func (d *Downloader) downloadContents(ctx context.Context, wg *sync.WaitGroup, path string, errCh chan error) {
	defer wg.Done()

	select {
	case <-ctx.Done():
		d.sendError(errCh, ctx.Err())
		return
	default:
	}

	opts := &github.RepositoryContentGetOptions{}
	if d.branch != "" {
		opts.Ref = d.branch
	}

	fileContent, directoryContent, err := d.client.GetContents(ctx, d.owner, d.repo, path, opts)
	if err != nil {
		d.sendError(errCh, fmt.Errorf("failed to get contents for path '%s': %w", path, err))
		return
	}

	if fileContent != nil {
		if err := d.downloadFile(ctx, fileContent.GetDownloadURL(), fileContent.GetPath()); err != nil {
			d.sendError(errCh, err)
		}
		return
	}

	if directoryContent != nil {
		d.processDirectoryContents(ctx, wg, directoryContent, errCh)
	}
}

func (d *Downloader) processDirectoryContents(ctx context.Context, wg *sync.WaitGroup, contents []*github.RepositoryContent, errCh chan error) {
	for _, content := range contents {
		select {
		case <-ctx.Done():
			d.sendError(errCh, ctx.Err())
			return
		default:
		}

		wg.Add(1)
		go func(content *github.RepositoryContent) {
			defer wg.Done()

			switch content.GetType() {
			case "file":
				if err := d.downloadFile(ctx, content.GetDownloadURL(), content.GetPath()); err != nil {
					d.sendError(errCh, err)
				}
			case "dir":
				wg.Add(1)
				go d.downloadContents(ctx, wg, content.GetPath(), errCh)
			}
		}(content)
	}
}

func (d *Downloader) downloadFile(ctx context.Context, downloadURL, path string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	fmt.Printf("Downloading: %s\n", path)

	req, err := http.NewRequestWithContext(ctx, "GET", downloadURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create request for %s: %w", path, err)
	}

	resp, err := d.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to download %s: %w", path, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download %s: received HTTP %d %s", path, resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	localPath, err := d.getExactPath(d.basePath, path)
	if err != nil {
		return fmt.Errorf("failed to determine local path for %s: %w", path, err)
	}

	if err := os.MkdirAll(filepath.Dir(localPath), 0755); err != nil {
		return fmt.Errorf("failed to create directory structure for %s: %w", localPath, err)
	}

	file, err := os.OpenFile(localPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", localPath, err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write file %s: %w", localPath, err)
	}

	return nil
}

func (d *Downloader) getExactPath(base, path string) (string, error) {
	if base == "" {
		return filepath.Join(d.repo, path), nil
	}

	relPath, err := filepath.Rel(base, path)
	if err != nil {
		return "", err
	}

	return filepath.Join(filepath.Base(base), relPath), nil
}

func (d *Downloader) sendError(errCh chan error, err error) {
	select {
	case errCh <- err:
	default:
	}
}
