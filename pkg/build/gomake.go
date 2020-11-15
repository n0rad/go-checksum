//+build build

package main

import (
	"github.com/n0rad/gomake"
)

func main() {
	gomake.ProjectBuilder().
		WithName("fim").
		WithStep(&gomake.StepBuild{
			Programs: []gomake.Program{
				{
					BinaryName: "fim",
					Package:    "github.com/n0rad/go-checksum/pkg/cli/fim",
				},
				{
					BinaryName: "filesum",
					Package:    "github.com/n0rad/go-checksum/pkg/cli/filesum",
				},
			},
		}).
		WithStep(&gomake.StepRelease{
			OsArchRelease: []string{"linux-amd64", "darwin-amd64"},
		}).
		MustBuild().MustExecute()
}
