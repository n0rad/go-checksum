package cmd

import (
	"github.com/n0rad/go-checksum/pkg/integrity"
	"github.com/spf13/cobra"
)

func setWatchCommand(config *Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "watch",
		Short: "Watch new file to add sum",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				directory := integrity.Directory{
					Regex:     config.regex,
					Inclusive: config.PatternIsInclusive,
					Strategy:  integrity.NewSumFileStrategy(config.Hash),
				}

				if err := directory.Watch(arg); err != nil {
					return err
				}
			}
			return nil
		},
	}

	return cmd
}
