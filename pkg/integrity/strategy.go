package integrity

import "github.com/n0rad/go-checksum/pkg/checksum"

type Strategy interface {
	IsSet(file string) (bool, error)
	GetSum(file string) (string, error)
	Sum(file string) (string, error)       // TODO generic
	SumAndSet(file string) (string, error) // TODO generic
	Set(file string, sum string) error
	Remove(file string) error
	Check(file string) (error, error) // TODO generic
	IsSumFile(file string) bool
}

func NewStrategy(strategyName string, hash checksum.Hash) Strategy {
	switch strategyName {
	case "sumfile":
		return StrategySumFile{
			Hash:     checksum.NewHash(hash),
			HashName: string(hash),
		}
	case "filename":
		return StrategyFilename{
			Hash:    checksum.NewHash(hash),
			OldHash: checksum.NewHash(hash), // support old HASH
		}
	}
	return nil
}
