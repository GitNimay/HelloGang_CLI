package cmd

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/inconshreveable/mousetrap"
	"github.com/spf13/cobra"
	"hellogang/internal/greeting"
)

var rootCmd = &cobra.Command{
	Use:   "hellogang",
	Short: "HelloGang — A beautiful CLI greeter with system stats",
	Long: `HelloGang is a terminal-based greeter that shows a fun animated
welcome screen along with real-time system statistics (CPU, RAM, date/time).

Run it without any sub-commands to launch the interactive TUI.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if runtime.GOOS == "windows" && mousetrap.StartedByExplorer() {
			fmt.Println("✨ Double-click detected! Starting the installation...")
			// Call the install command instead of normal run
			if err := installCmd.RunE(installCmd, args); err != nil {
				return err
			}
			// Pause so the user can read the message before closing
			fmt.Println("\nThis window will close in 3 seconds...")
			time.Sleep(3 * time.Second)
			return nil
		}
		return greeting.Run()
	},
}

// Execute runs the root command
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Disable default Cobra double-click prompt so we can handle it
	cobra.MousetrapHelpText = ""

	// Global flags can be added here
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	// Pretty error handling
	rootCmd.SilenceUsage = true
	rootCmd.SilenceErrors = true

	// Version
	rootCmd.Version = "1.0.0"
	rootCmd.SetVersionTemplate(fmt.Sprintf("HelloGang CLI v%s\n", "1.0.0"))

	// Override help template for nicer output
	rootCmd.SetUsageTemplate(`Usage:
  {{.UseLine}}

Available Commands:{{range .Commands}}{{if .IsAvailableCommand}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}

Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}

Use "{{.CommandPath}} [command] --help" for more information about a command.
`)

	_ = os.Stdout
}
