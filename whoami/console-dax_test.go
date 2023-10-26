package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sttk/sabi/errs"
)

func TestWhoami_ConsoleDax_PrintUserName(t *testing.T) {
	defer resetStdout()

	r, w, _ := os.Pipe()
	os.Stdout = w

	dax := NewConsoleDax()
	dax.PrintUserName("abc")

	w.Close()
	var buf bytes.Buffer
	_, e := buf.ReadFrom(r)
	assert.NoError(t, e)
	assert.Equal(t, buf.String(), "abc\n")
}

func TestWhoami_ConsoleDax_PrintErr_invalidOption(t *testing.T) {
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
	assert.Equal(t, buf.String(), "extra operand `--xxx'\nTry 'whoami --help' for more information.\n")

	switch r := err.Reason().(type) {
	case InvalidOption:
		assert.Equal(t, r.Option, "--xxx")
	default:
		assert.Fail(t, err.Error())
	}
}

func TestWhoami_ConsoleDax_PrintErr_failToGetUserName(t *testing.T) {
	defer resetStdout()

	r, w, _ := os.Pipe()
	os.Stderr = w

	dax := NewConsoleDax()
	err := errs.New(FailToGetUserName{Uid: "123"})
	dax.PrintErr(err)

	w.Close()
	var buf bytes.Buffer
	_, e := buf.ReadFrom(r)
	assert.NoError(t, e)
	assert.Equal(t, buf.String(), "cannot find name for user ID 123")

	switch r := err.Reason().(type) {
	case FailToGetUserName:
		assert.Equal(t, r.Uid, "123")
	default:
		assert.Fail(t, err.Error())
	}
}

func TestWhoami_ConsoleDax_PrintErr_others(t *testing.T) {
	defer resetStdout()

	type FailToDoSomething struct{}

	r, w, _ := os.Pipe()
	os.Stderr = w

	dax := NewConsoleDax()
	err := errs.New(FailToDoSomething{})
	dax.PrintErr(err)

	w.Close()
	var buf bytes.Buffer
	_, e := buf.ReadFrom(r)
	assert.NoError(t, e)
	assert.Equal(t, buf.String(), "{reason=FailToDoSomething}")

	switch err.Reason().(type) {
	case FailToDoSomething:
	default:
		assert.Fail(t, err.Error())
	}
}
