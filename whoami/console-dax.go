package main

import (
	"fmt"
	"io"
	"os"

	"github.com/sttk/sabi/errs"
)

type ConsoleDax struct {
	writer io.Writer
}

func NewConsoleDax() ConsoleDax {
	return ConsoleDax{writer: os.Stdout}
}

func (dax ConsoleDax) PrintUserName(userName string) {
	fmt.Fprintln(dax.writer, userName)
}

func (dax ConsoleDax) PrintErr(err errs.Err) {
	switch r := err.Reason().(type) {
	case InvalidOption:
		fmt.Fprintf(os.Stderr, "extra operand `%s'\nTry 'whoami --help' for more information.\n", r.Option)
	case FailToGetUserName:
		fmt.Fprintf(os.Stderr, "cannot find name for user ID %s", r.Uid)
	default:
		fmt.Fprintf(os.Stderr, err.Error())
	}
}
