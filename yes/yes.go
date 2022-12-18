package main

import (
	"github.com/sttk-go/sabi"
)

const (
	MODE_NOWORD = iota
	MODE_WORD
	MODE_HELP
	MODE_VERSION
)

type YesDax interface {
	GetMode() int
	GetWord() string
	PrintYes()
	PrintWord(word string)
	PrintVersion()
	PrintHelp()
}

func YesLogic(dax YesDax) sabi.Err {
	switch dax.GetMode() {
	case MODE_NOWORD:
		dax.PrintYes()
	case MODE_WORD:
		dax.PrintWord(dax.GetWord())
	case MODE_VERSION:
		dax.PrintVersion()
	default:
		dax.PrintHelp()
	}
	return sabi.Ok()
}
