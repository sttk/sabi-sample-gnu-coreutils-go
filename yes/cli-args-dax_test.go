package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var origOsArgs = os.Args

func resetOsArgs() {
	os.Args = origOsArgs
}

func TestYes_CliArgsDax_GetMode_modeIsNoWord(t *testing.T) {
	defer resetOsArgs()

	os.Args = []string{origOsArgs[0]}

	dax := NewCliArgsDax()
	assert.Equal(t, dax.GetMode(), MODE_NOWORD)
}

func TestYes_CliArgsDax_GetMode_modeIsWord(t *testing.T) {
	defer resetOsArgs()

	os.Args = []string{origOsArgs[0], "hello"}

	dax := NewCliArgsDax()
	assert.Equal(t, dax.GetMode(), MODE_WORD)
	assert.Equal(t, dax.GetWord(), "hello")
}

func TestYes_CliArgsDax_GetMode_modeIsHelp(t *testing.T) {
	defer resetOsArgs()

	os.Args = []string{origOsArgs[0], "--help"}

	dax := NewCliArgsDax()
	assert.Equal(t, dax.GetMode(), MODE_HELP)
}

func TestYes_CliArgsDax_GetMode_modeIsVersion(t *testing.T) {
	defer resetOsArgs()

	os.Args = []string{origOsArgs[0], "--version"}

	dax := NewCliArgsDax()
	assert.Equal(t, dax.GetMode(), MODE_VERSION)
}

func TestYes_CliArgsDax_GetMode_modeIsWord_ifMoreThanOneArg(t *testing.T) {
	defer resetOsArgs()

	os.Args = []string{origOsArgs[0], "hello", "world"}

	dax := NewCliArgsDax()
	assert.Equal(t, dax.GetMode(), MODE_WORD)
	assert.Equal(t, dax.GetWord(), "hello")
}

func TestYes_CliArgsDax_GetMode_modeIsWord_ifIncludesHelpButMoreThanOneArg(
	t *testing.T,
) {
	defer resetOsArgs()

	os.Args = []string{origOsArgs[0], "--help", "hello", "world"}

	dax := NewCliArgsDax()
	assert.Equal(t, dax.GetMode(), MODE_WORD)
	assert.Equal(t, dax.GetWord(), "--help")
}

func TestYes_CliArgsDax_GetMode_modeIsWord_ifIncludesVersionButMoreThanOneArg(
	t *testing.T,
) {
	defer resetOsArgs()

	os.Args = []string{origOsArgs[0], "--version", "hello", "world"}

	dax := NewCliArgsDax()
	assert.Equal(t, dax.GetMode(), MODE_WORD)
	assert.Equal(t, dax.GetWord(), "--version")
}
