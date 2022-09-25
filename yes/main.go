package main

import (
	"github.com/sttk-go/sabi"
)

func newProc() sabi.Proc[yesDax] {
	base := sabi.NewConnBase()
	dax := struct {
		argDax
		consoleDax
	}{
		argDax:     newArgDax(),
		consoleDax: newConsoleDax(),
	}
	return sabi.NewProc[yesDax](base, dax)
}

func main() {
	proc := newProc()
	proc.RunTxn(yesLogic)
}
