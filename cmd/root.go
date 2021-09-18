package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/knipferrc/bubbletea-starter/internal/constants"
	"github.com/knipferrc/bubbletea-starter/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "bubbletea-starter",
	Short:   "bubbletea-starter is a starting point for bubbletea apps",
	Version: constants.Versions.AppVersion,
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		m := ui.NewModel()
		var opts []tea.ProgramOption

		// Always append alt screen program option.
		opts = append(opts, tea.WithAltScreen())

		// Initialize new app.
		p := tea.NewProgram(m, opts...)
		if err := p.Start(); err != nil {
			log.Fatal("Failed to start bubbletea-starter", err)
			os.Exit(1)
		}
	},
}

// Execute executes the root command which starts the application.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
