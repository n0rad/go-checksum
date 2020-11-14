package integrity

import (
	"github.com/n0rad/go-erlog/logs"
	"os"
	"path/filepath"
	"regexp"
)

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

func CheckDir(path string, config IntegrityConfig) error {
	return directoryWalk(path, config.regex, config.PatternIsInclusive, func(path string, info os.FileInfo) {
		if SumFromFilename(config.hash, path) == "" {
			logs.WithField("path", path).Warn("Missing sum in filename")
			return
		}

		ok, err := Check(config.hash, path)
		if err != nil {
			logs.WithField("path", path).Error("Failed to check file integrity")
		}
		if !ok {
			logs.WithField("path", path).Error("File integrity failed")
		}
	})
}

func AddDir(path string, config IntegrityConfig) error {
	return directoryWalk(path, config.regex, config.PatternIsInclusive, func(path string, info os.FileInfo) {
		if SumFromFilename(config.hash, path) == "" {
			newName, err := Add(config.hash, path)
			if err != nil {
				logs.WithE(err).Error("Failed to Add sum to file")
				return
			}
			logs.WithField("path", path).WithField("name", filepath.Base(newName)).Warn("Added sum to filename")
		} else {
			logs.WithField("path", path).Warn("Sum already exists")
		}
	})
}
