package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/sttk-go/sabi"
	"testing"
)

type mapDax struct {
	m map[string]any
}

func newTestProc(m map[string]any) sabi.Proc[YesDax] {
	base := sabi.NewDaxBase()
	dax := mapDax{m: m}
	return sabi.NewProc[YesDax](base, dax)
}

func (dax mapDax) GetMode() int {
	return dax.m["mode"].(int)
}

func (dax mapDax) GetWord() string {
	return dax.m["word"].(string)
}

func (dax mapDax) PrintYes() {
	dax.m["print"] = "y"
}

func (dax mapDax) PrintWord(word string) {
	dax.m["print"] = word
}

func (dax mapDax) PrintVersion() {
	dax.m["print"] = "VERSION"
}

func (dax mapDax) PrintHelp() {
	dax.m["print"] = "HELP"
}

func TestYesLogic_mode_is_noword(t *testing.T) {
	m := make(map[string]any)
	m["mode"] = MODE_NOWORD

	proc := newTestProc(m)
	proc.RunTxn(YesLogic)

	assert.Equal(t, m["print"], "y")
}

func TestYesLogic_mode_is_word(t *testing.T) {
	m := make(map[string]any)
	m["mode"] = MODE_WORD
	m["word"] = "hello"

	proc := newTestProc(m)
	proc.RunTxn(YesLogic)

	assert.Equal(t, m["print"], "hello")
}

func TestYesLogic_mode_is_version(t *testing.T) {
	m := make(map[string]any)
	m["mode"] = MODE_VERSION

	proc := newTestProc(m)
	proc.RunTxn(YesLogic)

	assert.Equal(t, m["print"], "VERSION")
}

func TestYesLogic_mode_is_help(t *testing.T) {
	m := make(map[string]any)
	m["mode"] = MODE_HELP

	proc := newTestProc(m)
	proc.RunTxn(YesLogic)

	assert.Equal(t, m["print"], "HELP")
}
