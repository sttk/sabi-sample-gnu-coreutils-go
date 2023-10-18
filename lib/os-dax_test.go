package lib_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sttk/sabi-sample-gnu-coreutils-go/lib"
	"github.com/sttk/sabi/errs"
)

type noopAsyncGroup struct{}

func (ag *noopAsyncGroup) Add(fn func() errs.Err) {}

func TestNewOsDaxSrc(t *testing.T) {
	ds := lib.NewOsDaxSrc()
	assert.NotNil(t, ds)

	// for coverage
	ag := &noopAsyncGroup{}
	assert.True(t, ds.Setup(ag).IsOk())
	ds.Close()
}

func TestOsDaxSrc_CreateDaxConn(t *testing.T) {
	ds := lib.NewOsDaxSrc()
	conn, err := ds.CreateDaxConn()
	assert.True(t, err.IsOk())
	assert.IsType(t, conn, lib.OsDaxConn{})

	// for coverage
	ag := &noopAsyncGroup{}
	assert.True(t, conn.Commit(ag).IsOk())
	assert.True(t, conn.IsCommitted())
	conn.Rollback(ag)
	conn.ForceBack(ag)
	conn.Close()
}

// On `go test`, OsDaxConn#GetTtyName(0/1/2) always returns NOTTY error.
/*
func TestOsDaxConn_GetTtyName(t *testing.T) {
  ds := lib.NewOsDaxSrc()
  dc, err := ds.CreateDaxConn()
  assert.True(t, err.IsOk())
  conn, ok := dc.(lib.OsDaxConn)
  assert.True(t, ok)

  fd := int(os.Stdin.Fd())
  ttyname, err := conn.GetTtyName(fd)
  t.Logf("ttyname(%v) = %v (err=%v)\n", fd, ttyname, err)

  assert.True(t, err.IsOk())
  assert.Equal(t, ttyname[0:9], "/dev/ttys")
}
*/

func TestOsDaxConn_GetTtyName_ENOTTY(t *testing.T) {
	ds := lib.NewOsDaxSrc()
	dc, err := ds.CreateDaxConn()
	assert.True(t, err.IsOk())
	conn, ok := dc.(lib.OsDaxConn)
	assert.True(t, ok)

	fd := int(os.Stdin.Fd())
	_, err = conn.GetTtyName(fd)

	switch r := err.Reason().(type) {
	case lib.FailToGetTtyName:
		assert.Equal(t, r.Errno, lib.ENOTTY)
	default:
		assert.Fail(t, err.Error())
	}
}

func TestOsDaxConn_GetTtyName_EBADF(t *testing.T) {
	ds := lib.NewOsDaxSrc()
	dc, err := ds.CreateDaxConn()
	assert.True(t, err.IsOk())
	conn, ok := dc.(lib.OsDaxConn)
	assert.True(t, ok)

	fd := 300
	_, err = conn.GetTtyName(fd)

	switch r := err.Reason().(type) {
	case lib.FailToGetTtyName:
		assert.Equal(t, r.Errno, lib.EBADF)
	default:
		assert.Fail(t, err.Error())
	}
}
