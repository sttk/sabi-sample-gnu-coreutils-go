package lib

// #cgo CFLAGS: -g -Wall
// #include <unistd.h>
// #include <stdlib.h>
import "C"

import (
	"github.com/sttk-go/sabi"
	"unsafe"
)

type /* error reason */ (
	FailToGetTtyName struct{ Errno int }
)

type OsDaxSrc struct {
}

func NewOsDaxSrc() OsDaxSrc {
	return OsDaxSrc{}
}

func (ds OsDaxSrc) CreateDaxConn() (sabi.DaxConn, sabi.Err) {
	return &OsDaxConn{}, sabi.Ok()
}

type OsDaxConn struct {
	sabi.DaxConn
}

func (conn *OsDaxConn) Commit() sabi.Err {
	return sabi.Ok()
}

func (conn *OsDaxConn) Rollback() {
}

func (conn *OsDaxConn) Close() {
}

func (conn *OsDaxConn) GetTtyName(fd int) (string, sabi.Err) {
	const BUF_SIZE int = 512

	buf := (*C.char)(C.malloc(C.size_t(C.sizeof_char * BUF_SIZE)))
	defer C.free(unsafe.Pointer(buf))

	errno := int(C.ttyname_r(C.int(fd), buf, C.size_t(BUF_SIZE)))
	if errno == 0 {
		return C.GoString(buf), sabi.Ok()
	}

	// errno = 9:EBADF | 19:ENODEV | 25:ENOTTY | 34:ERANGE
	return "", sabi.NewErr(FailToGetTtyName{Errno: errno})
}

type OsDax struct {
	sabi.Dax
}

func NewOsDax(dax sabi.Dax) OsDax {
	return OsDax{Dax: dax}
}

func (dax OsDax) GetOsDaxConn(name string) (*OsDaxConn, sabi.Err) {
	conn, err := dax.GetDaxConn(name)
	return conn.(*OsDaxConn), err
}
