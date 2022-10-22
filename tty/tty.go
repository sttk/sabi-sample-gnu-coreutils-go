package main

import (
	"github.com/sttk-go/sabi"
	"os"
)

const (
	MODE_ERROR = iota - 1
	MODE_NORMAL
	MODE_SILENT
	MODE_HELP
	MODE_VERSION
)

type /* error reasons */ (
	InvalidOption struct{ Option string }
	StdinIsNotTty struct{}
	FailToPrint   struct{}
)

type TtyDax interface {
	GetMode() (mode int, err sabi.Err)
	GetTtyName(fd int) (ttyname string, err sabi.Err)
	PrintModeError(err sabi.Err) sabi.Err
	PrintTtyName(ttyname string) sabi.Err
	PrintVersion() sabi.Err
	PrintHelp() sabi.Err
}

func ttyLogic(dax TtyDax) sabi.Err {
	mode, err := dax.GetMode()
	if !err.IsOk() {
		dax.PrintModeError(err)
		return err
	}

	switch mode {
	case MODE_SILENT:
		fd := int(os.Stdin.Fd())
		_, err := dax.GetTtyName(fd)
		return err

	case MODE_NORMAL:
		fd := int(os.Stdin.Fd())
		ttyname, err := dax.GetTtyName(fd)
		if !err.IsOk() {
			return err
		}
		return dax.PrintTtyName(ttyname)

	case MODE_VERSION:
		return dax.PrintVersion()

	default: // MODE_HELP:
		return dax.PrintHelp()
	}
}
