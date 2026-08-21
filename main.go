package main

import (
	"fmt"

	"github.com/lxsh-S/dagit/TUI"
	"github.com/lxsh-S/dagit/internal/git"
)

func main() {
	if git.IsGitRepo() {
		// We'll begin the TUI here
		TUI.Main()
	} else {
		// If its not a git repo then we dont open our TUI
		fmt.Println("err: this isn't a git repo")
	}
}
