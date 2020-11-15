package integrity

import (
	"github.com/n0rad/go-erlog/logs"
	"os"
	"path/filepath"
	"regexp"
)

type Directory struct {
	Regex     *regexp.Regexp
	Inclusive bool
	Strategy  Strategy
}

func (d Directory) List(path string) error {
	return d.directoryWalk(path, func(path string, info os.FileInfo) {
		println(path)
	})
}

func (d Directory) Check(path string) error {
	return d.directoryWalk(path, func(path string, info os.FileInfo) {
		set, err := d.Strategy.IsSet(path)
		if err != nil {
			logs.WithE(err).Error("Failed to check if sum is set")
			return
		}
		if !set {
			logs.WithField("path", path).Warn("Missing sum")
			return
		}

		ok, err := d.Strategy.Check(path)
		if err != nil {
			logs.WithField("path", path).Error("Failed to check file integrity")
		}
		if ok != nil {
			logs.WithE(ok).WithField("path", path).Error("File integrity failed")
		}
	})
}

func (d Directory) Set(path string) error {
	return d.directoryWalk(path, func(path string, info os.FileInfo) {
		if d.Strategy.IsSumFile(path) {
			return
		}

		set, err := d.Strategy.IsSet(path)
		if err != nil {
			logs.WithE(err).Error("Failed to check if sum is set")
			return
		}

		if !set {
			logs.WithField("path", path).Info("Processing file")
			if _, err := d.Strategy.SumAndSet(path); err != nil {
				logs.WithE(err).Error("Failed to set sum")
				return
			}
		} else {
			logs.WithField("path", path).Debug("Sum already exists")
		}
	})
}

func (d Directory) Remove(path string) error {
	return d.directoryWalk(path, func(path string, info os.FileInfo) {
		if err := d.Strategy.Remove(path); err != nil {
			logs.WithE(err).WithField("path", path).Error("Failed to remove integrity")
		}
	})
}

////////////////////

func (d Directory) directoryWalk(path string, f func(path string, info os.FileInfo)) error {
	return filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if d.Strategy.IsSumFile(path) {
			return nil
		}

		if err != nil {
			logs.WithE(err).WithField("path", path).Error("Failed to process path")
			return nil
		}
		if info.IsDir() {
			return nil
		}
		if d.Inclusive && d.Regex.MatchString(path) ||
			!d.Inclusive && !d.Regex.MatchString(path) {
			f(path, info)
		}
		return nil
	})
}
