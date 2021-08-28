//+build build

package main

import (
	"github.com/n0rad/gomake"
)

func main() {
	gomake.ProjectBuilder().
		WithName("checksum").
		WithStep(&gomake.StepBuild{
			Programs: []gomake.Program{
				{
					BinaryName: "checksum",
					Package:    "github.com/n0rad/go-checksum/pkg/cli",
				},
			},
		}).
		WithStep(&gomake.StepRelease{
			OsArchRelease: []string{"linux-amd64", "darwin-amd64"},
		}).
		MustBuild().MustExecute()
}
