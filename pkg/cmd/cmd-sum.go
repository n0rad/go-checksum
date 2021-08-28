package cmd

import (
	"fmt"
	"github.com/n0rad/go-checksum/pkg/checksum"
	"github.com/spf13/cobra"
	"os"
)

func SumCommand() *cobra.Command {
	var hash string

	cmd := &cobra.Command{
		Use:   "sum",
		Short: "Sum file",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			h := checksum.MakeHashString(hash)
			if h == nil {
				println("Unsupported checksum : ", hash)
				os.Exit(1)
			}

			if len(os.Args) < 3 {
				fileSum, err := checksum.SumLineFromReader(h, os.Stdin, "-")
				if err != nil {
					println(os.Args[0], ": ", err.Error())
					os.Exit(1)
				}
				fmt.Print(fileSum)
			} else {
				for i := 2; i < len(os.Args); i++ {
					stat, err := os.Stat(os.Args[i])
					if err != nil {
						println(os.Args[0], ": ", os.Args[i], ": ", "No such file or directory")
						continue
					}
					if stat.IsDir() {
						println(os.Args[0], ": ", os.Args[i], ": ", "Is a directory")
						continue
					}
					fileSum, err := checksum.SumFilename(h, os.Args[i])
					if err != nil {
						println(os.Args[0], ": ", err.Error())
					}
					fmt.Print(fileSum)
					h.Reset()
				}
			}
			return nil
		},
	}
	cmd.Flags().StringVarP(&hash, "hash", "H", "sha1", "Hash")
	return cmd
}
