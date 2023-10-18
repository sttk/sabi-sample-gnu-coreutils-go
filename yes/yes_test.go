package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sttk/sabi"
	"github.com/sttk/sabi/errs"
)

type mapDax struct {
	m map[string]any
}

func newMapDaxBase(m map[string]any) sabi.DaxBase {
	base := sabi.NewDaxBase()
	return struct {
		sabi.DaxBase
		mapDax
	}{
		DaxBase: base,
		mapDax:  mapDax{m: m},
	}
}

func (dax mapDax) GetMode() int {
	return dax.m["mode"].(int)
}

func (dax mapDax) GetWord() string {
	return dax.m["word"].(string)
}

func (dax mapDax) PrintYes() errs.Err {
	if dax.m["error"] != nil {
		return errs.New(FailToPrint{})
	}
	dax.m["print"] = "y"
	return errs.Ok()
}

func (dax mapDax) PrintWord(word string) errs.Err {
	if dax.m["error"] != nil {
		return errs.New(FailToPrint{})
	}
	dax.m["print"] = word
	return errs.Ok()
}

func (dax mapDax) PrintVersion() {
	dax.m["print"] = "VERSION"
}

func (dax mapDax) PrintHelp() {
	dax.m["print"] = "HELP"
}

func TestYes_YesLogic_modeIsNoWord(t *testing.T) {
	m := make(map[string]any)
	m["mode"] = MODE_NOWORD

	base := newMapDaxBase(m)
	assert.True(t, sabi.Txn(base, YesLogic).IsOk())

	assert.Equal(t, m["print"], "y")
}

func TestYes_YesLogic_modeIsWord(t *testing.T) {
	m := make(map[string]any)
	m["mode"] = MODE_WORD
	m["word"] = "hello"

	dax := newMapDaxBase(m).(YesDax)
	assert.True(t, YesLogic(dax).IsOk())

	assert.Equal(t, m["print"], "hello")
}

func TestYes_YesLogic_modeIsVersion(t *testing.T) {
	m := make(map[string]any)
	m["mode"] = MODE_VERSION

	dax := newMapDaxBase(m).(YesDax)
	assert.True(t, YesLogic(dax).IsOk())

	assert.Equal(t, m["print"], "VERSION")
}

func TestYes_YesLogic_modeIsHelp(t *testing.T) {
	m := make(map[string]any)
	m["mode"] = MODE_HELP

	dax := newMapDaxBase(m).(YesDax)
	assert.True(t, YesLogic(dax).IsOk())

	assert.Equal(t, m["print"], "HELP")
}
