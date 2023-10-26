package main

import (
	"github.com/sttk/sabi"
	"github.com/sttk/sabi-sample-gnu-coreutils-go/lib"
	"github.com/sttk/sabi/errs"
)

func init() {
	sabi.Uses("os", lib.OsDaxSrc{})
}

func app() errs.Err {
	base := NewDaxBase()
	defer base.Close()
	return TtyLogic(base.(TtyDax))
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
		OsUserDax:  NewOsUserDax(base),
		ConsoleDax: NewConsoleDax(),
	}
}
