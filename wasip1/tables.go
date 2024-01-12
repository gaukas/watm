package wasip1

import (
	"syscall"
)

const (
	E2BIG           Errno = 1
	EACCES          Errno = 2
	EADDRINUSE      Errno = 3
	EADDRNOTAVAIL   Errno = 4
	EAFNOSUPPORT    Errno = 5
	EAGAIN          Errno = 6
	EALREADY        Errno = 7
	EBADF           Errno = 8
	EBADMSG         Errno = 9
	EBUSY           Errno = 10
	ECANCELED       Errno = 11
	ECHILD          Errno = 12
	ECONNABORTED    Errno = 13
	ECONNREFUSED    Errno = 14
	ECONNRESET      Errno = 15
	EDEADLK         Errno = 16
	EDESTADDRREQ    Errno = 17
	EDOM            Errno = 18
	EDQUOT          Errno = 19
	EEXIST          Errno = 20
	EFAULT          Errno = 21
	EFBIG           Errno = 22
	EHOSTUNREACH    Errno = 23
	EIDRM           Errno = 24
	EILSEQ          Errno = 25
	EINPROGRESS     Errno = 26
	EINTR           Errno = 27
	EINVAL          Errno = 28
	EIO             Errno = 29
	EISCONN         Errno = 30
	EISDIR          Errno = 31
	ELOOP           Errno = 32
	EMFILE          Errno = 33
	EMLINK          Errno = 34
	EMSGSIZE        Errno = 35
	EMULTIHOP       Errno = 36
	ENAMETOOLONG    Errno = 37
	ENETDOWN        Errno = 38
	ENETRESET       Errno = 39
	ENETUNREACH     Errno = 40
	ENFILE          Errno = 41
	ENOBUFS         Errno = 42
	ENODEV          Errno = 43
	ENOENT          Errno = 44
	ENOEXEC         Errno = 45
	ENOLCK          Errno = 46
	ENOLINK         Errno = 47
	ENOMEM          Errno = 48
	ENOMSG          Errno = 49
	ENOPROTOOPT     Errno = 50
	ENOSPC          Errno = 51
	ENOSYS          Errno = 52
	ENOTCONN        Errno = 53
	ENOTDIR         Errno = 54
	ENOTEMPTY       Errno = 55
	ENOTRECOVERABLE Errno = 56
	ENOTSOCK        Errno = 57
	ENOTSUP         Errno = 58
	ENOTTY          Errno = 59
	ENXIO           Errno = 60
	EOVERFLOW       Errno = 61
	EOWNERDEAD      Errno = 62
	EPERM           Errno = 63
	EPIPE           Errno = 64
	EPROTO          Errno = 65
	EPROTONOSUPPORT Errno = 66
	EPROTOTYPE      Errno = 67
	ERANGE          Errno = 68
	EROFS           Errno = 69
	ESPIPE          Errno = 70
	ESRCH           Errno = 71
	ESTALE          Errno = 72
	ETIMEDOUT       Errno = 73
	ETXTBSY         Errno = 74
	EXDEV           Errno = 75
	ENOTCAPABLE     Errno = 76
	// needed by src/net/error_unix_test.go
	EOPNOTSUPP = ENOTSUP
)

// TODO: Auto-generate some day. (Hard-coded in binaries so not likely to change.)
var errorstr = [...]string{
	E2BIG:           "Argument list too long",
	EACCES:          "Permission denied",
	EADDRINUSE:      "Address already in use",
	EADDRNOTAVAIL:   "Address not available",
	EAFNOSUPPORT:    "Address family not supported by protocol family",
	EAGAIN:          "Try again",
	EALREADY:        "Socket already connected",
	EBADF:           "Bad file number",
	EBADMSG:         "Trying to read unreadable message",
	EBUSY:           "Device or resource busy",
	ECANCELED:       "Operation canceled.",
	ECHILD:          "No child processes",
	ECONNABORTED:    "Connection aborted",
	ECONNREFUSED:    "Connection refused",
	ECONNRESET:      "Connection reset by peer",
	EDEADLK:         "Deadlock condition",
	EDESTADDRREQ:    "Destination address required",
	EDOM:            "Math arg out of domain of func",
	EDQUOT:          "Quota exceeded",
	EEXIST:          "File exists",
	EFAULT:          "Bad address",
	EFBIG:           "File too large",
	EHOSTUNREACH:    "Host is unreachable",
	EIDRM:           "Identifier removed",
	EILSEQ:          "EILSEQ",
	EINPROGRESS:     "Connection already in progress",
	EINTR:           "Interrupted system call",
	EINVAL:          "Invalid argument",
	EIO:             "I/O error",
	EISCONN:         "Socket is already connected",
	EISDIR:          "Is a directory",
	ELOOP:           "Too many symbolic links",
	EMFILE:          "Too many open files",
	EMLINK:          "Too many links",
	EMSGSIZE:        "Message too long",
	EMULTIHOP:       "Multihop attempted",
	ENAMETOOLONG:    "File name too long",
	ENETDOWN:        "Network interface is not configured",
	ENETRESET:       "Network dropped connection on reset",
	ENETUNREACH:     "Network is unreachable",
	ENFILE:          "File table overflow",
	ENOBUFS:         "No buffer space available",
	ENODEV:          "No such device",
	ENOENT:          "No such file or directory",
	ENOEXEC:         "Exec format error",
	ENOLCK:          "No record locks available",
	ENOLINK:         "The link has been severed",
	ENOMEM:          "Out of memory",
	ENOMSG:          "No message of desired type",
	ENOPROTOOPT:     "Protocol not available",
	ENOSPC:          "No space left on device",
	ENOSYS:          "Not implemented on wasip1",
	ENOTCONN:        "Socket is not connected",
	ENOTDIR:         "Not a directory",
	ENOTEMPTY:       "Directory not empty",
	ENOTRECOVERABLE: "State not recoverable",
	ENOTSOCK:        "Socket operation on non-socket",
	ENOTSUP:         "Not supported",
	ENOTTY:          "Not a typewriter",
	ENXIO:           "No such device or address",
	EOVERFLOW:       "Value too large for defined data type",
	EOWNERDEAD:      "Owner died",
	EPERM:           "Operation not permitted",
	EPIPE:           "Broken pipe",
	EPROTO:          "Protocol error",
	EPROTONOSUPPORT: "Unknown protocol",
	EPROTOTYPE:      "Protocol wrong type for socket",
	ERANGE:          "Math result not representable",
	EROFS:           "Read-only file system",
	ESPIPE:          "Illegal seek",
	ESRCH:           "No such process",
	ESTALE:          "Stale file handle",
	ETIMEDOUT:       "Connection timed out",
	ETXTBSY:         "Text file busy",
	EXDEV:           "Cross-device link",
	ENOTCAPABLE:     "Capabilities insufficient",
}

