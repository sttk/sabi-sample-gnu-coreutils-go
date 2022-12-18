package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var args []string = os.Args

func ResetArgs() {
	os.Args = args
}

func TestArgDax_GetMode_noword(t *testing.T) {
	defer ResetArgs()

	os.Args = make([]string, 1)
	os.Args[0] = args[0]

	dax := NewArgDax()
	assert.Equal(t, dax.GetMode(), MODE_NOWORD)
}

func TestArgDax_GetMode_word_when_one_arg(t *testing.T) {
	defer ResetArgs()

	os.Args = make([]string, 2)
	os.Args[0] = args[0]
	os.Args[1] = "hello"

	dax := NewArgDax()
	assert.Equal(t, dax.GetMode(), MODE_WORD)
	assert.Equal(t, dax.GetWord(), "hello")
}

func TestArgDax_GetMode_help_when_one_arg(t *testing.T) {
	defer ResetArgs()

	os.Args = make([]string, 2)
	os.Args[0] = args[0]
	os.Args[1] = "--help"

	dax := NewArgDax()
	assert.Equal(t, dax.GetMode(), MODE_HELP)
}

func TestArgDax_GetMode_version_when_one_arg(t *testing.T) {
	defer ResetArgs()

	os.Args = make([]string, 2)
	os.Args[0] = args[0]
	os.Args[1] = "--version"

	dax := NewArgDax()
	assert.Equal(t, dax.GetMode(), MODE_VERSION)
}

func TestArgDax_GetMode_word_when_more_than_one_args(t *testing.T) {
	defer ResetArgs()

	os.Args = make([]string, 3)
	os.Args[0] = args[0]
	os.Args[1] = "hello"
	os.Args[2] = "world"

	dax := NewArgDax()
	assert.Equal(t, dax.GetMode(), MODE_WORD)
	assert.Equal(t, dax.GetWord(), "hello")

	os.Args[1] = "--help"
	assert.Equal(t, dax.GetMode(), MODE_WORD)
	assert.Equal(t, dax.GetWord(), "--help")

	os.Args[1] = "--version"
	assert.Equal(t, dax.GetMode(), MODE_WORD)
	assert.Equal(t, dax.GetWord(), "--version")
}
