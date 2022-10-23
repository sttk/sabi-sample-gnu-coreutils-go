package main

import (
	"github.com/sttk-go/sabi"
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
	GetStdinTtyname() (ttyname string, err sabi.Err)
	PrintTtyname(ttyname string) sabi.Err
	PrintNotTty(err sabi.Err)
	PrintTtyError(err sabi.Err)
	PrintModeError(err sabi.Err)
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
		_, err := dax.GetStdinTtyname()
		return err

	case MODE_NORMAL:
		ttyname, err := dax.GetStdinTtyname()
		switch err.Reason().(type) {
		case sabi.NoError:
			return dax.PrintTtyname(ttyname)
		case StdinIsNotTty:
			dax.PrintNotTty(err)
			return err
		default:
			dax.PrintTtyError(err)
			return err
		}

	case MODE_VERSION:
		return dax.PrintVersion()

	default: // MODE_HELP:
		return dax.PrintHelp()
	}
}
