package fim

import (
	"regexp"
	"time"

	"github.com/n0rad/file-integrity-manager/pkg/checksum"
	"github.com/n0rad/go-app"
	"github.com/n0rad/go-erlog/data"
	"github.com/n0rad/go-erlog/errs"
)

var FIM Fim

func init() {
	FIM.Name = "fim"
}

type Fim struct {
	app.App

	Pattern            string
	PatternIsInclusive bool
	Hash               checksum.Hash
	Strategy           string

	Regex *regexp.Regexp

	Server struct {
		Paths       []string
		Interval    time.Duration
		ScrubWindow string
	}
}

func (f *Fim) Init() error {
	if err := f.App.Init(); err != nil {
		return err
	}

	if f.Pattern == "" {
		f.Pattern = `(?i)\.*$`
		f.PatternIsInclusive = true
	}

	if f.Strategy == "" {
		f.Strategy = "filename"
	}

	if f.Hash == "" {
		f.Hash = checksum.Sha256
	}

	var err error
	f.Regex, err = regexp.Compile(f.Pattern)
	if err != nil {
		return errs.WithEF(err, data.WithField("regex", f.Regex), "Failed to compile files regex")
	}

	return nil
}

//var presetFilenameCRC = fim.Config{
//	Pattern:            `(?i)\.(lock)$`,
//	PatternIsInclusive: false,
//	Hash:               "crc32",
//	Strategy:           "filename",
//}
