package TUI

// lipgloss test 2 -- With no hardcoded sizes
import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
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

var (
	sectionStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			Padding(0, 1)

	footerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
)

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyPressMsg:
		if msg.String() == "ctrl+c" || msg.String() == "q" {
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() tea.View {
	// Can we get resizing done :crying
	if m.width == 0 {
		v := tea.NewView("Loading...")
		v.AltScreen = true
		return v
	}
	// footer first
	footer := footerStyle.Render("q: quit")
	footerHeight := lipgloss.Height(footer)

	// les reserve the height for footer
	availabeHeight := m.height - footerHeight - 2 // -2 becuase -1 is too small (not always but for my hyprland config)
	topHeight := availabeHeight * 2 / 3
	logHeight := availabeHeight - topHeight

	// les split the panel and visualizer
	leftWidth := m.width * 36 / 100
	middleWidth := m.width - leftWidth // i accounted for padding before but it makes it look small idk why

	// left panel
	leftContent := fmt.Sprintf("DAGIT\n\nRepo: %s\n\nOwner: %s\n\nRemote URL:\n %s\n\nCurrent Branch: %s", m.reponame, m.ownername, m.url, m.currentbranch)

	leftSide := sectionStyle.Width(leftWidth).Height(topHeight).Render(leftContent)

	// Middle -- Our main visualiser
	middle := sectionStyle.Width(middleWidth).Height(topHeight).Render("visualiser")

	// Top row
	topRow := lipgloss.JoinHorizontal(lipgloss.Top, leftSide, middle)

	// The git log section
	var logLines []string
	for _, c := range m.logs {
		logLines = append(logLines, fmt.Sprintf("-> [%s] %s (%s)", c.Hash, c.Message, c.Author))
	}
	logContent := strings.Join(logLines, "\n\n")
	logSection := sectionStyle.Width(lipgloss.Width(topRow)).Height(logHeight).Render(logContent)

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
	logs, err := git.GetLog(5)
	if err == nil {
		m.logs = logs
	}
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("dagit refused to run: %v", err)
	}
}
