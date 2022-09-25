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

func TestArgDax_getMode_noword(t *testing.T) {
	defer ResetArgs()

	os.Args = make([]string, 1)
	os.Args[0] = args[0]

	dax := newArgDax()
	assert.Equal(t, dax.getMode(), mode_noword)
}

func TestArgDax_getMode_word_when_one_arg(t *testing.T) {
	defer ResetArgs()

	os.Args = make([]string, 2)
	os.Args[0] = args[0]
	os.Args[1] = "hello"

	dax := newArgDax()
	assert.Equal(t, dax.getMode(), mode_word)
	assert.Equal(t, dax.getWord(), "hello")
}

func TestArgDax_getMode_help_when_one_arg(t *testing.T) {
	defer ResetArgs()

	os.Args = make([]string, 2)
	os.Args[0] = args[0]
	os.Args[1] = "--help"

	dax := newArgDax()
	assert.Equal(t, dax.getMode(), mode_help)
}

func TestArgDax_getMode_version_when_one_arg(t *testing.T) {
	defer ResetArgs()

	os.Args = make([]string, 2)
	os.Args[0] = args[0]
	os.Args[1] = "--version"

	dax := newArgDax()
	assert.Equal(t, dax.getMode(), mode_version)
}

func TestArgDax_getMode_mode_when_more_than_one_args(t *testing.T) {
	defer ResetArgs()

	os.Args = make([]string, 3)
	os.Args[0] = args[0]
	os.Args[1] = "hello"
	os.Args[2] = "world"

	dax := newArgDax()
	assert.Equal(t, dax.getMode(), mode_word)
	assert.Equal(t, dax.getWord(), "hello")

	os.Args[1] = "--help"
	assert.Equal(t, dax.getMode(), mode_word)
	assert.Equal(t, dax.getWord(), "--help")

	os.Args[1] = "--version"
	assert.Equal(t, dax.getMode(), mode_word)
	assert.Equal(t, dax.getWord(), "--version")
}
