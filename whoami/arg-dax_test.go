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

func TestArgDax_getMode_noArg(t *testing.T) {
	defer ResetArgs()

	os.Args = make([]string, 1)
	os.Args[0] = args[0]

	dax := newArgDax()
	assert.Equal(t, dax.getMode(), mode_normal)
}

func TestArgDax_getMode_oneArg_help(t *testing.T) {
	defer ResetArgs()

	os.Args = make([]string, 2)
	os.Args[0] = args[0]
	os.Args[1] = "--help"

	dax := newArgDax()
	assert.Equal(t, dax.getMode(), mode_help)
}

func TestArgDax_getMode_oneArg_version(t *testing.T) {
	defer ResetArgs()

	os.Args = make([]string, 2)
	os.Args[0] = args[0]
	os.Args[1] = "--version"

	dax := newArgDax()
	assert.Equal(t, dax.getMode(), mode_version)
}

func TestArgDax_getMode_oneArg_normal(t *testing.T) {
	defer ResetArgs()

	os.Args = make([]string, 2)
	os.Args[0] = args[0]
	os.Args[1] = "abc"

	dax := newArgDax()
	assert.Equal(t, dax.getMode(), mode_normal)
}

func TestArgDax_getMode_twoArgs_help(t *testing.T) {
	defer ResetArgs()

	os.Args = make([]string, 3)
	os.Args[0] = args[0]
	os.Args[1] = "--help"
	os.Args[2] = "--version"

	dax := newArgDax()
	assert.Equal(t, dax.getMode(), mode_help)
}

func TestArgDax_getMode_twoArgs_version(t *testing.T) {
	defer ResetArgs()

	os.Args = make([]string, 3)
	os.Args[0] = args[0]
	os.Args[1] = "--version"
	os.Args[2] = "--help"

	dax := newArgDax()
	assert.Equal(t, dax.getMode(), mode_version)
}

func TestArgDax_getMode_twoArgs_normal(t *testing.T) {
	defer ResetArgs()

	os.Args = make([]string, 3)
	os.Args[0] = args[0]
	os.Args[1] = "--flag"
	os.Args[2] = "abc"

	dax := newArgDax()
	assert.Equal(t, dax.getMode(), mode_normal)
}
