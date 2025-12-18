// Package github provides a client to interact with the GitHub API for fetching repository contents.
package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const baseAPIURL = "https://api.github.com"

// Client represents a GitHub API client.
type Client struct {
	httpClient *http.Client
	repo       string // format: "owner/repo"
}

// NewClient creates a new GitHub client for the specified repository.
func NewClient(repo string) *Client {
	return &Client{
		httpClient: &http.Client{},
		repo:       repo,
	}
}

// ContentItem represents a file or directory in a GitHub repository.
type ContentItem struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Type        string `json:"type"` // "file" or "dir"
	DownloadURL string `json:"download_url,omitempty"`
}

// ListContents lists the contents of a directory in the repository.
func (c *Client) ListContents(path string) ([]ContentItem, error) {
	url := fmt.Sprintf("%s/repos/%s/contents/%s", baseAPIURL, c.repo, path)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch contents: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("GitHub API error (status %d): %s", resp.StatusCode, string(body))
	}

	var items []ContentItem
	if err := json.NewDecoder(resp.Body).Decode(&items); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return items, nil
}

// DownloadFile downloads the raw content of a file from its download URL.
func (c *Client) DownloadFile(downloadURL string) ([]byte, error) {
	resp, err := c.httpClient.Get(downloadURL)
	if err != nil {
		return nil, fmt.Errorf("failed to download file: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to download file (status %d)", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}
