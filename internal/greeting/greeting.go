package greeting

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"hellogang/internal/animation"
	"hellogang/internal/stats"
)

// Styles using Lip Gloss
var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4")).
			Padding(0, 2).
			Margin(0, 0, 1, 0)

	boxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#7D56F4")).
			Padding(1, 2).
			Margin(1, 0)

	statStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#7D56F4"))

	valueStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA")).
			Bold(true)

	progressBg = lipgloss.NewStyle().
			Background(lipgloss.Color("#3D3D3D"))

	progressFg = lipgloss.NewStyle().
			Background(lipgloss.Color("#7D56F4"))

	faintStyle = lipgloss.NewStyle().
			Faint(true)

	spinnerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#7D56F4"))
)

// Model holds the application state
type Model struct {
	stats      *stats.SystemStats
	animPlayer *animation.Player
	spinnerIdx int
	loading    bool
	err        error
}

// Messages for the update loop
type tickMsg time.Time
type statsLoadedMsg struct {
	stats *stats.SystemStats
	err   error
}

// NewModel creates the initial model
func NewModel() Model {
	return Model{
		animPlayer: animation.NewPlayer(animation.DancingFrames, 500*time.Millisecond),
		loading:    true,
		spinnerIdx: 0,
	}
}

// Init initializes the model
func (m Model) Init() tea.Cmd {
	return tea.Batch(
		loadStats(),
		tickCmd(100*time.Millisecond),
	)
}

// loadStats loads system statistics
func loadStats() tea.Cmd {
	return func() tea.Msg {
		s, err := stats.GetStats()
		return statsLoadedMsg{stats: s, err: err}
	}
}

// tickCmd returns a tick command
func tickCmd(d time.Duration) tea.Cmd {
	return tea.Tick(d, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

// Update handles events
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		}

	case tickMsg:
		// Update animation frame
		m.animPlayer.Next()
		m.spinnerIdx = (m.spinnerIdx + 1) % 10
		return m, tickCmd(300*time.Millisecond)

	case statsLoadedMsg:
		m.loading = false
		if msg.err != nil {
			m.err = msg.err
		} else {
			m.stats = msg.stats
		}
		return m, nil
	}

	return m, nil
}

// View renders the UI
func (m Model) View() string {
	var b strings.Builder

	// Title
	title := titleStyle.Render(" ✨ HelloGang CLI ✨ ")
	b.WriteString(title)
	b.WriteString("\n")

	if m.loading {
		// Loading state
		spinner := animation.GetSpinnerFrame(m.spinnerIdx)
		loadingBox := boxStyle.Render(fmt.Sprintf(
			"%s Loading system information...",
			spinnerStyle.Render(spinner),
		))
		b.WriteString(loadingBox)
		b.WriteString("\n")
	} else if m.err != nil {
		// Error state
		errorBox := boxStyle.Render(fmt.Sprintf("❌ Error: %v", m.err))
		b.WriteString(errorBox)
		b.WriteString("\n")
	} else {
		// Main content
		content := m.renderContent()
		b.WriteString(content)
	}

	// Footer
	footer := faintStyle.Render("Press 'q' to exit")
	b.WriteString("\n")
	b.WriteString(footer)

	return b.String()
}

// renderContent renders the main content
func (m Model) renderContent() string {
	var b strings.Builder

	// Greeting animation
	animFrame := m.animPlayer.Current()
	b.WriteString(boxStyle.Render(animFrame))

	// Date and Time section
	b.WriteString("\n")
	dateTimeBox := m.renderDateTime()
	b.WriteString(dateTimeBox)

	// System Stats section
	b.WriteString("\n")
	statsBox := m.renderStats()
	b.WriteString(statsBox)

	return b.String()
}

// renderDateTime renders date and time
func (m Model) renderDateTime() string {
	if m.stats == nil {
		return ""
	}

	var b strings.Builder

	// Date
	dateLabel := statStyle.Render("📅 Date:")
	dateValue := valueStyle.Render(m.stats.DateTime)
	b.WriteString(fmt.Sprintf("  %s %s\n", dateLabel, dateValue))

	// Time
	timeLabel := statStyle.Render("⏰ Time:")
	timeValue := valueStyle.Render(m.stats.Time)
	b.WriteString(fmt.Sprintf("  %s %s", timeLabel, timeValue))

	return boxStyle.Render(b.String())
}

// renderStats renders system statistics
func (m Model) renderStats() string {
	if m.stats == nil {
		return ""
	}

	var b strings.Builder

	// CPU Usage
	cpuLabel := statStyle.Render("💻 CPU:")
	cpuBar := renderProgressBar(m.stats.CPUPercent, 20)
	cpuValue := valueStyle.Render(fmt.Sprintf(" %.1f%%", m.stats.CPUPercent))
	b.WriteString(fmt.Sprintf("  %s %s%s\n", cpuLabel, cpuBar, cpuValue))

	// Memory Usage
	memLabel := statStyle.Render("🧠 RAM:")
	memBar := renderProgressBar(m.stats.MemPercent, 20)
	memValue := valueStyle.Render(fmt.Sprintf(" %.1f%% (%.1f/%.1f GB)",
		m.stats.MemPercent, m.stats.MemUsedGB, m.stats.MemTotalGB))
	b.WriteString(fmt.Sprintf("  %s %s%s", memLabel, memBar, memValue))

	return boxStyle.Render(b.String())
}

// renderProgressBar creates a progress bar string
func renderProgressBar(percent float64, width int) string {
	filled := int(float64(width) * percent / 100)
	if filled > width {
		filled = width
	}
	empty := width - filled

	bar := strings.Repeat("█", filled) + strings.Repeat("░", empty)
	return progressFg.Render(bar[:filled]) + progressBg.Render(bar[filled:])
}

// Run starts the TUI application
func Run() error {
	p := tea.NewProgram(
		NewModel(),
		tea.WithAltScreen(),
	)
	_, err := p.Run()
	return err
}