package token

import (
	"fmt"
	"regexp"
	"strings"
)

type Manager struct {
	storage *EnvironmentStorage
}

func NewManager() *Manager {
	return &Manager{storage: &EnvironmentStorage{}}
}

func (m *Manager) SetToken(token string) error {
	if err := m.validateToken(token); err != nil {
		return fmt.Errorf("invalid token: %w", err)
	}

	if err := m.storage.Set(token); err != nil {
		return fmt.Errorf("failed to store token: %w", err)
	}

	return nil
}

func (m *Manager) GetToken() (string, error) {
	token, err := m.storage.Get()
	if err != nil {
		return "", fmt.Errorf("failed to retrieve token: %w", err)
	}

	return token, nil
}

func (m *Manager) DeleteToken() error {
	if !m.storage.Exists() {
		fmt.Println("No GitHub token found to delete")
		return nil
	}

	if err := m.storage.Delete(); err != nil {
		return fmt.Errorf("failed to delete token: %w", err)
	}

	return nil
}

func (m *Manager) TokenExists() bool {
	return m.storage.Exists()
}

func (m *Manager) validateToken(token string) error {
	if token == "" {
		return fmt.Errorf("token cannot be empty")
	}

	if len(token) < 40 {
		return fmt.Errorf("token appears to be too short (minimum 40 characters)")
	}

	tokenRegex := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	if !tokenRegex.MatchString(token) {
		return fmt.Errorf("token contains invalid characters (only alphanumeric and underscore allowed)")
	}

	validPrefixes := []string{
		"ghp_",        // Personal access token
		"gho_",        // OAuth token
		"ghu_",        // User-to-server token
		"ghs_",        // Server-to-server token
		"ghr_",        // Refresh token
		"github_pat_", // Fine-grained personal access token
	}

	hasValidPrefix := false
	for _, prefix := range validPrefixes {
		if strings.HasPrefix(token, prefix) {
			hasValidPrefix = true
			break
		}
	}

	if !hasValidPrefix {
		return fmt.Errorf("token does not have a valid GitHub token prefix (ghp_, gho_, ghu_, ghs_, ghr_, github_pat_)")
	}

	return nil
}

func (m *Manager) GetStorageInfo() string {
	return "Shell Profile (zsh/bash)"
}

// ValidateToken validates a GitHub PAT without storing it
func ValidateToken(token string) error {
	manager := NewManager()
	return manager.validateToken(token)
}
