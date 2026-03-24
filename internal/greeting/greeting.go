package greeting

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/common-nighthawk/go-figure"
	"hellogang/internal/config"
	"hellogang/internal/stats"
	"hellogang/internal/terminal"
)

// Styles using Lip Gloss for a clean "Claude Orange" theme
var (
	// Claude Orange HEX color
	orange = lipgloss.Color("#DE5B38")
	white  = lipgloss.Color("#FAFAFA")

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
	termSize := terminal.GetSizeFromEnv()
	termWidth := termSize.Width

	s, err := stats.GetStats()
	if err != nil {
		fmt.Printf("Error fetching stats: %v\n", err)
		return err
	}

	var b strings.Builder

	name := config.GetName()
	greetingText := fmt.Sprintf("HI %s!!!", name)

	var asciiArt string
	font := "graffiti"

	if termWidth < 40 {
		asciiArt = greetingText
		font = "small"
	} else if termWidth < 60 {
		font = "straight"
	}

	myFigure := figure.NewFigure(greetingText, font, true)
	asciiArt = myFigure.String()

	b.WriteString("\x1b[38;2;222;91;56m\x1b[1m" + asciiArt + "\x1b[0m")
	b.WriteString("\n")

	barWidth := calculateBarWidth(termWidth)
	boxStyle := calculateBoxStyle(termWidth)

	infoContent := renderInfoContent(s, barWidth)
	b.WriteString(boxStyle.Render(infoContent))
	b.WriteString("\n")

	fmt.Println(b.String())

	return nil
}

func calculateBarWidth(termWidth int) int {
	if termWidth < 50 {
		return 10
	}
	if termWidth < 70 {
		return 15
	}
	if termWidth < 100 {
		return 20
	}
	return 25
}

func calculateBoxStyle(termWidth int) lipgloss.Style {
	style := boxStyle
	if termWidth < 50 {
		style = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(white).
			Padding(0, 2).
			Margin(0, 1, 1, 1)
	} else if termWidth < 70 {
		style = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(white).
			Padding(1, 2).
			Margin(1, 1, 1, 1)
	}
	return style
}

// renderInfoContent renders the information inside the box
func renderInfoContent(s *stats.SystemStats, barWidth int) string {
	var b strings.Builder

	dateLabel := statLabelStyle.Render("📅 Date:")
	dateValue := statValueStyle.Render(s.DateTime)
	b.WriteString(fmt.Sprintf("%s %s\n", dateLabel, dateValue))

	timeLabel := statLabelStyle.Render("⏰ Time:")
	timeValue := statValueStyle.Render(s.Time)
	b.WriteString(fmt.Sprintf("%s %s\n\n", timeLabel, timeValue))

	cpuLabel := statLabelStyle.Render("💻 CPU:")
	cpuBar := renderProgressBar(s.CPUPercent, barWidth)
	cpuValue := statValueStyle.Render(fmt.Sprintf(" %.1f%%", s.CPUPercent))
	b.WriteString(fmt.Sprintf("%s %s%s\n", cpuLabel, cpuBar, cpuValue))

	memLabel := statLabelStyle.Render("🧠 RAM:")
	memBar := renderProgressBar(s.MemPercent, barWidth)
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
