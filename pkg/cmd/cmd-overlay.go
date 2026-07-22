package cmd

import (
	"github.com/n0rad/file-integrity-manager/pkg/fim"
	"github.com/n0rad/file-integrity-manager/pkg/integrity"
	"github.com/spf13/cobra"
)

func OverlayCommand() *cobra.Command {
	var stripDepth int
	cmd := &cobra.Command{
		Use:   "overlay path target",
		Short: "Replicate the same files structure without integrity in filenames",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			directory := integrity.Path{
				Regex:      fim.FIM.Regex,
				Inclusive:  fim.FIM.PatternIsInclusive,
				Strategy:   integrity.NewStrategy(fim.FIM.Strategy, fim.FIM.Hash),
				StripDepth: stripDepth,
			}
			return directory.Overlay(args[0], args[1])
		},
	}
	cmd.Flags().IntVarP(&stripDepth, "strip", "s", 0, "number of top-level parent folders to strip away (0 = keep full structure)")
	return cmd
}
