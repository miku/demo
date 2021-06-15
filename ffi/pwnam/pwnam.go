// https://man7.org/linux/man-pages/man3/getpwnam.3.html
package pwnam

/*
#cgo LDFLAGS: -lpcap
#include <sys/types.h>
#include <pwd.h>
#include <stdlib.h>
*/
import "C" // "could not determine kind of name for" if there's an empty line above

import "unsafe"

type Passwd struct {
	Uid   uint32
	Gid   uint32
	Dir   string
	Shell string
}

func Getpwnam(name string) *Passwd {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cpw := C.getpwnam(cname)
	return &Passwd{
		Uid:   uint32(cpw.pw_uid),
		Gid:   uint32(cpw.pw_uid),
		Dir:   C.GoString(cpw.pw_dir),
		Shell: C.GoString(cpw.pw_shell),
	}
}
