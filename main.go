package main

import (
	"os"
	"syscall"

	"github.com/n0rad/file-integrity-manager/pkg/cmd"
	"github.com/n0rad/file-integrity-manager/pkg/fim"
	"github.com/n0rad/go-erlog/logs"
	_ "github.com/n0rad/go-erlog/register"
)

var Version = "0.0.0"

func main() {
	if err := fim.FIM.SetVersion(Version); err != nil {
		logs.WithE(err).Fatal("Failed to set internal build version")
	}

	if err := syscall.Setpriority(syscall.PRIO_PROCESS, syscall.Getpid(), 19); err != nil {
		logs.WithE(err).Warn("Failed to set process priority")
	}

	if err := cmd.RootCmd().Execute(); err != nil {
		logs.WithE(err).Fatal("Command failed")
	}
	os.Exit(0)
}
