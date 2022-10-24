package main

import (
	"github.com/sttk-go/sabi"
	"github.com/sttk-go/sabi-sample-gnu-coreutils/lib"
	"os"
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
	err := newProc().RunTxn(ttyLogic)
	switch err.Reason().(type) {
	case InvalidOption:
		os.Exit(2)
	case StdinIsNotTty:
		os.Exit(1)
	case FailToPrint:
		os.Exit(3)
	default:
		os.Exit(9)
	case sabi.NoError:
	}
}
