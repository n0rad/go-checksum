package main

import (
	"github.com/n0rad/go-checksum/pkg/cmd"
	"github.com/n0rad/go-erlog/logs"
	"math/rand"
	"os"
	"syscall"
	"time"
)

var Version = "0.0.0"

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	if err := syscall.Setpriority(syscall.PRIO_PROCESS, syscall.Getpid(), 19); err != nil {
		logs.WithE(err).Warn("Failed to set process priority")
	}

	if err := cmd.RootCmd().Execute(); err != nil {
		logs.WithE(err).Fatal("Command failed")
	}
	os.Exit(0)
}
