package main

import (
	"github.com/sttk/sabi"
	"github.com/sttk/sabi/errs"
)

func app() errs.Err {
	base := NewDaxBase()
	defer base.Close()
	return YesLogic(base.(YesDax))
}

func NewDaxBase() sabi.DaxBase {
	base := sabi.NewDaxBase()
	return struct {
		sabi.DaxBase
		CliArgsDax
		ConsoleDax
	}{
		DaxBase:    base,
		CliArgsDax: NewCliArgsDax(),
		ConsoleDax: NewConsoleDax(),
	}
}
