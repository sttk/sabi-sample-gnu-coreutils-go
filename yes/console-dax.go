package main

import (
	"fmt"
	"io"
	"os"

	"github.com/sttk/cliargs"
	"github.com/sttk/sabi/errs"
)

type ConsoleDax struct {
	writer io.Writer
}

func NewConsoleDax() ConsoleDax {
	return ConsoleDax{writer: os.Stdout}
}

func (dax ConsoleDax) PrintYes() errs.Err {
	for {
		_, e := fmt.Fprintln(dax.writer, "y")
		if e != nil {
			return errs.New(FailToPrint{}, e)
		}
	}
}

func (dax ConsoleDax) PrintWord(word string) errs.Err {
	for {
		_, e := fmt.Fprintln(dax.writer, word)
		if e != nil {
			return errs.New(FailToPrint{}, e)
		}
	}
}

func (dax ConsoleDax) PrintHelp() {
	help := cliargs.NewHelp()
	help.AddText("Usage: yes [STRING]...", 8, 0)
	help.AddText("  or:  yes OPTION", 6, 2)
	help.AddText(`Repeatedly output a line with all specified STRING(s), or 'y'
`)

	help.AddOpts([]cliargs.OptCfg{
		cliargs.OptCfg{
			Name: "help",
			Desc: "display this help and exit.",
		},
		cliargs.OptCfg{
			Name: "version",
			Desc: "output version information and exit.",
		},
	}, 14, 6)

	help.AddText("")
	help.Print()
}

func (dax ConsoleDax) PrintVersion() {
	help := cliargs.NewHelp()
	help.AddText(`yes 1.0
Copyright (C) 2022-2023 Takayuki Sato.
License GPLv3+: GNU GPL version 3 or later <https://gnu.org/licenses/gpl.html>.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.
`)
	help.Print()
}
