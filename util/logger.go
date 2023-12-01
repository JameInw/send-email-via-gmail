package util

import (
	"io"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

func NewCustomLogger(prefix string) *log.Logger {
	// Set Log Error Style
	log.ErrorLevelStyle = lipgloss.NewStyle().
		SetString("ERROR!!").
		Padding(0, 1, 0, 1).
		Background(lipgloss.AdaptiveColor{
			Light: "203",
			Dark:  "204",
		}).
		Foreground(lipgloss.Color("0"))
	log.KeyStyles["err"] = lipgloss.NewStyle().Foreground(lipgloss.Color("204"))
	log.ValueStyles["err"] = lipgloss.NewStyle().Bold(true)

	// Create a new log System
	logger := log.NewWithOptions(os.Stderr, log.Options{
		// ReportCaller:    true,
		ReportTimestamp: true,
		Prefix:          prefix,
	})

	// logger.SetLevel(log.DebugLevel)

	return logger
}

func NewCustomLoggerv2(prefix string, logFile *os.File) *log.Logger {
	// Set Log Error Style
	log.ErrorLevelStyle = lipgloss.NewStyle().
		SetString("ERROR!!").
		Padding(0, 1, 0, 1).
		Background(lipgloss.AdaptiveColor{
			Light: "203",
			Dark:  "204",
		}).
		Foreground(lipgloss.Color("0"))
	log.KeyStyles["err"] = lipgloss.NewStyle().Foreground(lipgloss.Color("204"))
	log.ValueStyles["err"] = lipgloss.NewStyle().Bold(true)

	MultiWriter := io.MultiWriter(os.Stderr, logFile)

	// Create a new log System
	logger := log.NewWithOptions(MultiWriter, log.Options{
		// ReportCaller:    true,
		ReportTimestamp: true,
		Prefix:          prefix,
	})

	// logger.SetLevel(log.DebugLevel)

	return logger
}
