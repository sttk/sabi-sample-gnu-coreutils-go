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

func (dax mapDax) GetMode() int {
	return dax.m["mode"].(int)
}

func (dax mapDax) GetEffectiveUserId() string {
	return dax.m["euid"].(string)
}

func (dax mapDax) GetUserNameByUserId(uid string) string {
	return dax.m["username"].(map[string]string)[uid]
}

func (dax mapDax) PrintUserName(userName string) {
	dax.m["print"] = userName
}

func (dax mapDax) PrintHelp() {
	dax.m["print"] = "HELP"
}

func (dax mapDax) PrintVersion() {
	dax.m["print"] = "VERSION"
}

func newTestProc(m map[string]any) sabi.Proc[WhoamiDax] {
	base := sabi.NewDaxBase()
	dax := newMapDax(m)
	return sabi.NewProc[WhoamiDax](base, dax)
}

func TestWhoamiLogic_if_mode_is_version(t *testing.T) {
	m := make(map[string]any)
	m["mode"] = MODE_VERSION

	proc := newTestProc(m)
	proc.RunTxn(WhoamiLogic)

	assert.Equal(t, m["print"], "VERSION")
}

func TestWhoamiLogic_if_mode_is_help(t *testing.T) {
	m := make(map[string]any)
	m["mode"] = MODE_HELP

	proc := newTestProc(m)
	proc.RunTxn(WhoamiLogic)

	assert.Equal(t, m["print"], "HELP")
}

func TestWhoamiLogic_if_mode_is_normal(t *testing.T) {
	users := make(map[string]string)
	users["111"] = "foo"
	users["123"] = "bar"

	m := make(map[string]any)
	m["mode"] = MODE_NORMAL
	m["euid"] = "123"
	m["username"] = users

	proc := newTestProc(m)
	proc.RunTxn(WhoamiLogic)

	assert.Equal(t, m["print"], "bar")
}