var mapSyscallErrno = map[syscall.Errno]Errno{
	syscall.E2BIG:           E2BIG,
	syscall.EACCES:          EACCES,
	syscall.EADDRINUSE:      EADDRINUSE,
	syscall.EADDRNOTAVAIL:   EADDRNOTAVAIL,
	syscall.EAFNOSUPPORT:    EAFNOSUPPORT,
	syscall.EAGAIN:          EAGAIN,
	syscall.EALREADY:        EALREADY,
	syscall.EBADF:           EBADF,
	syscall.EBADMSG:         EBADMSG,
	syscall.EBUSY:           EBUSY,
	syscall.ECANCELED:       ECANCELED,
	syscall.ECHILD:          ECHILD,
	syscall.ECONNABORTED:    ECONNABORTED,
	syscall.ECONNREFUSED:    ECONNREFUSED,
	syscall.ECONNRESET:      ECONNRESET,
	syscall.EDEADLK:         EDEADLK,
	syscall.EDESTADDRREQ:    EDESTADDRREQ,
	syscall.EDOM:            EDOM,
	syscall.EDQUOT:          EDQUOT,
	syscall.EEXIST:          EEXIST,
	syscall.EFAULT:          EFAULT,
	syscall.EFBIG:           EFBIG,
	syscall.EHOSTUNREACH:    EHOSTUNREACH,
	syscall.EIDRM:           EIDRM,
	syscall.EILSEQ:          EILSEQ,
	syscall.EINPROGRESS:     EINPROGRESS,
	syscall.EINTR:           EINTR,
	syscall.EINVAL:          EINVAL,
	syscall.EIO:             EIO,
	syscall.EISCONN:         EISCONN,
	syscall.EISDIR:          EISDIR,
	syscall.ELOOP:           ELOOP,
	syscall.EMFILE:          EMFILE,
	syscall.EMLINK:          EMLINK,
	syscall.EMSGSIZE:        EMSGSIZE,
	syscall.EMULTIHOP:       EMULTIHOP,
	syscall.ENAMETOOLONG:    ENAMETOOLONG,
	syscall.ENETDOWN:        ENETDOWN,
	syscall.ENETRESET:       ENETRESET,
	syscall.ENETUNREACH:     ENETUNREACH,
	syscall.ENFILE:          ENFILE,
	syscall.ENOBUFS:         ENOBUFS,
	syscall.ENODEV:          ENODEV,
	syscall.ENOENT:          ENOENT,
	syscall.ENOEXEC:         ENOEXEC,
	syscall.ENOLCK:          ENOLCK,
	syscall.ENOLINK:         ENOLINK,
	syscall.ENOMEM:          ENOMEM,
	syscall.ENOMSG:          ENOMSG,
	syscall.ENOPROTOOPT:     ENOPROTOOPT,
	syscall.ENOSPC:          ENOSPC,
	syscall.ENOSYS:          ENOSYS,
	syscall.ENOTCONN:        ENOTCONN,
	syscall.ENOTDIR:         ENOTDIR,
	syscall.ENOTEMPTY:       ENOTEMPTY,
	syscall.ENOTRECOVERABLE: ENOTRECOVERABLE,
	syscall.ENOTSOCK:        ENOTSOCK,
	syscall.ENOTSUP:         ENOTSUP,
	syscall.ENOTTY:          ENOTTY,
	syscall.ENXIO:           ENXIO,
	syscall.EOVERFLOW:       EOVERFLOW,
	syscall.EOWNERDEAD:      EOWNERDEAD,
	syscall.EPERM:           EPERM,
	syscall.EPIPE:           EPIPE,
	syscall.EPROTO:          EPROTO,
	syscall.EPROTONOSUPPORT: EPROTONOSUPPORT,
	syscall.EPROTOTYPE:      EPROTOTYPE,
	syscall.ERANGE:          ERANGE,
	syscall.EROFS:           EROFS,
	syscall.ESPIPE:          ESPIPE,
	syscall.ESRCH:           ESRCH,
	syscall.ESTALE:          ESTALE,
	syscall.ETIMEDOUT:       ETIMEDOUT,
	syscall.ETXTBSY:         ETXTBSY,
	syscall.EXDEV:           EXDEV,
}

