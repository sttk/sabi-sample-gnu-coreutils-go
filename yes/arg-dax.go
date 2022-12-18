package main

import (
	"os"
)

type ArgDax struct {
}

func NewArgDax() ArgDax {
	return ArgDax{}
}

func (dax ArgDax) GetMode() int {
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

func (dax ArgDax) GetWord() string {
	return os.Args[1]
}
