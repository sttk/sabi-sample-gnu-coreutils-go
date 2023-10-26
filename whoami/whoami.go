package main

import (
	"github.com/sttk/sabi/errs"
)

const (
	MODE_NORMAL = iota
	MODE_HELP
	MODE_VERSION
)

type (
	InvalidOption     struct{ Option string }
	FailToGetUserName struct{ Uid string }
)

type WhoamiDax interface {
	GetMode() (int, errs.Err)
	GetEffectiveUserId() string
	GetUserNameByUserId(uid string) (string, errs.Err)
	PrintUserName(userName string)
	PrintErr(err errs.Err)
	PrintHelp()
	PrintVersion()
}

func WhoamiLogic(dax WhoamiDax) errs.Err {
	mode, err := dax.GetMode()
	if err.IsNotOk() {
		dax.PrintErr(err)
		return err
	}

	switch mode {
	case MODE_VERSION:
		dax.PrintVersion()
		return errs.Ok()
	case MODE_HELP:
		dax.PrintHelp()
		return errs.Ok()
	}

	return runWhoami(dax)
}

func runWhoami(dax WhoamiDax) errs.Err {
	euid := dax.GetEffectiveUserId()

	userName, err := dax.GetUserNameByUserId(euid)
	if err.IsNotOk() {
		dax.PrintErr(err)
		return err
	}

	dax.PrintUserName(userName)
	return errs.Ok()
}
