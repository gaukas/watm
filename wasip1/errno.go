package wasip1

import (
	"fmt"
	"syscall"
)

// Errno is just a copy of syscall.Errno from the Go standard library.
//
// The values are defined in syscall/tables_wasip1.go.
type Errno int32

// FromSyscallErrno converts a syscall.Errno into a Errno on wasip1.
func FromSyscallErrno(errno syscall.Errno) Errno {
	if foundErrno, ok := mapSyscallErrno[errno]; ok {
		return foundErrno
	}
	panic(fmt.Sprintf("Errno %v(%d): %v(%d)", errno, errno, ENOSYS, ENOSYS))
}

func (e Errno) ToSyscallErrno() syscall.Errno {
	if foundSyscallErrno, ok := mapErrnoSyscall[e]; ok {
		return foundSyscallErrno
	}
	panic(fmt.Sprintf("Errno %v(%d): %v(%d)", e, e, syscall.ENOSYS, syscall.ENOSYS))
}

// Error returns the string representation of the error number.
func (e Errno) Error() string {
	if e == 0 {
		return "ESUCCESS"
	}

	if e < 0 || e >= Errno(len(errorstr)) {
		return fmt.Sprintf("unknown WATERErrno %d", e)
	}

	return errorstr[-e]
}

// WATERErrCode is defined as the negative value of Errno.
func FromWATERErrCode(errorCode int32) (n int32, err error) {
	if errorCode >= 0 {
		n = errorCode // so when error code is 0, it will return 0, nil
	} else if _, ok := mapErrnoSyscall[Errno(-errorCode)]; ok {
		// if the negative of the error code is a valid Errno, then it is a valid WATERErrno.
		err = Errno(-errorCode)
	} else {
		// otherwise, it is an unknown error code.
		err = Errno(ENOSYS)
	}
	return
}

func (e Errno) WATER() int32 {
	return int32(-e)
}
