package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sttk/sabi"
	"github.com/sttk/sabi/errs"

	"github.com/sttk/sabi-sample-gnu-coreutils-go/lib"
)

// On go test, `C.ttyname_r` always returns `ENOTTY`.
func TestTty_OsUserDax_GetStdinTtyName(t *testing.T) {
	base := sabi.NewDaxBase()
	defer base.Close()

	assert.True(t, base.Uses("os", lib.OsDaxSrc{}).IsOk())

	dax := NewOsUserDax(base)
	ttynm, err := dax.GetStdinTtyName()

	assert.Equal(t, ttynm, "")
	switch err.Reason().(type) {
	default:
		assert.Fail(t, err.Error())
	case StdinIsNotTty:
	}
}

type MyDaxSrc struct{}

func (ds MyDaxSrc) Setup(ag sabi.AsyncGroup) errs.Err { return errs.Ok() }
func (ds MyDaxSrc) Close()                            {}
func (ds MyDaxSrc) CreateDaxConn() (sabi.DaxConn, errs.Err) {
	return nil, errs.Ok()
}

func TestTty_OsUserDax_GetStdinTtyName_failToGetDaxConn(t *testing.T) {
	base := sabi.NewDaxBase()
	defer base.Close()

	base.Uses("os", MyDaxSrc{})

	dax := NewOsUserDax(base)
	ttynm, err := dax.GetStdinTtyName()

	assert.Equal(t, ttynm, "")
	switch r := err.Reason().(type) {
	default:
		assert.Fail(t, err.Error())
	case sabi.CreatedDaxConnIsNil:
		assert.Equal(t, r.Name, "os")
	}
}
