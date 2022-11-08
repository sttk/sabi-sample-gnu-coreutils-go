package main

import (
	"github.com/sttk-go/sabi"
)

const (
	mode_noword = iota
	mode_word
	mode_help
	mode_version
)

type yesDax interface {
	getMode() int
	getWord() string
	printYes()
	printWord(word string)
	printVersion()
	printHelp()
}

func yesLogic(dax yesDax) sabi.Err {
	switch dax.getMode() {
	case mode_noword:
		dax.printYes()
	case mode_word:
		dax.printWord(dax.getWord())
	case mode_version:
		dax.printVersion()
	default:
		dax.printHelp()
	}
	return sabi.Ok()
}
