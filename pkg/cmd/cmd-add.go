package cmd

import (
	"github.com/n0rad/go-checksum/pkg/integrity"
	"github.com/spf13/cobra"
)

func AddCommand() *cobra.Command {
	var configFile string
	var config integrity.Config

	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add integrity to filenames",
		Args:  cobra.MinimumNArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if configFile != "" {
				return config.Load(configFile)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if err := integrity.AddDir(arg, config); err != nil {
					return err
				}
			}
			return nil
		},
	}

	cmd.Flags().StringVarP(&configFile, "config", "c", "", "integrity configuration file")

	return cmd
}
