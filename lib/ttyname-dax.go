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

type TtyDax struct {
}

func NewTtyDax() TtyDax {
	return TtyDax{}
}

func (dax *TtyDax) GetTtyName(fd int) (string, sabi.Err) {
	const BUF_SIZE int = 512

	buf := (*C.char)(C.malloc(C.size_t(C.sizeof_char * BUF_SIZE)))
	defer C.free(unsafe.Pointer(buf))

	errno := int(C.ttyname_r(C.int(fd), buf, C.size_t(BUF_SIZE)))
	if errno == 0 {
		return C.GoString(buf), sabi.Ok()
	}

	// errno = 9:EBADF | 19:ENODEV | 25:ENOTTY | 34:ERANGE
	return "", sabi.ErrBy(FailToGetTtyName{Errno: errno})
}
