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
	for _, arg := range os.Args[1:] {
		switch arg {
		case "--version":
			return MODE_VERSION
		case "--help":
			return MODE_HELP
		}
	}
	return MODE_NORMAL
}
