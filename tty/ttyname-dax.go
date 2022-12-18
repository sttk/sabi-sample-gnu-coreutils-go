package main

import (
	"github.com/sttk-go/sabi"
	"github.com/sttk-go/sabi-sample-gnu-coreutils/lib"
	"os"
)

type TtyNameDax struct {
	lib.TtyDax
}

func NewTtyNameDax() TtyNameDax {
	return TtyNameDax{TtyDax: lib.NewTtyDax()}
}

func (dax TtyNameDax) GetStdinTtyName() (string, sabi.Err) {
	fd := int(os.Stdin.Fd())
	ttyname, err := dax.GetTtyName(fd)

	switch err.Reason().(type) {
	case lib.FailToGetTtyName:
		switch err.Reason().(lib.FailToGetTtyName).Errno {
		case lib.ENOTTY:
			return ttyname, sabi.ErrBy(StdinIsNotTty{})
		}
	}

	return ttyname, err
}
