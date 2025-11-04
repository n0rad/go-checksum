package cmd

import (
	"os"

	"github.com/n0rad/go-checksum/pkg/integrity"
	"github.com/n0rad/go-erlog/data"
	"github.com/n0rad/go-erlog/errs"
	"github.com/n0rad/go-erlog/logs"
	"github.com/spf13/cobra"
)

var presetFilenameCRC = Config{
	Pattern:            `(?i)\.(lock)$`,
	PatternIsInclusive: false,
	Hash:               "crc32",
	Strategy:           "filename",
}

func RootCmd() *cobra.Command {
	var configFile string
	var preset string
	var config = &Config{}

	var logLevel string
	cmd := &cobra.Command{
		Use:           os.Args[0],
		SilenceErrors: true,
		SilenceUsage:  true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if logLevel != "" {
				level, err := logs.ParseLevel(logLevel)
				if err != nil {
					logs.WithField("value", logLevel).Fatal("Unknown log level")
				}
				logs.SetLevel(level)
			}

			if cmd.Use == "sum" {
				return nil
			}

			if configFile != "" {
				return config.Load(configFile)
			} else {
				switch preset {
				case "filenameCRC":
					config.Set(&presetFilenameCRC)
					if err := config.Init(); err != nil {
						return errs.WithE(err, "Failed to init preset")
					}
				default:
					return errs.WithF(data.WithField("name", preset), "Unknown preset")
				}
			}
			return nil
		},
	}

	cmd.AddCommand(
		UnsetCommand(config),
		CheckCommand(config),
		ListCommand(config),
		SetCommand(config),
		SumCommand(),
	)

	cmd.PersistentFlags().StringVarP(&preset, "preset", "p", "default", "configuration preset: filenameCRC|")
	cmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "configuration file")
	cmd.PersistentFlags().StringVarP(&logLevel, "log-level", "L", "", "Set log level")
	cmd.MarkFlagsMutuallyExclusive("preset", "config")

	return cmd
}

func runCmdForPath(config *Config, path string, f func(d integrity.Path) func(path string) error) error {
	directory := integrity.Path{
		Regex:     config.regex,
		Inclusive: config.PatternIsInclusive,
		Strategy:  integrity.NewStrategy(config.Strategy, config.Hash),
	}

	return f(directory)(path)
}
