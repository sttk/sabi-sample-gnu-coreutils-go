package lib_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/sttk-go/sabi-sample-gnu-coreutils/lib"
	"os"
	"testing"
)

/* On go test, TtyConn#GetTtyname returns NOTTY error.
func TestTtyConn_GetTtyname(t *testing.T) {
  cfg := lib.NewTtyConnCfg()
  conn, err := cfg.CreateConn()
  assert.True(t, err.IsOk())

  fd := int(os.Stdin.Fd())
  ttyname, err := conn.GetTtyname(fd)
  t.Logf("ttyname(%v) = %v (%v)\n", fd, ttyname, err)
  assert.True(t, err.IsOK())
  assert.Equal(t, ttyname[0:9], "/dev/ttys")
}
*/

func TestTtyConn_GetTtyname_ENOTTY(t *testing.T) {
	cfg := lib.NewTtyConnCfg()
	conn, err := cfg.CreateConn()
	assert.True(t, err.IsOk())

	fd := int(os.Stdin.Fd())
	ttyname, err := conn.(*lib.TtyConn).GetTtyname(fd)
	switch err.Reason().(type) {
	case lib.FailToGetTtyname:
		assert.Equal(t, err.Get("Errno"), lib.ENOTTY)
	default:
		assert.Fail(t, err.Error())
	}
	assert.Equal(t, ttyname, "")

	conn.Commit()
	conn.Close()
}

func TestTtyConn_GetTtyname_EBADF(t *testing.T) {
	cfg := lib.NewTtyConnCfg()
	conn, err := cfg.CreateConn()
	assert.True(t, err.IsOk())

	fd := 99
	ttyname, err := conn.(*lib.TtyConn).GetTtyname(fd)
	switch err.Reason().(type) {
	case lib.FailToGetTtyname:
		assert.Equal(t, err.Get("Errno"), lib.EBADF)
	default:
		assert.Fail(t, err.Error())
	}
	assert.Equal(t, ttyname, "")

	conn.Rollback()
	conn.Close()
}
