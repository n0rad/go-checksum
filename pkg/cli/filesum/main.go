package main

import (
	"fmt"
	"github.com/n0rad/go-checksum/pkg/checksum"
	"os"
)

func usage() {
	println("Usage:", os.Args[0], "checksum [file...]\n")
	println("Supported checksums:")
	for _, hash := range checksum.Hashs {
		print(hash)
		print(" ")
	}
	println()
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	h := checksum.MakeHashString(os.Args[1])
	if h == nil {
		println("Unsupported checksum : ", os.Args[1])
		os.Exit(1)
	}

	if len(os.Args) < 3 {
		fileSum, err := checksum.SumLineFromReader(h, os.Stdin, "-")
		if err != nil {
			println(os.Args[0], ": ", err)
			os.Exit(1)
		}
		fmt.Print(fileSum)
	} else {
		for i := 2; i < len(os.Args); i++ {
			fileSum, err := checksum.SumFilename(h, os.Args[i])
			if err != nil {
				println(os.Args[0], ": ", err)
			}
			fmt.Print(fileSum)
			h.Reset()
		}
	}
}
