package main

import (
	"flag"
	"fmt"
	"github.com/n0rad/go-checksum"
	"log"
	"os"
)

func usage() {
	println("usage: checksum file ...\n")
	println("supported checksums:")

	for _, hash := range checksum.Hashs {
		println(hash)
		println(" ")
	}
	println()
	os.Exit(1)
}

func main() {
	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() < 1 {
		usage()
	}

	h := checksum.MakeHashString(flag.Arg(0))
	if h == nil {
		log.Printf("Unsupported checksum %q", flag.Arg(0))
		os.Exit(1)
	}

	if flag.NArg() < 2 {
		fileSum, err := checksum.SumFilenameReader(h, os.Stdin, "-")
		if err != nil {
			println(os.Args[0], ": ", err)
			os.Exit(1)
		}
		fmt.Print(fileSum)
	} else {
		for i := 1; i < flag.NArg(); i++ {
			filesum, err := checksum.SumFilename(h, flag.Arg(i))
			if err != nil {
				println(os.Args[0], ": ", err)
			}
			fmt.Print(filesum)
			h.Reset()
		}
	}
}
