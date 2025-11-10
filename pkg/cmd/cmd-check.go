package cmd

import (
	"os"

	"github.com/n0rad/file-integrity-manager/pkg/fim"
	"github.com/n0rad/file-integrity-manager/pkg/integrity"
	"github.com/n0rad/go-erlog/data"
	"github.com/n0rad/go-erlog/errs"
	"github.com/spf13/cobra"
)

func CheckCommand() *cobra.Command {
	var output string

	cmd := &cobra.Command{
		Use:   "check",
		Short: "check integrity of files",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var outputFile *os.File
			if output != "" {
				file, err := os.OpenFile(output, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
				if err != nil {
					return errs.WithEF(err, data.WithField("path", output), "Cannot open output file")
				}
				outputFile = file
				defer file.Close()
			}

			for _, arg := range args {
				directory := integrity.Path{
					Regex:     fim.FIM.Regex,
					Inclusive: fim.FIM.PatternIsInclusive,
					Strategy:  integrity.NewStrategy(fim.FIM.Strategy, fim.FIM.Hash),
				}
				return directory.Check(arg, outputFile)
			}
			return nil
		},
	}

	cmd.Flags().StringVarP(&output, "output", "o", "", "output file for invalid or missing")

	return cmd
}
