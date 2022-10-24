package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/sttk-go/sabi"
	"github.com/sttk-go/sabi-sample-gnu-coreutils/lib"
	"reflect"
	"testing"
)

func TestTtynameDax_getTtyConn(t *testing.T) {
	base := sabi.NewConnBase()
	base.AddLocalConnCfg("ttyname", lib.NewTtyConnCfg())
	dax := newTtynameDax(base)

	conn, err := dax.getTtyConn("ttyname")
	assert.True(t, err.IsOk())
	switch ((interface{})(conn)).(type) {
	default:
		assert.Fail(t, fmt.Sprintf("%v", reflect.TypeOf(conn)))
	case *lib.TtyConn:
	}
}

func TestTtynameDax_getTtyConn_ConnCfgIsNotFound(t *testing.T) {
	base := sabi.NewConnBase()
	dax := newTtynameDax(base)

	conn, err := dax.getTtyConn("ttt")
	assert.Nil(t, conn)
	switch err.Reason().(type) {
	default:
		assert.Fail(t, err.Error())
	case sabi.ConnCfgIsNotFound:
	}
}

// On go test, `C.ttyname_r` returns `ENOTTY`
func TestTtyDax_GetStdioTtyname(t *testing.T) {
	base := sabi.NewConnBase()
	base.AddLocalConnCfg("ttyname", lib.NewTtyConnCfg())
	dax := newTtynameDax(base)

	ttyname, err := dax.GetStdinTtyname()
	assert.Equal(t, ttyname, "")
	switch err.Reason().(type) {
	default:
		assert.Fail(t, err.Error())
	case StdinIsNotTty:
	}
}