var mapErrnoSyscall = map[Errno]syscall.Errno{
	E2BIG:           syscall.E2BIG,
	EACCES:          syscall.EACCES,
	EADDRINUSE:      syscall.EADDRINUSE,
	EADDRNOTAVAIL:   syscall.EADDRNOTAVAIL,
	EAFNOSUPPORT:    syscall.EAFNOSUPPORT,
	EAGAIN:          syscall.EAGAIN,
	EALREADY:        syscall.EALREADY,
	EBADF:           syscall.EBADF,
	EBADMSG:         syscall.EBADMSG,
	EBUSY:           syscall.EBUSY,
	ECANCELED:       syscall.ECANCELED,
	ECHILD:          syscall.ECHILD,
	ECONNABORTED:    syscall.ECONNABORTED,
	ECONNREFUSED:    syscall.ECONNREFUSED,
	ECONNRESET:      syscall.ECONNRESET,
	EDEADLK:         syscall.EDEADLK,
	EDESTADDRREQ:    syscall.EDESTADDRREQ,
	EDOM:            syscall.EDOM,
	EDQUOT:          syscall.EDQUOT,
	EEXIST:          syscall.EEXIST,
	EFAULT:          syscall.EFAULT,
	EFBIG:           syscall.EFBIG,
	EHOSTUNREACH:    syscall.EHOSTUNREACH,
	EIDRM:           syscall.EIDRM,
	EILSEQ:          syscall.EILSEQ,
	EINPROGRESS:     syscall.EINPROGRESS,
	EINTR:           syscall.EINTR,
	EINVAL:          syscall.EINVAL,
	EIO:             syscall.EIO,
	EISCONN:         syscall.EISCONN,
	EISDIR:          syscall.EISDIR,
	ELOOP:           syscall.ELOOP,
	EMFILE:          syscall.EMFILE,
	EMLINK:          syscall.EMLINK,
	EMSGSIZE:        syscall.EMSGSIZE,
	EMULTIHOP:       syscall.EMULTIHOP,
	ENAMETOOLONG:    syscall.ENAMETOOLONG,
	ENETDOWN:        syscall.ENETDOWN,
	ENETRESET:       syscall.ENETRESET,
	ENETUNREACH:     syscall.ENETUNREACH,
	ENFILE:          syscall.ENFILE,
	ENOBUFS:         syscall.ENOBUFS,
	ENODEV:          syscall.ENODEV,
	ENOENT:          syscall.ENOENT,
	ENOEXEC:         syscall.ENOEXEC,
	ENOLCK:          syscall.ENOLCK,
	ENOLINK:         syscall.ENOLINK,
	ENOMEM:          syscall.ENOMEM,
	ENOMSG:          syscall.ENOMSG,
	ENOPROTOOPT:     syscall.ENOPROTOOPT,
	ENOSPC:          syscall.ENOSPC,
	ENOSYS:          syscall.ENOSYS,
	ENOTCONN:        syscall.ENOTCONN,
	ENOTDIR:         syscall.ENOTDIR,
	ENOTEMPTY:       syscall.ENOTEMPTY,
	ENOTRECOVERABLE: syscall.ENOTRECOVERABLE,
	ENOTSOCK:        syscall.ENOTSOCK,
	ENOTSUP:         syscall.ENOTSUP,
	ENOTTY:          syscall.ENOTTY,
	ENXIO:           syscall.ENXIO,
	EOVERFLOW:       syscall.EOVERFLOW,
	EOWNERDEAD:      syscall.EOWNERDEAD,
	EPERM:           syscall.EPERM,
	EPIPE:           syscall.EPIPE,
	EPROTO:          syscall.EPROTO,
	EPROTONOSUPPORT: syscall.EPROTONOSUPPORT,
	EPROTOTYPE:      syscall.EPROTOTYPE,
	ERANGE:          syscall.ERANGE,
	EROFS:           syscall.EROFS,
	ESPIPE:          syscall.ESPIPE,
	ESRCH:           syscall.ESRCH,
	ESTALE:          syscall.ESTALE,
	ETIMEDOUT:       syscall.ETIMEDOUT,
	ETXTBSY:         syscall.ETXTBSY,
	EXDEV:           syscall.EXDEV,
}
