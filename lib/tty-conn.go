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
	FailToGetTtyname struct{ Errno int }
)

type TtyConn struct {
}

func (conn *TtyConn) GetTtyname(fd int) (string, sabi.Err) {
	const BUF_SIZE int = 512

	buf := (*C.char)(C.malloc(C.size_t(C.sizeof_char * BUF_SIZE)))
	defer C.free(unsafe.Pointer(buf))

	errno := int(C.ttyname_r(C.int(fd), buf, C.size_t(BUF_SIZE)))
	if errno == 0 {
		return C.GoString(buf), sabi.Ok()
	}
	return "", sabi.ErrBy(FailToGetTtyname{Errno: errno})
	// errno = 9:EBADF | 19:ENODEV | 25:ENOTTY | 34:ERANGE
}

func (conn *TtyConn) Commit() sabi.Err {
	return sabi.Ok()
}

func (conn *TtyConn) Rollback() {
}

func (conn *TtyConn) Close() {
}

type ttyConnCfg struct {
}

func NewTtyConnCfg() sabi.ConnCfg {
	return ttyConnCfg{}
}

func (cfg ttyConnCfg) CreateConn() (sabi.Conn, sabi.Err) {
	return &TtyConn{}, sabi.Ok()
}
