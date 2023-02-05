package lib_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/sttk-go/sabi"
	"github.com/sttk-go/sabi-sample-gnu-coreutils/lib"
	"os"
	"reflect"
	"testing"
)

func TestNewOsDaxSrc(t *testing.T) {
	ds := lib.NewOsDaxSrc()
	assert.NotNil(t, ds)
}

func TestOsDaxSrc_CreateDaxConn(t *testing.T) {
	ds := lib.NewOsDaxSrc()
	conn, err := ds.CreateDaxConn()
	assert.True(t, err.IsOk())
	assert.NotNil(t, conn)
}

// On go test, OsDaxConn#GetTtyName(0) returns NOTTTY error.
/*
func TestOsDaxConn_GetTtyName(t *testing.T) {
  ds := lib.NewOsDaxSrc()
	conn := ds.CreateDaxConn().(*lib.OsDaxConn)

	fd := int(os.Stdin.Fd())
	ttyname, err := conn.GetTtyName(fd)
	t.Logf("ttyname(%v) = %v (%v)\n", fd, ttyname, err)

	assert.True(t, err.IsOk())
	assert.Equal(t, ttyname[0:9], "/dev/ttys")
}
*/

func TestOsDaxConn_GetTtyName_ENOTTY(t *testing.T) {
	ds := lib.NewOsDaxSrc()
	conn0, err := ds.CreateDaxConn()
	assert.True(t, err.IsOk())

	conn := conn0.(*lib.OsDaxConn)

	fd := int(os.Stdin.Fd())
	ttyname, err := conn.GetTtyName(fd)

	switch err.Reason().(type) {
	case lib.FailToGetTtyName:
		assert.Equal(t, err.Get("Errno"), lib.ENOTTY)
		assert.Equal(t, err.Reason().(lib.FailToGetTtyName).Errno, lib.ENOTTY)
	default:
		assert.Fail(t, err.Error())
	}
	t.Logf("ttyname(%v) = %v (%v)\n", fd, ttyname, err)

	assert.Equal(t, ttyname, "")
}

func TestOsDax_GetOsDaxConn(t *testing.T) {
	base := sabi.NewDaxBase()
	base.AddLocalDaxSrc("os", lib.NewOsDaxSrc())

	dax := lib.NewOsDax(base)
	conn, err := dax.GetOsDaxConn("os")
	assert.True(t, err.IsOk())
	assert.Equal(t, reflect.TypeOf(conn).Elem().Name(), "OsDaxConn")
}
