package integrity

import (
	"encoding/hex"
	"github.com/n0rad/go-checksum/pkg/checksum"
	"github.com/n0rad/go-erlog/errs"
	"hash"
	"os"
	"path/filepath"
	"strings"
)


type StrategyFilename struct {
	NewHash hash.Hash
	OldHash hash.Hash

}

func TargetFileReplacingSum(oldHash hash.Hash, file string, newSum string) string {
	if file == "" {
		return ""
	}

	b := strings.Builder{}
	pathBase := strings.TrimSuffix(file, filepath.Ext(file))
	if SumFromFilename(oldHash, file) != "" {
		b.WriteString(pathBase[:len(pathBase)-(oldHash.Size()*2)-1])
	} else {
		b.WriteString(pathBase)
	}

	b.WriteRune('-')
	b.WriteString(newSum)
	b.WriteString(filepath.Ext(file))
	return b.String()
}

func SumFromFilename(hash hash.Hash, file string) string {
	fileWithoutExt := strings.TrimSuffix(file, filepath.Ext(file))
	hHexLen := hash.Size() * 2
	if len(filepath.Base(fileWithoutExt))-hHexLen <= 1 { // filename only contains CRC ?
		return ""
	}
	if (fileWithoutExt[len(fileWithoutExt)-hHexLen-1]) != '-' { // crc do not start with a -
		return ""
	}
	candidate := fileWithoutExt[len(fileWithoutExt)-hHexLen:]
	_, err := hex.DecodeString(candidate) // not a crc
	if err != nil {
		return ""
	}
	return candidate
}

//func ensureFilenameContainsHash(h hash.Hash, sumLen int, file string) error {
//	if (SumFromFilename(file, h)) == "" {
//		sum, err := SumFile(h, file)
//		if err != nil {
//			return err
//		})
//		if err := os.Rename(file, TargetFileReplacingSum(file, h, sum)); err != nil {
//			return err
//		}
//	}
//	return nil
//}

func Check(h hash.Hash, file string) (bool, error) {
	sumFile, err := checksum.SumFile(h, file)
	if err != nil {
		return false, err
	}

	sum := SumFromFilename(h, file)
	if sumFile != sum {
		return false, nil
	}
	return true, nil
}

func Add(h hash.Hash, file string) (string, error) {
	sum, err := checksum.SumFile(h, file)
	if err != nil {
		return file, errs.WithE(err, "Failed to create file sum")
	}
	newFile := TargetFileReplacingSum(h, file, sum)
	if err := os.Rename(file, newFile); err != nil {
		return file, err
	}
	return newFile, nil
}

func CheckOrAddIntegrityFilenameSum(h hash.Hash, file string) (bool, error) {
	sumFile, err := checksum.SumFile(h, file)
	if err != nil {
		return false, err
	}

	sum := SumFromFilename(h, file)
	if sum == "" {
		if err := os.Rename(file, TargetFileReplacingSum(h, file, sumFile)); err != nil {
			return false, err
		}
	} else if sumFile != sum {
		return false, nil
	}
	return true, nil
}
