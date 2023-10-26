package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var origStdout = os.Stdout

func resetStdout() {
	os.Stdout = origStdout
}

type errorableWriter struct {
	back  io.Writer
	count int
}

func (w *errorableWriter) Write(b []byte) (int, error) {
	if w.count >= 5 {
		return 0, fmt.Errorf("write error")
	}
	w.count++
	return w.back.Write(b)
}

func TestYes_ConsoleDax_PrintYes(t *testing.T) {
	defer resetStdout()

	r, w, _ := os.Pipe()

	dax := NewConsoleDax()
	dax.writer = &errorableWriter{back: w}
	err := dax.PrintYes()

	w.Close()
	var buf bytes.Buffer
	_, e := buf.ReadFrom(r)
	assert.NoError(t, e)
	assert.Equal(t, buf.String(), strings.Repeat("y\n", 5))

	switch err.Reason().(type) {
	case FailToPrint:
		assert.Equal(t, err.Cause().Error(), "write error")
	default:
		assert.Fail(t, err.Error())
	}
}

func TestYes_ConsoleDax_PrintWord(t *testing.T) {
	defer resetStdout()

	r, w, _ := os.Pipe()

	dax := NewConsoleDax()
	dax.writer = &errorableWriter{back: w}
	err := dax.PrintWord("hello")

	w.Close()
	var buf bytes.Buffer
	_, e := buf.ReadFrom(r)
	assert.NoError(t, e)
	assert.Equal(t, buf.String(), strings.Repeat("hello\n", 5))

	switch err.Reason().(type) {
	case FailToPrint:
		assert.Equal(t, err.Cause().Error(), "write error")
	default:
		assert.Fail(t, err.Error())
	}
}

func TestYes_ConsoleDax_PrintHelp(t *testing.T) {
	defer resetStdout()

	r, w, _ := os.Pipe()
	os.Stdout = w

	dax := NewConsoleDax()
	dax.PrintVersion()

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

func TestYes_ConsoleDax_PrintVersion(t *testing.T) {
	defer resetStdout()

	r, w, _ := os.Pipe()
	os.Stdout = w

	dax := NewConsoleDax()
	dax.PrintHelp()

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
