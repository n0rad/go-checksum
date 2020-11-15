package cmd

import (
	"github.com/n0rad/go-checksum/pkg/integrity"
	"github.com/spf13/cobra"
)

func ListCommand() *cobra.Command {
	var configFile string
	var config Config
	var reverse bool

	cmd := &cobra.Command{
		Use:   "list",
		Short: "list files",
		Args:  cobra.MinimumNArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return config.Load(configFile)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				inclusive := config.PatternIsInclusive
				if reverse {
					inclusive = !inclusive
				}
				directory := integrity.Directory{
					Regex:     config.regex,
					Inclusive: inclusive,
					Strategy:  integrity.NewSumFileStrategy(config.Hash),
				}

				if err := directory.List(arg); err != nil {
					return err
				}
			}
			return nil
		},
	}

	cmd.Flags().StringVarP(&configFile, "config", "c", `./fim.yaml`, "integrity configuration file")
	cmd.Flags().BoolVarP(&reverse, "reverse", "r", false, "Reverse regex match")

	return cmd
}
