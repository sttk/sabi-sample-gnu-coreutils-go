package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sttk/sabi"
)

func TestWhoami_app_NewDaxBase(t *testing.T) {
	_, ok := NewDaxBase().(WhoamiDax)
	assert.True(t, ok)
}

func TestWhoami_app_noArg(t *testing.T) {
	defer resetOsArgs()
	defer resetStdout()

	os.Args = []string{origOsArgs[0]}

	user := os.Getenv("USER")

	r, w, _ := os.Pipe()
	os.Stdout = w

	err := sabi.StartApp(app)
	assert.True(t, err.IsOk())

	w.Close()
	var buf bytes.Buffer
	_, e := buf.ReadFrom(r)
	assert.NoError(t, e)
	assert.Equal(t, buf.String(), user+"\n")
}

func TestWhoami_app_argIsVersion(t *testing.T) {
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
	assert.Equal(t, buf.String(), `whoami 1.0
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
	assert.Equal(t, buf.String(), `Usage: whoami [OPTION]...
Print the user name associated with the current effective user ID.
Same as id -un.

      --help        display this help and exit.
      --version     output version information and exit.

`)
}

func TestWhoami_app_argIsInvalidOption(t *testing.T) {
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
	assert.Equal(t, buf.String(), "extra operand `--xxx'\nTry 'whoami --help' for more information.\n")

	switch r := err.Reason().(type) {
	case InvalidOption:
		assert.Equal(t, r.Option, "--xxx")
	default:
		assert.Fail(t, err.Error())
	}
}
