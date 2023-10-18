package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sttk/sabi"
)

func TestTty_app_NewDaxBase(t *testing.T) {
	_, ok := NewDaxBase().(TtyDax)
	assert.True(t, ok)
}

// on go test, OsUserDax#GetStdinTtyName always returns error.
func TestTty_app_noArg(t *testing.T) {
	defer resetOsArgs()
	defer resetStdout()

	os.Args = []string{origOsArgs[0]}

	r, w, _ := os.Pipe()
	os.Stdout = w

	err := sabi.StartApp(app)
	switch err.Reason().(type) {
	default:
		assert.Fail(t, err.Error())
	case StdinIsNotTty:
	}

	w.Close()
	var buf bytes.Buffer
	_, e := buf.ReadFrom(r)
	assert.NoError(t, e)
	assert.Equal(t, buf.String(), "not a tty\n")
}

func TestTty_app_argIsVersion(t *testing.T) {
	defer resetOsArgs()
	defer resetStdout()

	os.Args = []string{origOsArgs[0], "--version"}

	r, w, _ := os.Pipe()
	os.Stdout = w

	err := sabi.StartApp(app)
	assert.True(t, err.IsOk())

	w.Close()
	var buf bytes.Buffer
	_, e := buf.ReadFrom(r)
	assert.NoError(t, e)
	assert.Equal(t, buf.String(), `tty 1.0
Copyright (C) 2022-2023 Takayuki Sato.
License GPLv3+: GNU GPL version 3 or later <https://gnu.org/licenses/gpl.html>.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

`)
}

func TestWhoami_app_argIsHelp(t *testing.T) {
	defer resetOsArgs()
	defer resetStdout()

	os.Args = []string{origOsArgs[0], "--help"}

	r, w, _ := os.Pipe()
	os.Stdout = w

	err := sabi.StartApp(app)
	assert.True(t, err.IsOk())

	w.Close()
	var buf bytes.Buffer
	_, e := buf.ReadFrom(r)
	assert.NoError(t, e)
	assert.Equal(t, buf.String(), `Usage: tty [OPTION]...
Print the file name of the terminal connected to standard input.

      -s, --silent, --quiet  print nothing, only return an exit status.
      --help                 display this help and exit.
      --version              output version information and exit.

`)
}

func TestTty_app_argIsInvalidOption(t *testing.T) {
	defer resetOsArgs()
	defer resetStdout()

	os.Args = []string{origOsArgs[0], "--xxx"}

	r, w, _ := os.Pipe()
	os.Stderr = w

	err := sabi.StartApp(app)

	w.Close()
	var buf bytes.Buffer
	_, e := buf.ReadFrom(r)
	assert.NoError(t, e)
	assert.Equal(t, buf.String(), "extra operand `--xxx'\nTry 'tty --help' for more information.\n")

	switch r := err.Reason().(type) {
	case InvalidOption:
		assert.Equal(t, r.Option, "--xxx")
	default:
		assert.Fail(t, err.Error())
	}
}
