package main

import (
	"github.com/sttk-go/sabi"
)

const (
	MODE_NORMAL = iota
	MODE_HELP
	MODE_VERSION
)

type WhoamiDax interface {
	GetMode() int
	GetEffectiveUserId() string
	GetUserNameByUserId(uid string) string
	PrintUserName(userName string)
	PrintHelp()
	PrintVersion()
}

func WhoamiLogic(dax WhoamiDax) sabi.Err {
	switch dax.GetMode() {
	case MODE_VERSION:
		dax.PrintVersion()
	case MODE_HELP:
		dax.PrintHelp()
	default:
		runWhoami(dax)
	}
	return sabi.Ok()
}

func runWhoami(dax WhoamiDax) {
	euid := dax.GetEffectiveUserId()
	userName := dax.GetUserNameByUserId(euid)
	dax.PrintUserName(userName)
}
