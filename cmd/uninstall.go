package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"hellogang/internal/install"
)

var (
	uninstallShell string
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Remove HelloGang from shell startup",
	Long: `Removes HelloGang from your shell's startup configuration.
This undoes what the 'install' command set up.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var shell install.ShellType

		switch uninstallShell {
		case "powershell", "ps":
			shell = install.ShellPowerShell
		case "cmd":
			shell = install.ShellCMD
		case "bash", "git-bash":
			shell = install.ShellBash
		case "auto", "":
			shell = install.DetectShell()
			fmt.Printf("🔍 Detected shell: %s\n", shell)
		case "prompt":
			shell = install.PromptForShell()
		default:
			return fmt.Errorf("unknown shell type: %s (use: powershell, cmd, bash, or auto)", uninstallShell)
		}

		return install.Uninstall(install.InstallOptions{
			Shell: shell,
		})
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)

	uninstallCmd.Flags().StringVarP(&uninstallShell, "shell", "s", "auto",
		"Shell to uninstall from (powershell, cmd, bash, auto, prompt)")
}
