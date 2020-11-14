package cmd

import (
	"github.com/n0rad/go-checksum/pkg/integrity"
	"github.com/spf13/cobra"
)

func CheckCommand() *cobra.Command {
	var configFile string
	var config integrity.IntegrityConfig

	cmd := &cobra.Command{
		Use:   "check",
		Short: "check integrity of files",
		Args:  cobra.MinimumNArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return config.Load(configFile)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if err := integrity.CheckDir(arg, config); err != nil {
					return err
				}
			}
			return nil
		},
	}

	cmd.Flags().StringVarP(&configFile, "config", "c", `./fim.yaml`, "integrity configuration file")
	cmd.MarkFlagRequired("config")

	return cmd
}
