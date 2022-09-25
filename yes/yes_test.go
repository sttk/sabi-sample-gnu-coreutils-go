package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/sttk-go/sabi"
	"testing"
)

type mapDax struct {
	m map[string]any
}

func newMapDax(m map[string]any) mapDax {
	return mapDax{m: m}
}

func (dax mapDax) getMode() int {
	return dax.m["mode"].(int)
}
func (dax mapDax) getWord() string {
	return dax.m["word"].(string)
}
func (dax mapDax) printYes() {
	dax.m["print"] = "y"
}
func (dax mapDax) printWord() {
	dax.m["print"] = dax.m["word"]
}
func (dax mapDax) printVersion() {
	dax.m["print"] = "VERSION"
}
func (dax mapDax) printHelp() {
	dax.m["print"] = "HELP"
}

func newTestProc(m map[string]any) sabi.Proc[yesDax] {
	base := sabi.NewConnBase()
	dax := struct {
		mapDax
	}{
		mapDax: newMapDax(m),
	}
	return sabi.NewProc[yesDax](base, dax)
}

func TestYesLogic_mode_is_noword(t *testing.T) {
	m := make(map[string]any)
	m["mode"] = mode_noword

	proc := newTestProc(m)
	proc.RunTxn(yesLogic)

	assert.Equal(t, m["print"], "y")
}

func TestYesLogic_mode_is_word(t *testing.T) {
	m := make(map[string]any)
	m["mode"] = mode_word
	m["word"] = "hello"

	proc := newTestProc(m)
	proc.RunTxn(yesLogic)

	assert.Equal(t, m["print"], "hello")
}

func TestYesLogic_mode_is_version(t *testing.T) {
	m := make(map[string]any)
	m["mode"] = mode_version

	proc := newTestProc(m)
	proc.RunTxn(yesLogic)

	assert.Equal(t, m["print"], "VERSION")
}

func TestYesLogic_mode_is_help(t *testing.T) {
	m := make(map[string]any)
	m["mode"] = mode_help

	proc := newTestProc(m)
	proc.RunTxn(yesLogic)

	assert.Equal(t, m["print"], "HELP")
}
