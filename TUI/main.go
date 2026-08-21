package TUI

// lipgloss test 2 -- With no hardcoded sizes
import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/lxsh-S/dagit/internal/git"
)

type model struct {
	reponame      string
	ownername     string
	url           string
	currentbranch string
	width, height int
}

var (
	sectionStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			Padding(0, 1)

	footerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
)

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
	// left panel
	leftContent := fmt.Sprintf("DAGIT\n\nRepo: %s\n\nOwner: %s\n\nRemote URL:\n %s\n\nCurrent Branch: %s", m.reponame, m.ownername, m.url, m.currentbranch)

	leftSide := sectionStyle.Width(38).Height(20).Render(leftContent)

	// Middle -- Our main visualiser
	middle := sectionStyle.Width(68).Height(20).Render("visualiser")

	// Top row
	topRow := lipgloss.JoinHorizontal(lipgloss.Top, leftSide, middle)

	// The git log section
	logSection := sectionStyle.Width(lipgloss.Width(topRow)).Height(8).Render("Log")

	// footer like nvim yayay!!
	footer := footerStyle.Render("q: quit")

	fullThing := lipgloss.JoinVertical(lipgloss.Left, topRow, logSection, footer)
	v := tea.NewView(fullThing)
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
