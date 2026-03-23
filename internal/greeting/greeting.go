package greeting

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/common-nighthawk/go-figure"
	"hellogang/internal/config"
	"hellogang/internal/stats"
)

// Styles using Lip Gloss for a clean "Claude Orange" theme
var (
	// Claude Orange HEX color
	orange = lipgloss.Color("#DE5B38")
	white = lipgloss.Color("#FAFAFA")

	boxStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(white).
		Padding(1, 4). // Added more padding for nicer spacing
		Margin(1, 0, 1, 2)

	statLabelStyle = lipgloss.NewStyle().
		Foreground(orange)

	statValueStyle = lipgloss.NewStyle().
		Foreground(white).
		Bold(true)

	progressBg = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#3D3D3D")) // Changed to Foreground for text blocks

	progressFg = lipgloss.NewStyle().
		Foreground(orange) // Changed to Foreground for solid text blocks
)

// Run executes the CLI
func Run() error {
	// Fetch system stats instantly (no blocking)
	s, err := stats.GetStats()
	if err != nil {
		fmt.Printf("Error fetching stats: %v\n", err)
		return err
	}

	var b strings.Builder

	// Get username from config
	name := config.GetName()
	greetingText := fmt.Sprintf("HI %s!!!", name)

	// Dynamically generate Graffiti ASCII art
	myFigure := figure.NewFigure(greetingText, "graffiti", true)
	asciiArt := myFigure.String()

	// Add ASCII Art with raw ANSI "Claude Orange" to escape Lipgloss multi-line bugs on Windows
	b.WriteString("\x1b[38;2;222;91;56m\x1b[1m" + asciiArt + "\x1b[0m")
	b.WriteString("\n")

	// Render Info Box with Date/Time and Stats
	infoContent := renderInfoContent(s)
	b.WriteString(boxStyle.Render(infoContent))
	b.WriteString("\n")

	// Print just once and return shell control
	fmt.Println(b.String())

	return nil
}

// renderInfoContent renders the information inside the box
func renderInfoContent(s *stats.SystemStats) string {
	var b strings.Builder

	// Date and Time
	dateLabel := statLabelStyle.Render("📅 Date:")
	dateValue := statValueStyle.Render(s.DateTime)
	b.WriteString(fmt.Sprintf("%s %s\n", dateLabel, dateValue))

	timeLabel := statLabelStyle.Render("⏰ Time:")
	timeValue := statValueStyle.Render(s.Time)
	b.WriteString(fmt.Sprintf("%s %s\n\n", timeLabel, timeValue))

	// CPU Usage
	cpuLabel := statLabelStyle.Render("💻 CPU:")
	cpuBar := renderProgressBar(s.CPUPercent, 20)
	cpuValue := statValueStyle.Render(fmt.Sprintf(" %.1f%%", s.CPUPercent))
	b.WriteString(fmt.Sprintf("%s %s%s\n", cpuLabel, cpuBar, cpuValue))

	// Memory Usage
	memLabel := statLabelStyle.Render("🧠 RAM:")
	memBar := renderProgressBar(s.MemPercent, 20)
	memValue := statValueStyle.Render(fmt.Sprintf(" %.1f%% (%.1f/%.1f GB)",
		s.MemPercent, s.MemUsedGB, s.MemTotalGB))
	b.WriteString(fmt.Sprintf("%s %s%s", memLabel, memBar, memValue))

	return b.String()
}

// renderProgressBar creates a generic progress bar string
func renderProgressBar(percent float64, width int) string {
	filled := int(float64(width) * percent / 100)
	if filled > width {
		filled = width
	}
	empty := width - filled

	filledBlocks := strings.Repeat("█", filled)
	emptyBlocks := strings.Repeat("░", empty)

	return progressFg.Render(filledBlocks) + progressBg.Render(emptyBlocks)
}