package TUI

// lipgloss test 2 -- With no hardcoded sizes
import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/lxsh-S/dagit/internal/git"
)

type model struct {
	reponame      string
	ownername     string
	url           string
	currentbranch string
	logs          []git.Commit
	width, height int
}

func (m model) Init() tea.Cmd {
	return nil
}

func Main() {
	m := model{}
	reponame, err := git.GetRepoName()
	if err == nil {
		m.reponame = reponame
	}
	ownername, err := git.GetRepoOwner()
	if err == nil {
		m.ownername = ownername
	}
	url, err := git.GetRemoteURL()
	if err == nil {
		m.url = url
	}
	branch, err := git.GetCurrentBranch()
	if err == nil {
		m.currentbranch = branch
	}
	logs, err := git.GetLog(5)
	if err == nil {
		m.logs = logs
	}

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("dagit refused to run: %v", err)
	}
}
