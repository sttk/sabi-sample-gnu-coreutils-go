package main

import (
	"os"
)

type CliArgsDax struct {
}

func NewCliArgsDax() CliArgsDax {
	return CliArgsDax{}
}

func (dax CliArgsDax) GetMode() int {
	switch len(os.Args) {
	case 1:
		return MODE_NOWORD
	case 2:
		switch os.Args[1] {
		case "--help":
			return MODE_HELP
		case "--version":
			return MODE_VERSION
		}
		fallthrough
	default:
		return MODE_WORD
	}
}

func (dax CliArgsDax) GetWord() string {
	return os.Args[1]
}
