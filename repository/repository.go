package repository

import (
	"fmt"
	"os/exec"
	"strings"
)

func FindRoot() (string, error) {
	path, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		return "", fmt.Errorf("failed running git: %w", err)
	}

	return strings.TrimSpace(string(path)), nil
}

func GetRevIDShort() (string, error) {
	result, err := exec.Command("git", "rev-parse", "--short", "HEAD").Output()
	if err != nil {
		return "", fmt.Errorf("failed running git: %w", err)
	}

	return strings.TrimSpace(string(result)), nil
}

func GetRevID() (string, error) {
	result, err := exec.Command("git", "rev-parse", "HEAD").Output()
	if err != nil {
		return "", fmt.Errorf("failed running git: %w", err)
	}

	return strings.TrimSpace(string(result)), nil
}

func GetBranch() (string, error) {
	result, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		return "", fmt.Errorf("failed running git: %w", err)
	}

	return strings.TrimSpace(string(result)), nil
}
