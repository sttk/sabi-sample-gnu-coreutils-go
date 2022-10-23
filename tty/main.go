package main

import (
	"github.com/sttk-go/sabi"
	"github.com/sttk-go/sabi-sample-gnu-coreutils/lib"
)

func init() {
	sabi.AddGlobalConnCfg("ttyname", lib.NewTtyConnCfg())
	sabi.SealGlobalConnCfgs()
}

func newProc() sabi.Proc[TtyDax] {
	base := sabi.NewConnBase()
	dax := struct {
		argDax
		ttynameDax
		consoleDax
	}{
		argDax:     newArgDax(),
		ttynameDax: newTtynameDax(base),
		consoleDax: newConsoleDax(),
	}
	return sabi.NewProc[TtyDax](base, dax)
}

func main() {
	newProc().RunTxn(ttyLogic)
}
