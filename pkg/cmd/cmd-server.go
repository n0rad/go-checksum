package cmd

import (
	"github.com/spf13/cobra"
)

func ServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "start a server",
		RunE: func(cmd *cobra.Command, args []string) error {

			return nil
		},
	}
	return cmd
}
