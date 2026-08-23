package TUI

import (
	tea "charm.land/bubbletea/v2"
	"github.com/lxsh-S/dagit/internal/git"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyPressMsg:
		if msg.String() == "ctrl+c" || msg.String() == "q" {
			return m, tea.Quit
		} else if msg.String() == "r" {
			logs, err := git.GetLog(5)
			if err == nil {
				m.logs = logs
			}

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
			return m, tea.ClearScreen
		}
	}
	return m, nil
}
