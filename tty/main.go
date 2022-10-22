package main

import (
	"github.com/sttk-go/sabi"
	"github.com/sttk-go/sabi-sample-gnu-coreutils/lib"
)

func newProc() sabi.Proc[TtyDax] {
	base := sabi.NewConnBase()
	dax := struct {
		argDax
		lib.TtyDax
		consoleDax
	}{
		argDax:     newArgDax(),
		TtyDax:     lib.NewTtyDax(),
		consoleDax: newConsoleDax(),
	}
	return sabi.NewProc[TtyDax](base, dax)
}

func main() {
	newProc().RunTxn(ttyLogic)
}
