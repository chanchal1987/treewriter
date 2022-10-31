package treewriter

import "strings"

const (
	DefaultLinkPrefix = "│  "
	DefaultMidPrefix  = "├──"
	DefaultEndPrefix  = "└──"
	DefaultNewline    = "\n"
)

type Config struct {
	linkPrefix string
	midPrefix  string
	endPrefix  string
	newline    string
	replacer   strings.Replacer
}

func NewConfig(linkPrefix, midPrefix, endPrefix, newline string) *Config {
	return &Config{
		linkPrefix: linkPrefix,
		midPrefix:  midPrefix,
		endPrefix:  endPrefix,
		newline:    newline,
		replacer:   *strings.NewReplacer(midPrefix, linkPrefix, endPrefix, linkPrefix),
	}
}

//nolint:gochecknoglobals
var DefaultConfig = NewConfig(DefaultLinkPrefix, DefaultMidPrefix, DefaultEndPrefix, DefaultNewline)
