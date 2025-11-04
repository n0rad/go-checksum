//go:build build
// +build build

package main

import (
	"github.com/n0rad/gomake"
)

func main() {
	gomake.ProjectBuilder().
		WithName("filesum").
		WithStep(&gomake.StepBuild{
			Programs: []gomake.Program{
				{
					BinaryName: "filesum",
				},
			},
		}).
		WithStep(&gomake.StepRelease{
			OsArchRelease: []string{"linux-amd64", "darwin-amd64"},
		}).
		MustBuild().MustExecute()
}
