package main

import (
	"github.com/sttk-go/sabi"
	"github.com/sttk-go/sabi-sample-gnu-coreutils/lib"
	"os"
)

func main() {
	err := newProc().RunTxn(TtyLogic)

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

func init() {
	sabi.AddGlobalDaxSrc("os", lib.NewOsDaxSrc())
}

func newProc() sabi.Proc[TtyDax] {
	base := sabi.NewDaxBase()
	dax := struct {
		ArgDax
		OsUserDax
		ConsoleDax
	}{
		ArgDax:     NewArgDax(),
		OsUserDax:  NewOsUserDax(base),
		ConsoleDax: NewConsoleDax(),
	}
	return sabi.NewProc[TtyDax](base, dax)
}
