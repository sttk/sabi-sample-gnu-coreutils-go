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

func (dax mapDax) GetMode() (mode int, err errs.Err) {
	e := dax.m["error"]
	if e == "bad option" {
		o := dax.m["option"].(string)
		return MODE_NORMAL, errs.New(InvalidOption{Option: o})
	}
	return dax.m["mode"].(int), errs.Ok()
}

func (dax mapDax) GetStdinTtyName() (ttyName string, err errs.Err) {
	switch dax.m["error"] {
	case "notty":
		ttyName = "not a tty"
		err = errs.New(StdinIsNotTty{})
	default:
		ttyName = dax.m["ttyname"].(string)
		err = errs.Ok()
	}
	return
}

func (dax mapDax) PrintTtyName(ttyName string) errs.Err {
	if dax.m["error"] == "failToPrint" {
		return errs.New(FailToPrint{})
	}
	dax.m["print"] = ttyName
	return errs.Ok()
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

func TestTty_TtyLogic_modeIsNormal(t *testing.T) {
	m := make(map[string]any)
	m["mode"] = MODE_NORMAL
	m["ttyname"] = "abc"
	dax := newMapDaxBase(m)

	err := TtyLogic(dax.(TtyDax))
	assert.True(t, err.IsOk())

	assert.Equal(t, m["print"], "abc")
}

func TestTty_TtyLogic_modeIsNormal_butStdinIsNotTty(t *testing.T) {
	m := make(map[string]any)
	m["mode"] = MODE_NORMAL
	m["error"] = "notty"
	dax := newMapDaxBase(m)

	err := TtyLogic(dax.(TtyDax))
	assert.Equal(t, m["print"], "{reason=StdinIsNotTty}")

	switch err.Reason().(type) {
	case StdinIsNotTty:
	default:
		assert.Fail(t, err.Error())
	}
}

func TestTty_TtyLogic_modeIsSilent(t *testing.T) {
	m := make(map[string]any)
	m["mode"] = MODE_SILENT
	m["ttyname"] = "abc"
	dax := newMapDaxBase(m)

	err := TtyLogic(dax.(TtyDax))
	assert.Nil(t, m["print"])

	assert.True(t, err.IsOk())
}

func TestTty_TtyLogic_modeIsSilent_butStdinIsNotTty(t *testing.T) {
	m := make(map[string]any)
	m["mode"] = MODE_SILENT
	m["error"] = "notty"
	dax := newMapDaxBase(m)

	err := TtyLogic(dax.(TtyDax))
	assert.Nil(t, m["print"])

	switch err.Reason().(type) {
	case StdinIsNotTty:
	default:
		assert.Fail(t, err.Error())
	}
}

func TestTty_TtyLogic_modeIsHelp(t *testing.T) {
	m := make(map[string]any)
	m["mode"] = MODE_HELP
	dax := newMapDaxBase(m)

	err := TtyLogic(dax.(TtyDax))
	assert.Equal(t, m["print"], "HELP")

	assert.True(t, err.IsOk())
}

func TestTty_TtyLogic_modeIsVersion(t *testing.T) {
	m := make(map[string]any)
	m["mode"] = MODE_VERSION
	dax := newMapDaxBase(m)

	err := TtyLogic(dax.(TtyDax))
	assert.Equal(t, m["print"], "VERSION")

	assert.True(t, err.IsOk())
}

func TestTty_TtyLogic_invalidOption(t *testing.T) {
	m := make(map[string]any)
	m["error"] = "bad option"
	m["option"] = "-x"
	dax := newMapDaxBase(m)

	err := TtyLogic(dax.(TtyDax))
	assert.Equal(t, m["print"], "{reason=InvalidOption, Option=-x}")

	switch err.Reason().(type) {
	case InvalidOption:
	default:
		assert.Fail(t, err.Error())
	}
}
