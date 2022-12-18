package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// On go test, `C.ttyname_r` returns `ENOTTY`
func TestTtyNameDax_GetStdinTtyName(t *testing.T) {
	dax := NewTtyNameDax()

	ttyname, err := dax.GetStdinTtyName()
	assert.Equal(t, ttyname, "")
	switch err.Reason().(type) {
	default:
		assert.Fail(t, err.Error())
	case StdinIsNotTty:
	}
}
