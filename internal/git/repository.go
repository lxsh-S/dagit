package git

import (
	"os/exec"
	"path/filepath"
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
	return url, nil
}
