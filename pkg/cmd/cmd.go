package cmd

import (
	"os"

	"github.com/n0rad/file-integrity-manager/pkg/fim"
	"github.com/n0rad/file-integrity-manager/pkg/integrity"
	"github.com/n0rad/go-erlog/logs"
	"github.com/spf13/cobra"
)

func RootCmd() *cobra.Command {
	var logLevel string
	var home string

	cmd := &cobra.Command{
		Use:           os.Args[0],
		SilenceErrors: true,
		SilenceUsage:  true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if logLevel != "" {
				if level, err := logs.ParseLevel(logLevel); err != nil {
					logs.WithField("value", logLevel).Fatal("Unknown log level")
				} else {
					logs.SetLevel(level)
				}
			}

			if err := fim.FIM.Init(home); err != nil {
				return err
			}

			if cmd.Use == "sum" {
				return nil
			}

			return nil
		},
	}

	cmd.AddCommand(
		UnsetCommand(),
		CheckCommand(),
		ListCommand(),
		SetCommand(),
		ServerCommand(),
		SumCommand(),
		OverlayCommand(),
	)

	cmd.PersistentFlags().StringVarP(&logLevel, "log-level", "L", "", "Set log level")
	cmd.PersistentFlags().StringVarP(&home, "home", "H", fim.FIM.DefaultHomeFolder(), "fim home directory")

	return cmd
}

func runCmdForPath(path string, f func(d integrity.Path) func(path string) error) error {

	directory := integrity.Path{
		Regex:     fim.FIM.Regex,
		Inclusive: fim.FIM.PatternIsInclusive,
		Strategy:  integrity.NewStrategy(fim.FIM.Strategy, fim.FIM.Hash),
	}

	return f(directory)(path)
}
