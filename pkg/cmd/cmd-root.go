package cmd

import (
	"github.com/n0rad/go-checksum/pkg/integrity"
	"github.com/n0rad/go-erlog/logs"
	_ "github.com/n0rad/go-erlog/register"
	"github.com/spf13/cobra"
	"os"
)

func RootCmd() *cobra.Command {
	var configFile string
	var config = &Config{}

	var logLevel string
	cmd := &cobra.Command{
		Use:           os.Args[0],
		SilenceErrors: true,
		SilenceUsage:  true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return config.Load(configFile)
		},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if logLevel != "" {
				level, err := logs.ParseLevel(logLevel)
				if err != nil {
					logs.WithField("value", logLevel).Fatal("Unknown log level")
				}
				logs.SetLevel(level)
			}
		},
	}

	cmd.AddCommand(
		RemoveCommand(config),
		CheckCommand(config),
		ListCommand(config),
		SetCommand(config),
	)

	cmd.PersistentFlags().StringVarP(&configFile, "config", "c", `./fim.yaml`, "configuration file")
	cmd.PersistentFlags().StringVarP(&logLevel, "log-level", "L", "", "Set log level")

	cmd.MarkFlagRequired("config")

	return cmd
}

func runCmdForPath(config *Config, path string, f func(d integrity.Directory) func(path string) error) error {
	directory := integrity.Directory{
		Regex:     config.regex,
		Inclusive: config.PatternIsInclusive,
		Strategy:  integrity.NewStrategy(config.Strategy, config.Hash),
	}

	return f(directory)(path)
}
