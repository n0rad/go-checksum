package cmd

import (
	"github.com/n0rad/go-erlog/logs"
	_ "github.com/n0rad/go-erlog/register"
	"github.com/spf13/cobra"
	"os"
)

func RootCmd() *cobra.Command {
	var logLevel string
	cmd := &cobra.Command{
		Use:           os.Args[0],
		SilenceErrors: true,
		SilenceUsage:  true,
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
		RemoveCommand(),
		CheckCommand(),
		ListCommand(),
		SetCommand(),
	)

	cmd.PersistentFlags().StringVarP(&logLevel, "log-level", "L", "", "Set log level")
	return cmd
}
