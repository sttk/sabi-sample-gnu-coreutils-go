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

func (dax mapDax) GetMode() (int, sabi.Err) {
	switch dax.m["mode"].(int) {
	case MODE_ERROR:
		return MODE_ERROR, sabi.NewErr(InvalidOption{Option: "--opt"})
	default:
		return dax.m["mode"].(int), sabi.Ok()
	}
}

type TtyError struct{}

func (dax mapDax) GetStdinTtyName() (string, sabi.Err) {
	switch dax.m["error"] {
	case "notty":
		return "not a tty", sabi.NewErr(StdinIsNotTty{})
	case "ttyError":
		return "tty error", sabi.NewErr(TtyError{})
	default:
		return dax.m["ttyname"].(string), sabi.Ok()
	}
}

func (dax mapDax) PrintTtyName(ttyname string) sabi.Err {
	if dax.m["error"] == "failToPrint" {
		return sabi.NewErr(FailToPrint{})
	}
	dax.m["print"] = ttyname
	return sabi.Ok()
}

func (dax mapDax) PrintNotTty(err sabi.Err) {
	dax.m["print"] = err
}

func (dax mapDax) PrintTtyError(err sabi.Err) {
	dax.m["print"] = err
}

func (dax mapDax) PrintModeError(err sabi.Err) {
	dax.m["print"] = err
}

func (dax mapDax) PrintVersion() sabi.Err {
	dax.m["print"] = "VERSION"
	return sabi.Ok()
}

func (dax mapDax) PrintHelp() sabi.Err {
	dax.m["print"] = "HELP"
	return sabi.Ok()
}

func newTestProc(m map[string]any) sabi.Proc[TtyDax] {
	base := sabi.NewDaxBase()
	dax := newMapDax(m)
	return sabi.NewProc[TtyDax](base, dax)
}

func TestTtyLogic_if_mode_is_version(t *testing.T) {
	m := make(map[string]any)
	m["mode"] = MODE_VERSION

	proc := newTestProc(m)
	err := proc.RunTxn(TtyLogic)

	assert.Equal(t, m["print"], "VERSION")
	assert.True(t, err.IsOk())
}

func TestTtyLogic_if_mode_is_help(t *testing.T) {
	m := make(map[string]any)
	m["mode"] = MODE_HELP

	proc := newTestProc(m)
	err := proc.RunTxn(TtyLogic)

	assert.Equal(t, m["print"], "HELP")
	assert.True(t, err.IsOk())
}

func TestTtyLogic_if_mode_is_normal(t *testing.T) {
	m := make(map[string]any)
	m["mode"] = MODE_NORMAL
	m["ttyname"] = "/dev/tty001"

	proc := newTestProc(m)
	err := proc.RunTxn(TtyLogic)

	assert.Equal(t, m["print"], "/dev/tty001")
	assert.True(t, err.IsOk())
}

func TestTtyLogic_if_mode_is_silent(t *testing.T) {
	m := make(map[string]any)
	m["mode"] = MODE_SILENT
	m["ttyname"] = "/dev/tty001"

	proc := newTestProc(m)
	err := proc.RunTxn(TtyLogic)

	assert.Nil(t, m["print"])
	assert.True(t, err.IsOk())
}

func TestTtyLogic_if_mode_is_error(t *testing.T) {
	m := make(map[string]any)
	m["mode"] = MODE_ERROR

	proc := newTestProc(m)
	err := proc.RunTxn(TtyLogic)

	assert.Equal(t, m["print"], err)
	switch err.Reason().(type) {
	case InvalidOption:
		assert.Equal(t, err.Get("Option"), "--opt")
		assert.Equal(t, err.Reason().(InvalidOption).Option, "--opt")
	default:
		assert.Fail(t, err.Error())
	}
}

func TestTtyLogic_if_mode_is_normal_but_notty(t *testing.T) {
	m := make(map[string]any)
	m["mode"] = MODE_NORMAL
	m["error"] = "notty"

	proc := newTestProc(m)
	err := proc.RunTxn(TtyLogic)

	assert.Equal(t, m["print"], err)
	switch err.Reason().(type) {
	default:
		assert.Fail(t, err.Error())
	case StdinIsNotTty:
	}
}

func TestTtyLogic_if_mode_is_silent_but_notty(t *testing.T) {
	m := make(map[string]any)
	m["mode"] = MODE_SILENT
	m["error"] = "notty"

	proc := newTestProc(m)
	err := proc.RunTxn(TtyLogic)

	assert.Nil(t, m["print"])
	switch err.Reason().(type) {
	default:
		assert.Fail(t, err.Error())
	case StdinIsNotTty:
	}
}

func TestTtyLogic_if_mode_is_normal_but_tty_error(t *testing.T) {
	m := make(map[string]any)
	m["mode"] = MODE_NORMAL
	m["error"] = "ttyError"

	proc := newTestProc(m)
	err := proc.RunTxn(TtyLogic)

	assert.Equal(t, m["print"], err)
	switch err.Reason().(type) {
	default:
		assert.Fail(t, err.Error())
	case TtyError:
	}
}

func TestTtyLogic_if_mode_is_silent_but_tty_error(t *testing.T) {
	m := make(map[string]any)
	m["mode"] = MODE_SILENT
	m["error"] = "ttyError"

	proc := newTestProc(m)
	err := proc.RunTxn(TtyLogic)

	assert.Nil(t, m["print"])
	switch err.Reason().(type) {
	default:
		assert.Fail(t, err.Error())
	case TtyError:
	}
}

func TestTtyLogic_if_mode_is_normal_but_fail_to_write(t *testing.T) {
	m := make(map[string]any)
	m["mode"] = MODE_NORMAL
	m["ttyname"] = "/dev/tty001"
	m["error"] = "failToPrint"

	proc := newTestProc(m)
	err := proc.RunTxn(TtyLogic)

	assert.Nil(t, m["print"])
	switch err.Reason().(type) {
	default:
		assert.Fail(t, err.Error())
	case FailToPrint:
	}
}
