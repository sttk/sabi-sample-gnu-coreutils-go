package main

import (
	"github.com/sttk/sabi/errs"
)

const (
	MODE_NOWORD = iota
	MODE_WORD
	MODE_HELP
	MODE_VERSION
)

type (
	FailToPrint struct{}
)

type YesDax interface {
	GetMode() int
	GetWord() string
	PrintYes() errs.Err
	PrintWord(word string) errs.Err
	PrintVersion()
	PrintHelp()
}

func YesLogic(dax YesDax) errs.Err {
	switch dax.GetMode() {
	case MODE_NOWORD:
		return dax.PrintYes()

	case MODE_WORD:
		return dax.PrintWord(dax.GetWord())

	case MODE_VERSION:
		dax.PrintVersion()

	default: // MODE_HELP
		dax.PrintHelp()
	}

	return errs.Ok()
}
