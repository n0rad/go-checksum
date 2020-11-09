package hashsum

import (
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"os"
)

// calculate sum of file, with filename, as checksum
func SumFilename(h hash.Hash, file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	return SumFilenameReader(h, f, file)
}

// calculate sum of file
func SumFile(h hash.Hash, file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	return SumReader(h, f)
}

// calculate sum of reader
func SumFilenameReader(h hash.Hash, r io.Reader, filename string) (string, error) {
	sum, err := SumReader(h, r)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s  %s\n", sum, filename), nil
}

// calculate sum of reader
func SumReader(h hash.Hash, r io.Reader) (string, error) {
	_, err := io.Copy(h, r)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
