package cmd

import (
	"github.com/n0rad/go-erlog/logs"
	_ "github.com/n0rad/go-erlog/register"
	"github.com/spf13/cobra"
	"os"
)

// -h hash algorism
// --dry-run
// -L log level
// -p pattern
//

// pattern test
// create		// create sum on file missing
// check		// check integrity for file with sum
// run         	// create sum if missing, validate if exists
// agent		// like run but with an agent watching files and periodicly check
// hash replace // replace
//

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
		PatternCommand(),
		SetCommand(),
	)

	cmd.PersistentFlags().StringVarP(&logLevel, "log-level", "L", "", "Set log level")
	return cmd
}
