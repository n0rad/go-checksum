package main

import (
	"github.com/n0rad/go-checksum/pkg/cmd"
	"github.com/n0rad/go-erlog/logs"
	"math/rand"
	"os"
	"time"
)

var Version = "0.0.0"

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	if err := cmd.RootCmd().Execute(); err != nil {
		logs.WithE(err).Fatal("Command failed")
	}
	os.Exit(0)
}
