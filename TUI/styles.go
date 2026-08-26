package TUI

import "charm.land/lipgloss/v2"

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

	headHashStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("13")).Bold(true)
	headRefStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("46")).Bold(true)
	headNodeChar  = "●"
)

// For statusPanel
var (
	statusStagedStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("10"))
	statusModifiedStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("3"))
	statusUntrackedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
)
