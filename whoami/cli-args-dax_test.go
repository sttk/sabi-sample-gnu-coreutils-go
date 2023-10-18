package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var origOsArgs = os.Args
var origStdout = os.Stdout

func resetOsArgs() {
	os.Args = origOsArgs
}
func resetStdout() {
	os.Stdout = origStdout
}

func TestWhoami_CliArgsDax_GetMode_noArg(t *testing.T) {
	defer resetOsArgs()

	os.Args = []string{origOsArgs[0]}

	dax := NewCliArgsDax()
	mode, err := dax.GetMode()
	assert.Equal(t, mode, MODE_NORMAL)
	assert.True(t, err.IsOk())
}

func TestWhoami_CliArgsDax_GetMode_oneArg_help(t *testing.T) {
	defer resetOsArgs()

	os.Args = []string{origOsArgs[0], "--help"}

	dax := NewCliArgsDax()
	mode, err := dax.GetMode()
	assert.Equal(t, mode, MODE_HELP)
	assert.True(t, err.IsOk())
}

func TestWhoami_CliArgsDax_GetMode_oneArg_version(t *testing.T) {
	defer resetOsArgs()

	os.Args = []string{origOsArgs[0], "--version"}

	dax := NewCliArgsDax()
	mode, err := dax.GetMode()
	assert.Equal(t, mode, MODE_VERSION)
	assert.True(t, err.IsOk())
}

func TestWhoami_CliArgsDax_GetMode_oneArg_invalidOption(t *testing.T) {
	defer resetOsArgs()

	os.Args = []string{origOsArgs[0], "abc"}

	dax := NewCliArgsDax()
	mode, err := dax.GetMode()
	assert.Equal(t, mode, MODE_HELP)

	switch r := err.Reason().(type) {
	case InvalidOption:
		assert.Equal(t, r.Option, "abc")
	default:
		assert.Fail(t, err.Error())
	}
}

func TestWhoami_CliArgsDax_GetMode_twoArg_help(t *testing.T) {
	defer resetOsArgs()

	os.Args = []string{origOsArgs[0], "--help", "--version"}

	dax := NewCliArgsDax()
	mode, err := dax.GetMode()
	assert.Equal(t, mode, MODE_HELP)
	assert.True(t, err.IsOk())
}

func TestWhoami_CliArgsDax_GetMode_twoArg_version(t *testing.T) {
	defer resetOsArgs()

	os.Args = []string{origOsArgs[0], "--version", "-x"}

	dax := NewCliArgsDax()
	mode, err := dax.GetMode()
	assert.Equal(t, mode, MODE_VERSION)
	assert.True(t, err.IsOk())
}

func TestWhoami_CliArgsDax_GetMode_twoArg_invalidLongOption(t *testing.T) {
	defer resetOsArgs()

	os.Args = []string{origOsArgs[0], "--flag", "abc"}

	dax := NewCliArgsDax()
	mode, err := dax.GetMode()

	assert.Equal(t, mode, MODE_HELP)

	switch r := err.Reason().(type) {
	case InvalidOption:
		assert.Equal(t, r.Option, "--flag")
	default:
		assert.Fail(t, err.Error())
	}
}

func TestWhoami_CliArgsDax_GetMode_twoArg_invalidShortOption(t *testing.T) {
	defer resetOsArgs()

	os.Args = []string{origOsArgs[0], "-f", "abc"}

	dax := NewCliArgsDax()
	mode, err := dax.GetMode()

	assert.Equal(t, mode, MODE_HELP)

	switch r := err.Reason().(type) {
	case InvalidOption:
		assert.Equal(t, r.Option, "-f")
	default:
		assert.Fail(t, err.Error())
	}
}

func TestWhoami_CliArgsDax_PrintVersion(t *testing.T) {
	defer resetStdout()

	r, w, _ := os.Pipe()
	os.Stdout = w

	dax := NewCliArgsDax()
	dax.PrintVersion()

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

func TestWhoami_CliArgsDax_PrintHelp(t *testing.T) {
	defer resetStdout()

	r, w, _ := os.Pipe()
	os.Stdout = w

	dax := NewCliArgsDax()
	dax.PrintHelp()

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
