package main

import (
	"github.com/sttk/sabi"
	"github.com/sttk/sabi/errs"
)

func app() errs.Err {
	base := NewDaxBase()
	defer base.Close()
	return WhoamiLogic(base.(WhoamiDax))
}

func NewDaxBase() sabi.DaxBase {
	base := sabi.NewDaxBase()
	return struct {
		sabi.DaxBase
		CliArgsDax
		OsUserDax
		ConsoleDax
	}{
		DaxBase:    base,
		CliArgsDax: NewCliArgsDax(),
		OsUserDax:  NewOsUserDax(),
		ConsoleDax: NewConsoleDax(),
	}
}
