package lib_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/sttk-go/sabi-sample-gnu-coreutils/lib"
	"os"
	"testing"
)

/* On go test, dax.GetTtyName returns NOTTY error.
func TestGetTtyName(t *testing.T) {
  dax := lib.NewTtyDax()

  fd0 := int(os.Stdin.Fd())
  ttyname, err := dax.GetTtyName(fd0)
  t.Logf("ttyname(%v) = %v (%v)\n", fd0, ttyname, err)
  assert.True(t, err.IsOk())
  assert.Equal(t, ttyname[0:9], "/dev/ttys")
}
*/

func TestGetTtyName_ENOTTY(t *testing.T) {
	dax := lib.NewTtyDax()

	fd0 := int(os.Stdin.Fd())
	ttyname, err := dax.GetTtyName(fd0)
	t.Logf("ttyname(%v) = %v (%v)\n", fd0, ttyname, err)
	switch err.Reason().(type) {
	case lib.FailToGetTtyName:
		assert.Equal(t, err.Get("Errno"), 25)
	default:
		assert.Fail(t, err.Error())
	}
	assert.Equal(t, ttyname, "")
}

func TestGetTtyName_EBADF(t *testing.T) {
	dax := lib.NewTtyDax()

	fd := 99
	ttyname, err := dax.GetTtyName(fd)
	t.Logf("ttyname(%v) = %v (%v)\n", fd, ttyname, err)
	switch err.Reason().(type) {
	case lib.FailToGetTtyName:
		assert.Equal(t, err.Get("Errno"), 9)
	default:
		assert.Fail(t, err.Error())
	}
	assert.Equal(t, ttyname, "")
}
