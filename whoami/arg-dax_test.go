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

func TestArgDax_GetMode_noArg(t *testing.T) {
	defer ResetArgs()

	os.Args = make([]string, 1)
	os.Args[0] = args[0]

	dax := NewArgDax()
	assert.Equal(t, dax.GetMode(), MODE_NORMAL)
}

func TestArgDax_GetMode_oneArg_help(t *testing.T) {
	defer ResetArgs()

	os.Args = make([]string, 2)
	os.Args[0] = args[0]
	os.Args[1] = "--help"

	dax := NewArgDax()
	assert.Equal(t, dax.GetMode(), MODE_HELP)
}

func TestArgDax_GetMode_oneArg_version(t *testing.T) {
	defer ResetArgs()

	os.Args = make([]string, 2)
	os.Args[0] = args[0]
	os.Args[1] = "--version"

	dax := NewArgDax()
	assert.Equal(t, dax.GetMode(), MODE_VERSION)
}

func TestArgDax_GetMode_oneArg_normal(t *testing.T) {
	defer ResetArgs()

	os.Args = make([]string, 2)
	os.Args[0] = args[0]
	os.Args[1] = "abc"

	dax := NewArgDax()
	assert.Equal(t, dax.GetMode(), MODE_NORMAL)
}

func TestArgDax_GetMode_twoArgs_help(t *testing.T) {
	defer ResetArgs()

	os.Args = make([]string, 3)
	os.Args[0] = args[0]
	os.Args[1] = "--help"
	os.Args[2] = "--version"

	dax := NewArgDax()
	assert.Equal(t, dax.GetMode(), MODE_HELP)
}

func TestArgDax_GetMode_twoArgs_version(t *testing.T) {
	defer ResetArgs()

	os.Args = make([]string, 3)
	os.Args[0] = args[0]
	os.Args[1] = "--version"
	os.Args[2] = "--help"

	dax := NewArgDax()
	assert.Equal(t, dax.GetMode(), MODE_VERSION)
}

func TestArgDax_GetMode_twoArgs_normal(t *testing.T) {
	defer ResetArgs()
	os.Args = make([]string, 3)
	os.Args[0] = args[0]
	os.Args[1] = "--flag"
	os.Args[2] = "abc"

	dax := NewArgDax()
	assert.Equal(t, dax.GetMode(), MODE_NORMAL)
}
