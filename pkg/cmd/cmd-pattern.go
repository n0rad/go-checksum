package cmd

import "github.com/spf13/cobra"

func PatternCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pattern",
		Short: "Pattern",
	}

	cmd.AddCommand(patternTestCommand())

	return cmd
}
