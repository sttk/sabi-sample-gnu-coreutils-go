package main

import (
	"github.com/sttk-go/sabi"
)

const (
	mode_normal = iota
	mode_help
	mode_version
)

type whoamiDax interface {
	getMode() int
	getEffectiveUserId() string
	getUsernameByUserId(uid string) string
	printUsername(username string)
	printVersion()
	printHelp()
}

func whoamiLogic(dax whoamiDax) sabi.Err {
	switch dax.getMode() {
	case mode_version:
		dax.printVersion()
	case mode_help:
		dax.printHelp()
	default:
		runWhoami(dax)
	}
	return sabi.Ok()
}

func runWhoami(dax whoamiDax) {
	euid := dax.getEffectiveUserId()
	username := dax.getUsernameByUserId(euid)
	dax.printUsername(username)
}
