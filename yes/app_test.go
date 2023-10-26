package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sttk/sabi"
)

func TestYes_app_NewDaxBase(t *testing.T) {
	_, ok := NewDaxBase().(YesDax)
	assert.True(t, ok)
}

func TestYes_app_argsIsVersion(t *testing.T) {
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
	assert.Equal(t, buf.String(), `yes 1.0
Copyright (C) 2022-2023 Takayuki Sato.
License GPLv3+: GNU GPL version 3 or later <https://gnu.org/licenses/gpl.html>.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

`)
}

func TestYes_app_argIsHelp(t *testing.T) {
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
	assert.Equal(t, buf.String(), `Usage: yes [STRING]...
  or:  yes OPTION
Repeatedly output a line with all specified STRING(s), or 'y'

      --help        display this help and exit.
      --version     output version information and exit.

`)
}
