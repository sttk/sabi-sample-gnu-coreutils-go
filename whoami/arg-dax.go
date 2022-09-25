package main

import (
	"os"
)

type argDax struct {
}

func newArgDax() argDax {
	return argDax{}
}

func (dax argDax) getMode() int {
	for _, arg := range os.Args[1:] {
		switch arg {
		case "--version":
			return mode_version
		case "--help":
			return mode_help
		}
	}
	return mode_normal
}
