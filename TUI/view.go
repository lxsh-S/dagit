package TUI

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/lxsh-S/dagit/internal/git"
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

	isHead := strings.Contains(refs, "HEAD")

	renderedGraph := graphStyle.Render(graphPrefix)
	if isHead {
		renderedGraph = strings.Replace(graphPrefix, "*", headNodeChar, 1)
		renderedGraph = lipgloss.NewStyle().Foreground(lipgloss.Color("46")).Render(renderedGraph)
	}

	var out string
	if isHead {
		out = renderedGraph + headHashStyle.Render(hash) + " "
	} else {
		out = renderedGraph + hashStyle.Render(hash) + " "
	}
	out += authorStyle.Render(author) + " "
	out += timeStyle.Render(when)

	// out := graphStyle.Render(graphPrefix) + hashStyle.Render(hash) + " "
	// out += authorStyle.Render(author) + " "
	// out += timeStyle.Render(when)

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
			if isHead {
				out += refStyle.Render(refs)
			} else {
				out += refStyle.Render(refs)
			}
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
