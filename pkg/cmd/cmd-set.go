package cmd

import (
	"github.com/n0rad/go-checksum/pkg/integrity"
	"github.com/spf13/cobra"
)

func SetCommand(config *Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set",
		Short: "Set integrity",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if err := runCmdForPath(config, arg, func(d integrity.Directory) func(path string) error {
					return d.Set
				}); err != nil {
					return err
				}
			}
			return nil
		},
	}
	cmd.AddCommand(setWatchCommand(config))
	return cmd
}
