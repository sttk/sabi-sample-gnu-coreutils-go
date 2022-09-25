package main

import (
	"github.com/sttk-go/sabi"
)

func newProc() sabi.Proc[whoamiDax] {
	base := sabi.NewConnBase()
	dax := struct {
		argDax
		osDax
		consoleDax
	}{
		argDax:     newArgDax(),
		osDax:      newOsDax(),
		consoleDax: newConsoleDax(),
	}
	return sabi.NewProc[whoamiDax](base, dax)
}

func main() {
	newProc().RunTxn(whoamiLogic)
}
