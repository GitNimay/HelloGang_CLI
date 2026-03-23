package stats

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

// SystemStats holds system information
type SystemStats struct {
	DateTime   string
	Time       string
	CPUPercent float64
	MemUsed    uint64
	MemTotal   uint64
	MemPercent float64
	MemUsedGB  float64
	MemTotalGB float64
}

// GetStats retrieves current system statistics
func GetStats() (*SystemStats, error) {
	// Get current date/time
	now := time.Now()

	// Get CPU usage (0 for instant, non-blocking measurement)
	cpuPercents, err := cpu.Percent(0, false)
	if err != nil {
		return nil, fmt.Errorf("failed to get CPU stats: %w", err)
	}

	// Get memory stats
	memStats, err := mem.VirtualMemory()
	if err != nil {
		return nil, fmt.Errorf("failed to get memory stats: %w", err)
	}

	var cpuPercent float64
	if len(cpuPercents) > 0 {
		cpuPercent = cpuPercents[0]
	}

	const GB = 1024 * 1024 * 1024

	return &SystemStats{
		DateTime:   now.Format("Monday, January 02, 2006"),
		Time:       now.Format("15:04:05"),
		CPUPercent: cpuPercent,
		MemUsed:    memStats.Used,
		MemTotal:   memStats.Total,
		MemPercent: memStats.UsedPercent,
		MemUsedGB:  float64(memStats.Used) / GB,
		MemTotalGB: float64(memStats.Total) / GB,
	}, nil
}

// GetDateTime returns formatted current date and time
func GetDateTime() (date string, timeStr string) {
	now := time.Now()
	return now.Format("Monday, January 02, 2006"), now.Format("15:04:05")
}

// FormatBytes converts bytes to human-readable string
func FormatBytes(b uint64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := uint64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "KMGTPE"[exp])
}