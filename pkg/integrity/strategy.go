package integrity

import "github.com/n0rad/go-checksum/pkg/checksum"

type Strategy interface {
	IsSet(file string) (bool, error)
	GetSum(file string) (string, error)
	Sum(file string) (string, error)       // TODO generic
	SumAndSet(file string) (string, error) // TODO generic
	Set(file string, sum string) error
	Remove(file string) error
	Check(file string) (bool, error) // TODO generic
}

func NewSumFileStrategy(hash checksum.Hash) StrategySumFile {
	return StrategySumFile{
		Hash:     checksum.NewHash(hash),
		HashName: string(hash),
	}
}

func NewFilenameStrategy(hash checksum.Hash, oldHash checksum.Hash) StrategyFilename {
	return StrategyFilename{
		Hash:    checksum.NewHash(hash),
		OldHash: checksum.NewHash(oldHash),
	}
}
