package main

import (
	"flag"
	"fmt"
	"github.com/n0rad/go-hashsum"
	"log"
	"os"
)

func usage() {
	fmt.Fprintln(os.Stderr, "usage: checksum file ...\n")
	fmt.Fprintln(os.Stderr, "supported checksums:")

	for _, hash := range hashsum.Hashs {
		fmt.Fprint(os.Stderr, hash)
		fmt.Fprint(os.Stderr, " ")
	}
	fmt.Fprintln(os.Stderr)
	os.Exit(1)
}

func main() {
	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() < 1 {
		usage()
	}

	h := hashsum.MakeHash(hashsum.Hash(flag.Arg(0)))
	if h == nil {
		log.Fatalf("Unsupported checksum %q", flag.Arg(0))
	}

	if flag.NArg() < 2 {
		filesum, err := hashsum.SumFilenameReader(h, os.Stdin, "-")
		if err != nil {
			fmt.Fprintln(os.Stderr, os.Args[0], ": ", err)
			os.Exit(1)
		}
		fmt.Print(filesum)
	} else {
		for i := 1; i < flag.NArg(); i++ {
			filesum, err := hashsum.SumFilename(h, flag.Arg(i))
			if err != nil {
				fmt.Fprintln(os.Stderr, os.Args[0], ": ", err)
			}
			fmt.Print(filesum)
			h.Reset()
		}
	}
}
