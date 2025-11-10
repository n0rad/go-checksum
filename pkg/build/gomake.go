//go:build build
// +build build

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
				},
			},
		}).
		WithStep(&gomake.StepRelease{
			OsArchRelease: []string{"linux-amd64", "darwin-amd64", "darwin-arm64", "linux-arm64"},
		}).
		MustBuild().MustExecute()
}
