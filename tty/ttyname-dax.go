package main

import (
	"github.com/sttk-go/sabi"
	"github.com/sttk-go/sabi-sample-gnu-coreutils/lib"
	"os"
)

type ttynameDax struct {
	sabi.Dax
}

func newTtynameDax(base *sabi.ConnBase) ttynameDax {
	return ttynameDax{Dax: base}
}

func (dax ttynameDax) getTtyConn(name string) (*lib.TtyConn, sabi.Err) {
	conn, err := dax.GetConn(name)
	if !err.IsOk() {
		return nil, err
	}
	return conn.(*lib.TtyConn), sabi.Ok()
}

func (dax ttynameDax) GetStdinTtyname() (string, sabi.Err) {
	conn, err := dax.getTtyConn("ttyname")
	if !err.IsOk() {
		return "", err
	}

	var ttyname string
	fd := int(os.Stdin.Fd())
	ttyname, err = conn.GetTtyname(fd)

	switch err.Reason().(type) {
	case lib.FailToGetTtyname:
		switch err.Reason().(lib.FailToGetTtyname).Errno {
		case lib.ENOTTY:
			return ttyname, sabi.ErrBy(StdinIsNotTty{})
		}
	}
	return ttyname, err
}
