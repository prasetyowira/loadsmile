package lunchcli

import (
	"github.com/spf13/cobra"

	"github.com/prasetyowira/loadsmile/internal/app/lunchcli/command"
)

// Configure configures a root command.
func Configure(rootCmd *cobra.Command) {
	command.AddCommands(rootCmd)
}
