package TUI

import (
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/lxsh-S/dagit/internal/git"
)

type tickMsg time.Time

var tickRates = []time.Duration{
	500 * time.Millisecond,
	1 * time.Second,
	3 * time.Second,
	5 * time.Second,
	10 * time.Second,
}

func nextTicketRate(current time.Duration) time.Duration {
	for i, r := range tickRates {
		if r == current {
			return tickRates[(i+1)%len(tickRates)]
		}
	}
	return tickRates[0] // If cuurent not in list we fall here
}

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

	if status, err := git.GetStatus(); err == nil {
		m.status = status
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tickMsg:
		m.refreshData()
		return m, tickEveryWhen(m.tickRate)

	case tea.KeyPressMsg:
		if msg.String() == "ctrl+c" || msg.String() == "q" {
			return m, tea.Quit
		} else if msg.String() == "v" {
			m.showStatus = !m.showStatus
		} else if msg.String() == "+" || msg.String() == "=" {
			for i, r := range tickRates {
				if r == m.tickRate && i < len(tickRates)-1 {
					m.tickRate = tickRates[i+1] // slowrr
				}
			}
		} else if msg.String() == "-" {
			for i, r := range tickRates {
				for r == m.tickRate && i > 0 {
					m.tickRate = tickRates[i-1] // faster
				}
			}
		}
	}
	return m, nil
}
