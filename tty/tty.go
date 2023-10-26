package main

import (
	"github.com/sttk/sabi/errs"
)

const (
	MODE_NORMAL = iota
	MODE_SILENT
	MODE_HELP
	MODE_VERSION
)

type (
	InvalidOption struct{ Option string }
	StdinIsNotTty struct{}
	FailToPrint   struct{}
)

type TtyDax interface {
	GetMode() (mode int, err errs.Err)
	GetStdinTtyName() (ttyName string, err errs.Err)
	PrintTtyName(ttyName string) errs.Err
	PrintErr(err errs.Err)
	PrintHelp()
	PrintVersion()
}

func TtyLogic(dax TtyDax) errs.Err {
	mode, err := dax.GetMode()
	if err.IsNotOk() {
		dax.PrintErr(err)
		return err
	}

	switch mode {
	case MODE_SILENT:
		_, err := dax.GetStdinTtyName()
		return err

	case MODE_NORMAL:
		ttynm, err := dax.GetStdinTtyName()
		if err.IsNotOk() {
			dax.PrintErr(err)
			return err
		}
		return dax.PrintTtyName(ttynm)

	case MODE_HELP:
		dax.PrintHelp()

	case MODE_VERSION:
		dax.PrintVersion()
	}

	return errs.Ok()
}
