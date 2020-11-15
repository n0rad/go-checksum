package cmd

import (
	"github.com/n0rad/go-checksum/pkg/integrity"
	"github.com/spf13/cobra"
)

func ListCommand(config *Config) *cobra.Command {
	var reverse bool
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list files",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				inclusive := config.PatternIsInclusive
				if reverse {
					inclusive = !inclusive
				}
				if err := runCmdForPath(config, arg, func(d integrity.Directory) func(path string) error {
					d.Inclusive = inclusive
					return d.List
				}); err != nil {
					return err
				}
			}
			return nil
		},
	}

	cmd.Flags().BoolVarP(&reverse, "reverse", "r", false, "Reverse regex match")
	return cmd
}
