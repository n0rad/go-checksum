package cmd

import (
	"os"
	"regexp"

	"github.com/ghodss/yaml"
	"github.com/n0rad/go-checksum/pkg/checksum"
	"github.com/n0rad/go-erlog/data"
	"github.com/n0rad/go-erlog/errs"
)

type Config struct {
	Pattern            string
	PatternIsInclusive bool
	Hash               checksum.Hash
	Strategy           string

	regex *regexp.Regexp
}

func (h *Config) Init() error {
	if h.Pattern == "" {
		h.Pattern = `(?i)\.*$`
	}

	if h.Hash == "" {
		h.Hash = checksum.Sha256
	}

	var err error
	h.regex, err = regexp.Compile(h.Pattern)
	if err != nil {
		return errs.WithEF(err, data.WithField("regex", h.regex), "Failed to compile files regex")
	}

	return nil
}

func (h *Config) Load(configPath string) error {
	bytes, err := os.ReadFile(configPath)
	if err != nil {
		return errs.WithEF(err, data.WithField("path", configPath), "Failed to read config file")
	}

	if err := yaml.Unmarshal(bytes, h); err != nil {
		return errs.WithEF(err, data.WithField("content", string(bytes)).WithField("path", configPath), "Failed to parse config file")
	}

	if err := h.Init(); err != nil {
		return errs.WithEF(err, data.WithField("content", string(bytes)), "Failed to init config")
	}
	return nil
}

func (h *Config) Set(c *Config) {
	config := *c
	*h = config
}
