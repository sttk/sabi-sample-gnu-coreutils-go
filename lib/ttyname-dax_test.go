package lib_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/sttk-go/sabi-sample-gnu-coreutils/lib"
	"os"
	"testing"
)

// On go test, TtyDax#GetTtyname(0) returns NOTTY error.
/*
func TestTtyDax_GetTtyName(t *testing.T) {
	dax := lib.NewTtyDax()

	fd := int(os.Stdin.Fd())
	ttyname, err := dax.GetTtyName(fd)
	t.Logf("ttyname(%v) = %v (%v)\n", fd, ttyname, err)

	assert.True(t, err.IsOk())
	assert.Equal(t, ttyname[0:9], "/dev/ttys")
}
*/

func TestTtyDax_GetTtyName_ENOTTY(t *testing.T) {
	dax := lib.NewTtyDax()

	fd := int(os.Stdin.Fd())
	ttyname, err := dax.GetTtyName(fd)

	switch err.Reason().(type) {
	case lib.FailToGetTtyName:
		assert.Equal(t, err.Get("Errno"), lib.ENOTTY)
		assert.Equal(t, err.Reason().(lib.FailToGetTtyName).Errno, lib.ENOTTY)
	default:
		assert.Fail(t, err.Error())
	}
	assert.Equal(t, ttyname, "")
}

func TestTtyDax_GetTtyName_EBADF(t *testing.T) {
	dax := lib.NewTtyDax()

	fd := 99
	ttyname, err := dax.GetTtyName(fd)

	switch err.Reason().(type) {
	case lib.FailToGetTtyName:
		assert.Equal(t, err.Get("Errno"), lib.EBADF)
		assert.Equal(t, err.Reason().(lib.FailToGetTtyName).Errno, lib.EBADF)
	default:
		assert.Fail(t, err.Error())
	}
	assert.Equal(t, ttyname, "")
}
