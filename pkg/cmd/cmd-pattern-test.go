package cmd

import (
	"github.com/spf13/cobra"
)

func patternTestCommand() *cobra.Command {

	var pattern string
	var patternInclusive bool

	cmd := &cobra.Command{
		Use:   "test path",
		Short: "test pattern on tree",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			//p, err := regexp.Compile(pattern)
			//if err != nil {
			//	return errs.WithE(err, "Pattern is not valid")
			//}

			//checksum.DirectoryWalk(args[0], p, patternInclusive, func(path string, info os.FileInfo) {
			//	fmt.Println(path, info.Size())
			//})

			return nil
		},
	}

	//"(?i)\\.(socket|lock)$"
	cmd.Flags().StringVarP(&pattern, "pattern", "p", `(?i)\.*$`, "pattern for files")
	cmd.Flags().BoolVarP(&patternInclusive, "inclusive", "i", false, "pattern is inclusive")

	return cmd
}
