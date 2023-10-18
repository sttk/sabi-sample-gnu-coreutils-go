package lib

// #cgo CFLAGS: -g -Wall
// #include <unistd.h>
// #include <stdlib.h>
import "C"

import (
	"unsafe"

	"github.com/sttk/sabi"
	"github.com/sttk/sabi/errs"
)

type (
	FailToGetTtyName struct{ Errno int }
)

type OsDaxSrc struct {
}

func NewOsDaxSrc() OsDaxSrc {
	return OsDaxSrc{}
}

func (ds OsDaxSrc) Setup(ag sabi.AsyncGroup) errs.Err {
	return errs.Ok()
}

func (ds OsDaxSrc) Close() {
}

func (ds OsDaxSrc) CreateDaxConn() (sabi.DaxConn, errs.Err) {
	return OsDaxConn{}, errs.Ok()
}

type OsDaxConn struct {
}

func (conn OsDaxConn) Commit(ag sabi.AsyncGroup) errs.Err {
	return errs.Ok()
}

func (conn OsDaxConn) IsCommitted() bool {
	return true
}

func (conn OsDaxConn) Rollback(ag sabi.AsyncGroup) {
	// never be run because IsCommitted always returns true.
}

func (conn OsDaxConn) ForceBack(ag sabi.AsyncGroup) {
}

func (conn OsDaxConn) Close() {
}

func (conn OsDaxConn) GetTtyName(fd int) (string, errs.Err) {
	const BUF_SIZE int = 512

	buf := (*C.char)(C.malloc(C.size_t(C.sizeof_char * BUF_SIZE)))
	defer C.free(unsafe.Pointer(buf))

	errno := int(C.ttyname_r(C.int(fd), buf, C.size_t(BUF_SIZE)))
	if errno == 0 {
		return C.GoString(buf), errs.Ok()
	}

	// errno = 9:EBADF | 19:ENODEV | 25:ENOTTY | 34:ERANGE
	return "", errs.New(FailToGetTtyName{Errno: errno})
}
