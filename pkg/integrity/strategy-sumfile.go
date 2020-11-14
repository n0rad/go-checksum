package integrity

import (
	"github.com/n0rad/go-checksum/pkg/checksum"
	"github.com/n0rad/go-erlog/data"
	"github.com/n0rad/go-erlog/errs"
	"hash"
	"io/ioutil"
	"os"
)

type StrategySumFile struct {
	Hash     hash.Hash
	HashName string
}

func (s StrategySumFile) IsSet(file string) (bool, error) {
	if _, err := os.Stat(s.sumFilename(file)); err == nil {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, nil
	} else {
		return false, err
	}
}

func (s StrategySumFile) GetSum(file string) (string, error) {
	sumFilename := s.sumFilename(file)
	readFile, err := ioutil.ReadFile(sumFilename)
	if err != nil {
		return "", errs.WithEF(err, data.WithField("file", sumFilename), "Failed to read sum file")
	}
	if s.HashName == string(checksum.Crc32_ieee) {
		return checksum.SumFromSumSfvLine(string(readFile)), nil
	} else {
		return checksum.SumFromSumLine(string(readFile)), nil
	}
}

func (s StrategySumFile) Sum(file string) (string, error) {
	sum, err := checksum.SumFile(s.Hash, file)
	if err != nil {
		return "", errs.WithEF(err, data.WithField("file", file), "Failed to sum file")
	}
	return sum, nil
}

func (s StrategySumFile) SumAndSet(file string) (string, error) {
	sum, err := checksum.SumFile(s.Hash, file)
	if err != nil {
		return "", err
	}
	return sum, s.Set(file, sum)
}

func (s StrategySumFile) Set(file string, sum string) error {
	stat, err := os.Stat(file)
	if err != nil {
		return errs.WithEF(err, data.WithField("file", file), "Failed to get stat of file")
	}

	var line string
	if s.HashName == string(checksum.Crc32_ieee) {
		line = checksum.SumSfvLine(file, sum)
	} else {
		line = checksum.SumLine(file, sum)
	}

	sumFilename := s.sumFilename(file)
	if err := ioutil.WriteFile(sumFilename, []byte(line), stat.Mode()); err != nil {
		return errs.WithEF(err, data.WithField("file", sumFilename), "Failed to write sum file")
	}
	return nil
}

func (s StrategySumFile) Remove(file string) error {
	sumFilename := s.sumFilename(file)
	if err := os.Remove(sumFilename); err != nil {
		return errs.WithEF(err, data.WithField("file", sumFilename), "Failed to remove sum file")
	}
	return nil
}

func (s StrategySumFile) Check(file string) (bool, error) {
	sum, err := s.Sum(file)
	if err != nil {
		return false, err
	}

	savedSum, err := s.GetSum(file)
	if savedSum != sum {
		return false, nil
	}
	return true, nil
}

////////////////////////////////////////

func (s StrategySumFile) sumFilename(file string) string {
	if s.HashName == string(checksum.Crc32_ieee) {
		return file + ".sfv"
	}
	return file + "." + s.HashName
}
