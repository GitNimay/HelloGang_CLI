package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"hellogang/internal/install"
	"hellogang/internal/config"
)

var (
	installShell string
	installForce bool
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install HelloGang to run on shell startup",
	Long: `Installs HelloGang so it runs automatically every time you open
a new terminal session. Supports PowerShell, CMD, and Git Bash.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// --- Prompt for Name ---
		fmt.Print("✨ What is your name? : ")
		reader := bufio.NewReader(os.Stdin)
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)
		if name != "" {
			if err := config.SetName(name); err != nil {
				fmt.Printf("⚠️  Could not save name to config: %v\n", err)
			}
		}

		var shell install.ShellType

		switch installShell {
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
			return fmt.Errorf("unknown shell type: %s (use: powershell, cmd, bash, or auto)", installShell)
		}

		return install.Install(install.InstallOptions{
			Shell: shell,
			Force: installForce,
		})
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	installCmd.Flags().StringVarP(&installShell, "shell", "s", "auto",
		"Shell to install to (powershell, cmd, bash, auto, prompt)")
	installCmd.Flags().BoolVarP(&installForce, "force", "f", false,
		"Force reinstall even if already installed")
}
