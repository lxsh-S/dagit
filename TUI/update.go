package TUI

import (
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/lxsh-S/dagit/internal/git"
)

type tickMsg time.Time

func tickEveryWhen(d time.Duration) tea.Cmd {
	return tea.Tick(d, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m *model) refreshData() {
	if logs, err := git.GetLog(5); err == nil {
		m.logs = logs
	}

	if reponame, err := git.GetRepoName(); err == nil {
		m.reponame = reponame
	}

	if ownername, err := git.GetRepoOwner(); err == nil {
		m.url = ownername
	}

	if url, err := git.GetRemoteURL(); err == nil {
		m.url = url
	}

	if branch, err := git.GetCurrentBranch(); err == nil {
		m.currentbranch = branch
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tickMsg:
		m.refreshData()
		return m, tickEveryWhen(1 * time.Second)

	case tea.KeyPressMsg:
		if msg.String() == "ctrl+c" || msg.String() == "q" {
			return m, tea.Quit
		}
	}
	return m, nil
}
