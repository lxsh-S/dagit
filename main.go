package main

import (
	"fmt"

	"github.com/lxsh-S/dagit/internal/git"
)

func main() {
	if git.IsGitRepo() {
		// We'll begin the TUI here
		fmt.Println("Yes we are in a git repo")
		repoName, err := git.GetRepoName()
		if err == nil {
			fmt.Println(repoName)
		}
		repoURL, err := git.GetRemoteURL()
		if err == nil {
			fmt.Println(repoURL)
		}

		repoOwner, err := git.GetRepoOwner()
		if err == nil {
			fmt.Println(repoOwner)
		}

		crBranch, err := git.GetCurrentBranch()
		if err == nil {
			fmt.Println(crBranch)
		}
	} else {
		// If its not a git repo then we dont open our TUI
		fmt.Println("err: this isn't a git repo")
	}
}
