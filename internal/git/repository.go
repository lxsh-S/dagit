package git

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

func IsGitRepo() bool {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	return cmd.Run() == nil
}

func GetRepoName() (string, error) {
	out, err := exec.Command("git", "config", "--get", "remote.origin.url").Output()
	if err != nil {
		return "Failed!", nil
	}
	// trimout
	url := strings.TrimSpace(string(out))

	// Remove tht .git thing
	url = strings.TrimSuffix(url, ".git")

	// return the name and error *if any
	return filepath.Base(url), nil
}

func GetRemoteURL() (string, error) {
	out, err := exec.Command("git", "config", "--get", "remote.origin.url").Output()
	if err != nil {
		return "Failed!", nil
	}
	// trimout here too
	url := strings.TrimSpace(string(out))

	// Remove the .git thing because it makes the url look ugly :P
	url = strings.TrimSuffix(url, ".git")
	return url, err
}

func GetRepoOwner() (string, error) {
	out, err := exec.Command("git", "remote", "get-url", "origin").Output()
	if err != nil {
		return "", err
	}

	url := strings.TrimSpace(string(out))

	// For both SSH and HTTP

	re := regexp.MustCompile(`[:/]([^/]+)/([^/]+?)(\.git)?$`)
	matches := re.FindStringSubmatch(url)
	if len(matches) < 3 {
		return "", fmt.Errorf("could not get remote URL: %s", url)
	}

	owner, _ := matches[1], matches[2]
	return owner, nil
}

func GetCurrentBranch() (string, error) {
	out, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		return "", err
	}
	branch := strings.TrimSpace(string(out))
	return branch, nil
}
