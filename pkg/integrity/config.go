package integrity

import (
	"github.com/ghodss/yaml"
	"github.com/n0rad/go-checksum/pkg/checksum"
	"github.com/n0rad/go-erlog/data"
	"github.com/n0rad/go-erlog/errs"
	"hash"
	"io/ioutil"
	"regexp"
)

var FIM IntegrityConfig

type IntegrityConfig struct {
	Pattern            string
	PatternIsInclusive bool
	Hash               string
	Strategy           string

	regex *regexp.Regexp
	hash  hash.Hash
}

func (h *IntegrityConfig) Init() error {
	if h.Pattern == "" {
		h.Pattern = `(?i)\.*$`
	}

	var err error
	h.regex, err = regexp.Compile(h.Pattern)
	if err != nil {
		return errs.WithEF(err, data.WithField("regex", h.regex), "Failed to compile files regex")
	}

	h.hash = checksum.MakeHashString(h.Hash)
	if h.hash == nil {
		return errs.WithF(data.WithField("hash", h.Hash), "Unknown hash algorithm")
	}
	return nil
}

func (h *IntegrityConfig) Load(configPath string) error {
	bytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		return errs.WithEF(err, data.WithField("path", configPath), "Failed to read fim config file")
	}

	if err := yaml.Unmarshal(bytes, h); err != nil {
		return errs.WithEF(err, data.WithField("content", string(bytes)).WithField("path", configPath), "Failed to parse fim file")
	}

	if err := h.Init(); err != nil {
		return errs.WithEF(err, data.WithField("content", string(bytes)), "Failed to init hdm file")
	}
	return nil
}
