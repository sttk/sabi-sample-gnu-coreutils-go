package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/sttk-go/sabi"
	"testing"
)

// On go test, `C.ttyname_r` returns `ENOTTY`
func TestOsUserDax_GetStdinTtyName(t *testing.T) {
	base := sabi.NewDaxBase()
	dax := NewOsUserDax(base)

	ttyname, err := dax.GetStdinTtyName()
	assert.Equal(t, ttyname, "")
	switch err.Reason().(type) {
	default:
		assert.Fail(t, err.Error())
	case StdinIsNotTty:
	}
}
