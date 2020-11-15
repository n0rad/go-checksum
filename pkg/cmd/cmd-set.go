package cmd

import (
	"github.com/n0rad/go-checksum/pkg/integrity"
	"github.com/spf13/cobra"
)

func SetCommand() *cobra.Command {
	var configFile string
	var config Config

	cmd := &cobra.Command{
		Use:   "set",
		Short: "Set integrity",
		Args:  cobra.MinimumNArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if configFile != "" {
				return config.Load(configFile)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				directory := integrity.Directory{
					Regex:     config.regex,
					Inclusive: config.PatternIsInclusive,
					Strategy:  integrity.NewSumFileStrategy(config.Hash),
				}

				if err := directory.Set(arg); err != nil {
					return err
				}
			}
			return nil
		},
	}

	cmd.Flags().StringVarP(&configFile, "config", "c", "", "integrity configuration file")

	return cmd
}
