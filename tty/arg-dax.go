package main

import (
	"github.com/sttk-go/sabi"
	"os"
)

type ArgDax struct {
}

func NewArgDax() ArgDax {
	return ArgDax{}
}

func (dax ArgDax) GetMode() (int, sabi.Err) {
	mode := MODE_NORMAL
	err := sabi.Ok()

	for _, arg := range os.Args[1:] {
		switch arg {
		case "--version":
			return MODE_VERSION, sabi.Ok()

		case "--help":
			return MODE_HELP, sabi.Ok()

		case "-s":
			fallthrough
		case "--silent":
			fallthrough
		case "--quiet":
			mode = MODE_SILENT

		default:
			mode = MODE_ERROR
			err = sabi.NewErr(InvalidOption{Option: arg})
		}
	}

	return mode, err
}
