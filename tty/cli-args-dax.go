package main

import (
	"os"

	"github.com/sttk/cliargs"
	"github.com/sttk/sabi/errs"
)

var optCfgs = []cliargs.OptCfg{
	cliargs.OptCfg{
		Name:    "s",
		Desc:    "print nothing, only return an exit status.",
		Aliases: []string{"silent", "quiet"},
	},
	cliargs.OptCfg{
		Name: "help",
		Desc: "display this help and exit.",
	},
	cliargs.OptCfg{
		Name: "version",
		Desc: "output version information and exit.",
	},
}

type CliArgsDax struct {
}

func NewCliArgsDax() CliArgsDax {
	return CliArgsDax{}
}

func (dax CliArgsDax) GetMode() (int, errs.Err) {
	cmd, e := cliargs.ParseWith(os.Args, optCfgs)
	if cmd.HasOpt("help") {
		return MODE_HELP, errs.Ok()
	}
	if cmd.HasOpt("version") {
		return MODE_VERSION, errs.Ok()
	}

	if e == nil {
		args := cmd.Args()
		if len(args) == 0 {
			return MODE_NORMAL, errs.Ok()
		}
		return MODE_HELP, errs.New(InvalidOption{Option: args[0]}, e)
	}

	var opt string
	ee, ok := e.(cliargs.InvalidOption)
	if ok {
		opt = ee.GetOpt()
		if len(opt) == 1 {
			opt = "-" + opt
		} else {
			opt = "--" + opt
		}
	}
	return MODE_HELP, errs.New(InvalidOption{Option: opt}, e)
}

func (dax CliArgsDax) PrintHelp() {
	help := cliargs.NewHelp()
	help.AddText("Usage: tty [OPTION]...", 7, 0)
	help.AddText(`Print the file name of the terminal connected to standard input.
`)

	help.AddOpts(optCfgs, 0, 6)
	help.AddText("")
	help.Print()
}

func (dax CliArgsDax) PrintVersion() {
	help := cliargs.NewHelp()
	help.AddText(`tty 1.0
Copyright (C) 2022-2023 Takayuki Sato.
License GPLv3+: GNU GPL version 3 or later <https://gnu.org/licenses/gpl.html>.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.
`)

	help.Print()
}
