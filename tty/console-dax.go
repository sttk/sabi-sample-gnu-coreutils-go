package main

import (
	"fmt"
	"github.com/sttk-go/sabi"
)

type consoleDax struct {
}

func newConsoleDax() consoleDax {
	return consoleDax{}
}

func (dax consoleDax) PrintVersion() sabi.Err {
	_, err := fmt.Print(`tty 1.0
Copyright (C) 2022 sttk-go project.
License GPLv3+: GNU GPL version 3 or later <https://gnu.org/licenses/gpl.html>.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Written by Takayuki Sato.
`)
	if err != nil {
		return sabi.ErrBy(FailToPrint{})
	}
	return sabi.Ok()
}

func (dax consoleDax) PrintHelp() sabi.Err {
	_, err := fmt.Print(`Usage: tty [OPTION]...
Print the file name of the terminal connected to standard input.

      -s, --silent, --quiet   print nothing, only return an exit status
      --help        display this help and exit
      --version     output version information and exit
`)
	if err != nil {
		return sabi.ErrBy(FailToPrint{})
	}
	return sabi.Ok()
}

func (dax consoleDax) PrintTtyName(ttyname string) sabi.Err {
	_, err := fmt.Println(ttyname)
	if err != nil {
		return sabi.ErrBy(FailToPrint{})
	}
	return sabi.Ok()
}

func (dax consoleDax) PrintModeError(err sabi.Err) sabi.Err {
	_, e := fmt.Printf("tty: illegal option: %v\n", err.Get("Option"))
	if e != nil {
		return sabi.ErrBy(FailToPrint{})
	}
	return dax.PrintHelp()
}
