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
	// Les make with something bluish and purple becuause my current wallaper is in something like that too
	colorBorder = lipgloss.Color("61")
	colorAccent = lipgloss.Color("212")
	colorMuted  = lipgloss.Color("241")
	colorSubtle = lipgloss.Color("245")

	sectionStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(colorBorder).
			Padding(0, 1)

	titleStyle = lipgloss.NewStyle().
			Foreground(colorAccent).
			Bold(true)

	footerStyle = lipgloss.NewStyle().
			Foreground(colorMuted).
			Padding(0, 1)
)

// For left panel
var (
	labelStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	valueStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("225"))
	urlStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("1"))
)

// For log
var (
	logHashStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("214"))
	logAuthorStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("81"))
	logArrowStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	logMessageStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("225"))
)

// For visualiser
var (
	graphStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("78"))
	hashStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("214"))
	refStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Bold(true)
	timeStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
	authorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("81"))
)

func renderLeftPanel(m model) string {
	return renderTitle("DAGIT") + fmt.Sprintf(
		"\n%s %s\n%s %s\n%s %s\n\n%s\n%s",
		labelStyle.Render("Repo:  "), valueStyle.Render(m.reponame),
		labelStyle.Render("Owner: "), valueStyle.Render(m.ownername),
		labelStyle.Render("Branch:"), valueStyle.Render(m.currentbranch),
		labelStyle.Render("URL:  "), urlStyle.Render(m.url),
	)
}

func truncateMessages(msg string, maxWidth int) string {
	if lipgloss.Width(msg) <= maxWidth {
		return msg
	}
	runes := []rune(msg)
	if maxWidth <= 3 {
		return string(runes[:maxWidth])
	}
	cut := maxWidth - 3
	if cut > len(runes) {
		cut = len(runes)
	}

	return string(runes[:cut]) + "..."
}

func renderLogLines(c git.Commit, availableWidth int) string {
	arrow := logArrowStyle.Render("->")
	hash := logHashStyle.Render("[" + c.Hash + "]")
	author := logAuthorStyle.Render("[" + c.Author + "]")

	fixedWidth := lipgloss.Width("-> ["+c.Hash+"] ["+c.Author+"]") + 2
	msgWidth := availableWidth - fixedWidth
	if msgWidth < 10 {
		msgWidth = 10
	}

	message := logMessageStyle.Render(truncateMessages(c.Message, msgWidth))
	return fmt.Sprintf("%s %s %s %s", arrow, hash, message, author)
}

func renderTitle(text string) string {
	return titleStyle.Render(text) + "\n" + lipgloss.NewStyle().Foreground(colorBorder).Render(strings.Repeat("-", lipgloss.Width(text)))
}

func renderGraphLine(line string) string {
	parts := strings.SplitN(line, "\x02", 2)
	if len(parts) < 2 {
		return graphStyle.Render(line) // Only the connectors
	}

	graphPrefix, data := parts[0], parts[1]

	fields := strings.SplitN(data, "\x1f", 4)
	if len(fields) < 4 {
		return graphStyle.Render(graphPrefix)
	}

	hash, author, refs, when := fields[0], strings.TrimSpace(fields[1]), fields[2], fields[3]

	out := graphStyle.Render(graphPrefix) + hashStyle.Render(hash) + " "
	out += authorStyle.Render(author) + " "
	out += timeStyle.Render(when)

	if refs != "" {
		refs = strings.Trim(refs, "()")

		// remove Noise
		refs = strings.ReplaceAll(refs, "origin/HEAD", "")
		refs = strings.ReplaceAll(refs, ", ,", ",")
		refs = strings.Trim(refs, ", ")

		// Stop 'refs' being too long
		if lipgloss.Width(refs) > 30 {
			refs = refs[:27] + "..."
		}

		if refs != "" {
			out += " "
			out += refStyle.Render(refs)
		}
	}
	return out
}

func makeTheLinesStop(lines []string, max int) []string {
	if len(lines) > max {
		return lines[:max]
	}
	return lines
}

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

func (m model) View() tea.View {
	// Can we get resizing done :crying
	if m.width == 0 {
		v := tea.NewView("Loading...")
		v.AltScreen = true
		return v
	}

	// ---------
	// ||	HEIGHT|
	// ---------

	// footer first
	footer := footerStyle.Render("q: quit • r: refresh")
	footerHeight := lipgloss.Height(footer)

	verticalFrame := sectionStyle.GetVerticalFrameSize()

	// Now reserve space for footer
	availabeHeight := m.height - footerHeight
	// availabeHeight := m.height - footerHeight - 1 // Before  [switched to 1 form 2]

	topOuterHeight := availabeHeight * 2 / 3
	lopOuterHeight := availabeHeight - topOuterHeight

	topHeight := topOuterHeight - verticalFrame
	logHeight := lopOuterHeight - verticalFrame

	// for a lil safety
	if topHeight < 1 {
		topHeight = 1
	}
	if logHeight < 1 {
		logHeight = 1
	}
	//
	// ---------
	// ||	WIDTH|
	// ---------

	// left panel
	// Attempt to fix right side border issue

	// horizontalFrame := sectionStyle.GetHorizontalFrameSize()

	leftOuterWidth := m.width * 36 / 100
	middleOuterWidth := m.width - leftOuterWidth

	leftWidth := leftOuterWidth
	middleWidth := middleOuterWidth

	// leftContent := renderTitle("DAGIT") + fmt.Sprintf(
	// 	"\nRepo:    %s\nOwner:   %s\nBranch:  %s\n\nURL:     %s", m.reponame, m.ownername, m.currentbranch, m.url)

	leftContent := renderLeftPanel(m)

	leftSide := sectionStyle.Width(leftWidth).Height(topHeight).Render(leftContent)

	// Middle -- Our main visualiser
	graphLines, _ := git.GetGraph(15)                        // We first get the graph lines
	graphLines = git.SpreadGraph(graphLines)                 // Then we spread then (look better)
	graphLines = makeTheLinesStop(graphLines, topHeight-2-2) // Added one more -2 accounting the heading fixed? -> yes

	var rendered []string
	for _, l := range graphLines {
		rendered = append(rendered, renderGraphLine(l))
	}

	middleContent := renderTitle("VISUALIZER") + "\n" + strings.Join(rendered, "\n")
	middle := sectionStyle.Width(middleWidth).Height(topHeight).Render(middleContent)

	// Top row
	topRow := lipgloss.JoinHorizontal(lipgloss.Top, leftSide, middle)

	// The git log section

	maxLogLines := (logHeight - 3)
	if maxLogLines < 1 {
		maxLogLines = 1
	}

	logWidth := lipgloss.Width(topRow) - sectionStyle.GetVerticalFrameSize() - 3

	logs := m.logs

	if len(logs) > maxLogLines {
		logs = logs[:maxLogLines]
	}

	var logLines []string

	for _, c := range logs {
		logLines = append(logLines, renderLogLines(c, logWidth))
	}
	logContent := renderTitle("LOG") + "\n" + strings.Join(logLines, "\n\n")
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
