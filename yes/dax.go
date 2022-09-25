package main

import (
	"fmt"
	"os"
)

type argDax struct {
}

func newArgDax() argDax {
	return argDax{}
}

func (dax argDax) getMode() int {
	switch len(os.Args) {
	case 1:
		return mode_noword
	case 2:
		switch os.Args[1] {
		case "--help":
			return mode_help
		case "--version":
			return mode_version
		}
		fallthrough
	default:
		return mode_word
	}
}

func (dax argDax) getWord() string {
	return os.Args[1]
}

type consoleDax struct {
}

func newConsoleDax() consoleDax {
	return consoleDax{}
}

func (dax consoleDax) printYes() {
	for {
		fmt.Println("y")
	}
}

func (dax consoleDax) printWord() {
	for {
		fmt.Println(os.Args[1])
	}
}

func (dax consoleDax) printVersion() {
	fmt.Print(`Usage: yes [STRING]...
  or:  yes OPTION
Repeatedly output a line with all specified STRING(s), or 'y'

      --help        display this help and exit.
      --version     output version information and exit.
`)
}

func (dax consoleDax) printHelp() {
	fmt.Print(`yes 1.0
Copyright (C) 2022 sttk-go project.
License GPLv3+: GNU GPL version 3 or later <https://gnu.org/licenses/gpl.html>.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Written by Takayuki Sato.
`)
}
