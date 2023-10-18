package main

import (
	"fmt"
	"os"

	"github.com/sttk/sabi/errs"
)

type ConsoleDax struct {
}

func NewConsoleDax() ConsoleDax {
	return ConsoleDax{}
}

func (dax ConsoleDax) PrintTtyName(ttynm string) errs.Err {
	_, e := fmt.Println(ttynm)
	if e != nil {
		return errs.New(FailToPrint{}, e)
	}
	return errs.Ok()
}

func (dax ConsoleDax) PrintErr(err errs.Err) {
	switch r := err.Reason().(type) {
	case InvalidOption:
		fmt.Fprintf(os.Stderr, "extra operand `%s'\nTry 'tty --help' for more information.\n", r.Option)
	case StdinIsNotTty:
		fmt.Println("not a tty")
	default:
		fmt.Println(err.Error())
	}
}
