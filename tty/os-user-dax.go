package main

import (
	"github.com/sttk-go/sabi"
	"github.com/sttk-go/sabi-sample-gnu-coreutils/lib"
	"os"
)

type OsUserDax struct {
	lib.OsDax
}

func NewOsUserDax(dax sabi.Dax) OsUserDax {
	return OsUserDax{OsDax: lib.NewOsDax(dax)}
}

func (dax OsUserDax) GetStdinTtyName() (string, sabi.Err) {
	conn, err := dax.GetOsDaxConn("os")
	if !err.IsOk() {
		return "", err
	}

	fd := int(os.Stdin.Fd())
	ttyname, err := conn.GetTtyName(fd)

	switch err.Reason().(type) {
	case lib.FailToGetTtyName:
		switch err.Reason().(lib.FailToGetTtyName).Errno {
		case lib.ENOTTY:
			return ttyname, sabi.ErrBy(StdinIsNotTty{})
		}
	}

	return ttyname, err
}
