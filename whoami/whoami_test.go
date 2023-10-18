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

func (dax mapDax) GetMode() (int, errs.Err) {
	e := dax.m["error"]
	mode := dax.m["mode"].(int)
	if e != nil {
		return mode, errs.New(InvalidOption{Option: e.(string)})
	}
	return mode, errs.Ok()
}

func (dax mapDax) GetEffectiveUserId() string {
	return dax.m["euid"].(string)
}

func (dax mapDax) GetUserNameByUserId(uid string) (string, errs.Err) {
	a := dax.m["username"]
	if a != nil {
		users, ok := a.(map[string]string)
		if ok {
			unm, ok := users[uid]
			if ok {
				return unm, errs.Ok()
			}
		}
	}
	return "", errs.New(FailToGetUserName{Uid: uid})
}

func (dax mapDax) PrintUserName(userName string) {
	dax.m["print"] = userName
}

func (dax mapDax) PrintErr(err errs.Err) {
	dax.m["print"] = err.Error()
}

func (dax mapDax) PrintHelp() {
	dax.m["print"] = "HELP"
}

func (dax mapDax) PrintVersion() {
	dax.m["print"] = "VERSION"
}

func TestWhoami_WhoamiLogic_modeIsNormal(t *testing.T) {
	users := make(map[string]string)
	users["111"] = "foo"
	users["123"] = "bar"

	m := make(map[string]any)
	m["mode"] = MODE_NORMAL
	m["euid"] = "123"
	m["username"] = users

	dax := newMapDaxBase(m).(WhoamiDax)
	assert.True(t, WhoamiLogic(dax).IsOk())

	assert.Equal(t, m["print"], "bar")
}

func TestWhoami_WhoamiLogic_modeIsNormal_failToGetUserName(t *testing.T) {
	users := make(map[string]string)
	users["111"] = "foo"
	users["123"] = "bar"

	m := make(map[string]any)
	m["mode"] = MODE_NORMAL
	m["euid"] = "xxx"
	m["username"] = users

	dax := newMapDaxBase(m).(WhoamiDax)
	err := WhoamiLogic(dax)

	switch r := err.Reason().(type) {
	case FailToGetUserName:
		assert.Equal(t, r.Uid, "xxx")
	default:
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, m["print"], "{reason=FailToGetUserName, Uid=xxx}")
}

func TestWhoami_WhoamiLogic_modeIsHelp(t *testing.T) {
	m := make(map[string]any)
	m["mode"] = MODE_HELP

	dax := newMapDaxBase(m).(WhoamiDax)
	assert.True(t, WhoamiLogic(dax).IsOk())

	assert.Equal(t, m["print"], "HELP")
}

func TestWhoami_WhoamiLogic_modeIsVersion(t *testing.T) {
	m := make(map[string]any)
	m["mode"] = MODE_VERSION

	dax := newMapDaxBase(m).(WhoamiDax)
	assert.True(t, WhoamiLogic(dax).IsOk())

	assert.Equal(t, m["print"], "VERSION")
}

func TestWhoami_WhoamiLogic_invalidOption(t *testing.T) {
	m := make(map[string]any)
	m["error"] = "--option"
	m["mode"] = MODE_NORMAL

	dax := newMapDaxBase(m).(WhoamiDax)
	err := WhoamiLogic(dax)

	switch r := err.Reason().(type) {
	case InvalidOption:
		assert.Equal(t, r.Option, "--option")
	default:
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, m["print"], "{reason=InvalidOption, Option=--option}")
}
