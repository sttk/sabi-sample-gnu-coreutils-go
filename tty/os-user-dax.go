package main

import (
	"os"

	"github.com/sttk/sabi"
	"github.com/sttk/sabi-sample-gnu-coreutils-go/lib"
	"github.com/sttk/sabi/errs"
)

type OsUserDax struct {
	sabi.Dax
}

func NewOsUserDax(dax sabi.Dax) OsUserDax {
	return OsUserDax{Dax: dax}
}

func (dax OsUserDax) GetStdinTtyName() (string, errs.Err) {
	conn, err := sabi.GetDaxConn[lib.OsDaxConn](dax, "os")
	if err.IsNotOk() {
		return "", err
	}

	fd := int(os.Stdin.Fd())
	ttynm, err := conn.GetTtyName(fd)

	if err.IsOk() {
		return ttynm, err
	}

	switch r := err.Reason().(type) {
	case lib.FailToGetTtyName:
		switch r.Errno {
		case lib.ENOTTY:
			err = errs.New(StdinIsNotTty{}, err)
		}
	}

	return "", err
}
