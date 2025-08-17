package token

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const envVarName = "PGIT_GITHUB_TOKEN"

type EnvironmentStorage struct{}

func (e *EnvironmentStorage) Set(token string) error {
	if token == "" {
		return fmt.Errorf("token cannot be empty")
	}

	if err := os.Setenv(envVarName, token); err != nil {
		return fmt.Errorf("failed to set environment variable: %w", err)
	}

	if err := e.addToShellProfile(token); err != nil {
		fmt.Printf("⚠️  Warning: Failed to add to shell profile: %v\n", err)
		fmt.Printf("✓ Token set for current session only.\n")
		fmt.Printf("\n To make it permanent, manually add this line to your shell profile:\n")
		fmt.Printf("   export %s=%s\n", envVarName, token)
		fmt.Printf("\n Then restart your terminal or run: source ~/.zshrc\n")
		return nil
	}

	fmt.Printf("✓ Token stored in shell profile and current session\n")
	fmt.Printf("✓ Token will persist across terminal sessions\n")
	fmt.Printf("\n IMPORTANT: To use the token in new terminal sessions:\n")
	fmt.Printf("   • Open a new terminal window/tab, OR\n")
	fmt.Printf("   • Run: source ~/.zshrc (or source ~/.bashrc)\n")
	fmt.Printf("\n The token is available in this current session immediately.\n")

	return nil
}

func (e *EnvironmentStorage) Get() (string, error) {
	token := os.Getenv(envVarName)
	if token == "" {
		return "", fmt.Errorf("GitHub token not found in environment variable %s", envVarName)
	}

	return token, nil
}

func (e *EnvironmentStorage) Delete() error {
	if err := os.Unsetenv(envVarName); err != nil {
		return fmt.Errorf("failed to unset environment variable: %w", err)
	}

	if err := e.removeFromShellProfile(); err != nil {
		fmt.Printf("⚠️  Warning: Failed to remove from shell profile: %v\n", err)
		fmt.Printf("✓ Token removed from current session only.\n")
		fmt.Printf("\nYou may need to manually remove this line from your shell profile:\n")
		fmt.Printf("   export %s=...\n", envVarName)
		fmt.Printf("\nThen restart your terminal or run: source ~/.zshrc\n")
		return nil
	}

	fmt.Printf("✓ Token removed from shell profile and current session\n")
	fmt.Printf("\nIMPORTANT: To apply changes in new terminal sessions:\n")
	fmt.Printf("   • Open a new terminal window/tab, OR\n")
	fmt.Printf("   • Run: source ~/.zshrc (or source ~/.bashrc)\n")
	return nil
}

func (e *EnvironmentStorage) Exists() bool {
	return os.Getenv(envVarName) != ""
}

func (e *EnvironmentStorage) getShellProfilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user home directory: %w", err)
	}

	shell := os.Getenv("SHELL")

	profiles := []string{}

	if strings.Contains(shell, "zsh") {
		profiles = []string{".zshrc", ".zprofile"}
	} else if strings.Contains(shell, "bash") {
		profiles = []string{".bashrc", ".bash_profile", ".profile"}
	} else {
		profiles = []string{".zshrc", ".bashrc", ".profile"}
	}

	for _, profile := range profiles {
		profilePath := filepath.Join(homeDir, profile)
		if _, err := os.Stat(profilePath); err == nil {
			return profilePath, nil
		}
	}

	var defaultProfile string
	if strings.Contains(shell, "zsh") {
		defaultProfile = ".zshrc"
	} else if strings.Contains(shell, "bash") {
		defaultProfile = ".bashrc"
	} else {
		defaultProfile = ".zshrc"
	}

	return filepath.Join(homeDir, defaultProfile), nil
}

func (e *EnvironmentStorage) addToShellProfile(token string) error {
	profilePath, err := e.getShellProfilePath()
	if err != nil {
		return err
	}

	if e.tokenExistsInProfile(profilePath) {
		if err := e.removeFromShellProfile(); err != nil {
			return fmt.Errorf("failed to remove existing token: %w", err)
		}
	}

	exportLine := fmt.Sprintf("export %s=%s", envVarName, token)

	file, err := os.OpenFile(profilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open shell profile: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("\n# PGIT GitHub token\n%s\n", exportLine))
	if err != nil {
		return fmt.Errorf("failed to write to shell profile: %w", err)
	}

	return nil
}

func (e *EnvironmentStorage) removeFromShellProfile() error {
	profilePath, err := e.getShellProfilePath()
	if err != nil {
		return err
	}

	if _, err := os.Stat(profilePath); os.IsNotExist(err) {
		return nil
	}

	file, err := os.Open(profilePath)
	if err != nil {
		return fmt.Errorf("failed to open shell profile: %w", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	skipNext := false

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "# PGIT GitHub token") {
			skipNext = true
			continue
		}

		if skipNext && strings.HasPrefix(strings.TrimSpace(line), fmt.Sprintf("export %s=", envVarName)) {
			skipNext = false
			continue
		}

		if strings.HasPrefix(strings.TrimSpace(line), fmt.Sprintf("export %s=", envVarName)) {
			continue
		}

		skipNext = false
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("failed to read shell profile: %w", err)
	}

	err = os.WriteFile(profilePath, []byte(strings.Join(lines, "\n")+"\n"), 0644)
	if err != nil {
		return fmt.Errorf("failed to write shell profile: %w", err)
	}

	return nil
}

func (e *EnvironmentStorage) tokenExistsInProfile(profilePath string) bool {
	file, err := os.Open(profilePath)
	if err != nil {
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, fmt.Sprintf("export %s=", envVarName)) {
			return true
		}
	}

	return false
}
