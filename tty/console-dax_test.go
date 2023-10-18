package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sttk/sabi/errs"
)

func TestTty_ConsoleDax_PrintTtyName(t *testing.T) {
	defer resetStdout()

	r, w, _ := os.Pipe()
	os.Stdout = w

	dax := NewConsoleDax()
	dax.PrintTtyName("abc")

	w.Close()
	var buf bytes.Buffer
	_, e := buf.ReadFrom(r)
	assert.NoError(t, e)
	assert.Equal(t, buf.String(), "abc\n")
}

func TestTty_ConsoleDax_PrintErr_invalidOption(t *testing.T) {
	defer resetStdout()

	r, w, _ := os.Pipe()
	os.Stderr = w

	dax := NewConsoleDax()
	err := errs.New(InvalidOption{Option: "--xxx"})
	dax.PrintErr(err)

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
