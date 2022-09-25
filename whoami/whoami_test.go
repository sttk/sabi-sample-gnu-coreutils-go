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

func (dax mapDax) getEffectiveUserId() string {
	return dax.m["euid"].(string)
}

func (dax mapDax) getUsernameByUserId(uid string) string {
	return dax.m["username"].(map[string]string)[uid]
}

func (dax mapDax) printUsername(username string) {
	dax.m["print"] = username
}

func (dax mapDax) printVersion() {
	dax.m["print"] = "VERSION"
}

func (dax mapDax) printHelp() {
	dax.m["print"] = "HELP"
}

func newTestProc(m map[string]any) sabi.Proc[whoamiDax] {
	base := sabi.NewConnBase()
	dax := struct {
		mapDax
	}{
		mapDax: newMapDax(m),
	}
	return sabi.NewProc[whoamiDax](base, dax)
}

func TestWhoamiLogic_if_mode_is_version(t *testing.T) {
	m := make(map[string]any)
	m["mode"] = mode_version

	proc := newTestProc(m)
	proc.RunTxn(whoamiLogic)

	assert.Equal(t, m["print"], "VERSION")
}

func TestWhoamiLogic_if_mode_is_help(t *testing.T) {
	m := make(map[string]any)
	m["mode"] = mode_help

	proc := newTestProc(m)
	proc.RunTxn(whoamiLogic)

	assert.Equal(t, m["print"], "HELP")
}

func TestWhoamiLogic_if_mode_is_normal(t *testing.T) {
	m := make(map[string]any)
	m["mode"] = mode_normal
	m["euid"] = "123"
	m["username"] = make(map[string]string)
	m["username"].(map[string]string)["111"] = "foo"
	m["username"].(map[string]string)["123"] = "bar"

	proc := newTestProc(m)
	proc.RunTxn(whoamiLogic)

	assert.Equal(t, m["print"], "bar")
}
