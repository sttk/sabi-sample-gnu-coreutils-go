package main

import (
	"fmt"
)

type ConsoleDax struct {
}

func NewConsoleDax() ConsoleDax {
	return ConsoleDax{}
}

func (dax ConsoleDax) PrintYes() {
	for {
		fmt.Println("y")
	}
}

func (dax ConsoleDax) PrintWord(word string) {
	for {
		fmt.Println(word)
	}
}

func (dax ConsoleDax) PrintVersion() {
	fmt.Print(`Usage: yes [STRING]...
  or:  yes OPTION
Repeatedly output a line with all specified STRING(s), or 'y'

      --help        display this help and exit.
      --version     output version information and exit.
`)
}

func (dax ConsoleDax) PrintHelp() {
	fmt.Print(`yes 1.0
Copyright (C) 2022 sttk-go project.
License GPLv3+: GNU GPL version 3 or later <https://gnu.org/licenses/gpl.html>.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Written by Takayuki Sato.
`)
}
