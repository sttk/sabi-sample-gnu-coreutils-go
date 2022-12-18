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

func TestArgDax_GetMode_if_no_arg(t *testing.T) {
	defer ResetArgs()

	os.Args = make([]string, 1)
	os.Args[0] = args[0]

	dax := NewArgDax()
	mode, err := dax.GetMode()

	assert.Equal(t, mode, MODE_NORMAL)
	assert.True(t, err.IsOk())
}

func TestArgDax_GetMode_if_arg_is_silent(t *testing.T) {
	defer ResetArgs()

	os.Args = make([]string, 2)
	os.Args[0] = args[0]
	os.Args[1] = "--silent"

	dax := NewArgDax()
	mode, err := dax.GetMode()

	assert.Equal(t, mode, MODE_SILENT)
	assert.True(t, err.IsOk())
}

func TestArgDax_GetMode_if_arg_is_s(t *testing.T) {
	defer ResetArgs()

	os.Args = make([]string, 2)
	os.Args[0] = args[0]
	os.Args[1] = "-s"

	dax := NewArgDax()
	mode, err := dax.GetMode()

	assert.Equal(t, mode, MODE_SILENT)
	assert.True(t, err.IsOk())
}

func TestArgDax_GetMode_if_arg_is_quiet(t *testing.T) {
	defer ResetArgs()

	os.Args = make([]string, 2)
	os.Args[0] = args[0]
	os.Args[1] = "--quiet"

	dax := NewArgDax()
	mode, err := dax.GetMode()

	assert.Equal(t, mode, MODE_SILENT)
	assert.True(t, err.IsOk())
}

func TestArgDax_GetMode_if_arg_is_version(t *testing.T) {
	defer ResetArgs()

	os.Args = make([]string, 2)
	os.Args[0] = args[0]
	os.Args[1] = "--version"

	dax := NewArgDax()
	mode, err := dax.GetMode()

	assert.Equal(t, mode, MODE_VERSION)
	assert.True(t, err.IsOk())
}

func TestArgDax_GetMode_if_args_are_version_and_another(t *testing.T) {
	defer ResetArgs()

	os.Args = make([]string, 3)
	os.Args[0] = args[0]
	os.Args[1] = "--silent"
	os.Args[2] = "--version"

	dax := NewArgDax()
	mode, err := dax.GetMode()

	assert.Equal(t, mode, MODE_VERSION)
	assert.True(t, err.IsOk())
}

func TestArgDax_GetMode_if_arg_is_help(t *testing.T) {
	defer ResetArgs()

	os.Args = make([]string, 2)
	os.Args[0] = args[0]
	os.Args[1] = "--help"

	dax := NewArgDax()
	mode, err := dax.GetMode()

	assert.Equal(t, mode, MODE_HELP)
	assert.True(t, err.IsOk())
}

func TestArgDax_GetMode_if_args_are_help_and_another(t *testing.T) {
	defer ResetArgs()

	os.Args = make([]string, 3)
	os.Args[0] = args[0]
	os.Args[1] = "--quiet"
	os.Args[2] = "--help"

	dax := NewArgDax()
	mode, err := dax.GetMode()

	assert.Equal(t, mode, MODE_HELP)
	assert.True(t, err.IsOk())
}

func TestArgDax_GetMode_if_arg_is_invalid(t *testing.T) {
	defer ResetArgs()

	os.Args = make([]string, 2)
	os.Args[0] = args[0]
	os.Args[1] = "-x"

	dax := NewArgDax()
	mode, err := dax.GetMode()

	assert.Equal(t, mode, MODE_ERROR)
	switch err.Reason().(type) {
	case InvalidOption:
		assert.Equal(t, err.Get("Option"), "-x")
		assert.Equal(t, err.Reason().(InvalidOption).Option, "-x")
	default:
		assert.Fail(t, err.Error())
	}
}

func TestArgDax_GetMode_if_args_include_invalid_and_version(t *testing.T) {
	defer ResetArgs()

	os.Args = make([]string, 3)
	os.Args[0] = args[0]
	os.Args[1] = "-x"
	os.Args[2] = "--version"

	dax := NewArgDax()
	mode, err := dax.GetMode()

	assert.Equal(t, mode, MODE_VERSION)
	assert.True(t, err.IsOk())
}

func TestArgDax_GetMode_if_args_include_invalid_and_help(t *testing.T) {
	defer ResetArgs()

	os.Args = make([]string, 3)
	os.Args[0] = args[0]
	os.Args[1] = "-x"
	os.Args[2] = "--help"

	dax := NewArgDax()
	mode, err := dax.GetMode()

	assert.Equal(t, mode, MODE_HELP)
	assert.True(t, err.IsOk())
}
