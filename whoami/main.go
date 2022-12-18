package main

import (
	"github.com/sttk-go/sabi"
)

func main() {
	proc := newProc()
	proc.RunTxn(WhoamiLogic)
}

func newProc() sabi.Proc[WhoamiDax] {
	base := sabi.NewDaxBase()
	dax := struct {
		ArgDax
		OsDax
		ConsoleDax
	}{
		ArgDax:     NewArgDax(),
		OsDax:      NewOsDax(),
		ConsoleDax: NewConsoleDax(),
	}
	return sabi.NewProc[WhoamiDax](base, dax)
}
