package main

import (
	"github.com/sttk-go/sabi"
)

func main() {
	proc := newProc()
	proc.RunTxn(YesLogic)
}

func newProc() sabi.Proc[YesDax] {
	base := sabi.NewDaxBase()
	dax := struct {
		ArgDax
		ConsoleDax
	}{
		ArgDax:     NewArgDax(),
		ConsoleDax: NewConsoleDax(),
	}
	return sabi.NewProc[YesDax](base, dax)
}
