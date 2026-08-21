package TUI

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
}

// WAIT! WHY not do it in main()
//
// func (m *model) assigndata() (string, string, string, string, error) {
// 	var err error
// 	// We get our repo Name
// 	m.reponame, err = git.GetRepoName()
// 	if err != nil {
// 		return "", "", "", "", err
// 	}
//
// 	// Owner Name
// 	m.ownername, err = git.GetRepoOwner()
// 	if err != nil {
// 		return "", "", "", "", err
// 	}
//
// 	// The ur;
// 	m.url, err = git.GetRemoteURL()
// 	if err != nil {
// 		return "", "", "", "", err
// 	}
//
// 	// Current Branch name
// 	m.currentbranch, err = git.GetCurrentBranch()
// 	if err != nil {
// 		return "", "", "", "", nil
// 	}
//
// 	// Now we return all the values
// 	return m.reponame, m.ownername, m.url, m.currentbranch, err
// }

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		if msg.String() == "ctrl+c" || msg.String() == "q" {
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() tea.View {
	dagit := "DAGIT"
	str := fmt.Sprintf("%s\n\nRepo: %s\nOwner: %sRemote URL:\n %s\nCurrent Branch: %s", dagit, m.reponame, m.ownername, m.url, m.currentbranch)

	v := tea.NewView(str)
	v.AltScreen = true

	return v
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
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v", err)
	}
}
