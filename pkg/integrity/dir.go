package integrity

import (
	"github.com/n0rad/go-erlog/logs"
	"os"
	"path/filepath"
	"regexp"
)

type Directory struct {
	fileRegexp *regexp.Regexp
	inclusive  bool
	strategy   Strategy
}

func (d Directory) Check(path string) {
	directoryWalk(path, d.fileRegexp, d.inclusive, checkFunc(d.strategy))
}

func (d Directory) Set(path string) {
	directoryWalk(path, d.fileRegexp, d.inclusive, setFunc(d.strategy))
}

func directoryWalk(path string, fileRegexp *regexp.Regexp, inclusive bool, f func(path string, info os.FileInfo)) error {
	return filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			logs.WithE(err).WithField("path", path).Error("Failed to process path")
			return err
		}
		if info.IsDir() {
			return nil
		}
		if inclusive && fileRegexp.MatchString(path) ||
			!inclusive && !fileRegexp.MatchString(path) {
			f(path, info)
		}
		return nil
	})
}

func checkFunc(s Strategy) func(path string, info os.FileInfo) {
	return func(path string, info os.FileInfo) {
		set, err := s.IsSet(path)
		if err != nil {
			logs.WithE(err).Error("Failed to check if sum is set")
			return
		}
		if !set {
			logs.WithField("path", path).Warn("Missing sum in filename")
			return
		}

		ok, err := s.Check(path)
		if err != nil {
			logs.WithField("path", path).Error("Failed to check file integrity")
		}
		if !ok {
			logs.WithField("path", path).Error("File integrity failed")
		}
	}
}

func setFunc(s Strategy) func(path string, info os.FileInfo) {
	return func(path string, info os.FileInfo) {
		set, err := s.IsSet(path)
		if err != nil {
			logs.WithE(err).Error("Failed to check if sum is set")
			return
		}

		if !set {
			logs.WithField("path", path).Info("Processing file")
			if _, err := s.SumAndSet(path); err != nil {
				logs.WithE(err).Error("Failed to set sum")
				return
			}
			logs.WithField("path", path).Info("Sum added")
		} else {
			logs.WithField("path", path).Debug("Sum already exists")
		}
	}
}
