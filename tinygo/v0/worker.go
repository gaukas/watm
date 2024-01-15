package v0

import (
	"errors"
	"io"
	"log"
	"net"
	"syscall"

	"github.com/gaukas/watm/wasip1"
)

type identity uint8

var workerIdentity identity = identity_uninitialized

const (
	identity_uninitialized identity = iota
	identity_dialer
	identity_listener
	identity_relay
)

var identityStrings = map[identity]string{
	identity_dialer:   "dialer",
	identity_listener: "listener",
	identity_relay:    "relay",
}

var sourceConn net.Conn // sourceConn is used to communicate between WASM and the host application or a dialing party (for relay only)
var remoteConn net.Conn // remoteConn is used to communicate between WASM and a dialed remote destination (for dialer/relay) or a dialing party (for listener only)
var cancelConn net.Conn // cancelConn is used to cancel the entire worker.

var pollFn func() int32 = unfairPoll // by default, use unfairPoll

// PollingFairness sets the fairness of polling.
//
// If sourceConn or remoteConn will not work in non-blocking mode,
// it is highly recommended to set fair to true, otherwise it is most
// likely that the worker will block on reading from a blocking
// connection forever and therefore make no progress in the other
// direction.
func PollingFairness(fair bool) {
	if fair {
		pollFn = fairPoll
	} else {
		pollFn = unfairPoll
	}
}

//export _water_worker
func _water_worker() int32 {
	if workerIdentity == identity_uninitialized {
		log.Println("worker: uninitialized")
		return wasip1.EncodeWATERError(syscall.ENOTCONN) // socket not connected
	}
	log.Printf("worker: working as %s", identityStrings[workerIdentity])
	return poll()
}

func poll() int32 {
	defer _import_host_defer()

	if sourceConn == nil || remoteConn == nil || cancelConn == nil {
		log.Println("worker: unfairPoll: sourceConn, remoteConn, or cancelConn is nil")
		return wasip1.EncodeWATERError(syscall.EBADF) // bad file descriptor
	}

	return pollFn()
}

// untilError executes the given function until non-nil error is returned
func untilError(f func() error) error {
	var err error
	for err == nil {
		err = f()
	}
	return err
}

// unfairPoll works on all three connections with a priority order
// of cancelConn > sourceConn > remoteConn.
//
// It keeps working on the current connection until it returns an error,
// and if the error is EAGAIN, it switches to the next connection. If the
// connection is not properly set to non-blocking mode, i.e., never returns
// EAGAIN, this function will block forever and never work on a lower priority
// connection. Thus it is called unfairPoll.
func unfairPoll() int32 {
	var readBuf []byte = make([]byte, 65536)
	for {
		// first priority: cancelConn
		_, err := cancelConn.Read(readBuf)
		if !errors.Is(err, syscall.EAGAIN) {
			if errors.Is(err, io.EOF) || err == nil {
				log.Println("worker: unfairPoll: cancelConn is closed")
				return wasip1.EncodeWATERError(syscall.ECANCELED) // operation canceled
			}
			log.Println("worker: unfairPoll: cancelConn.Read:", err)
			return wasip1.EncodeWATERError(syscall.EIO) // input/output error
		}

		// second priority: sourceConn
		if err := untilError(func() error {
			readN, readErr := sourceConn.Read(readBuf)
			if readErr != nil {
				return readErr
			}

			writeN, writeErr := remoteConn.Write(readBuf[:readN])
			if writeErr != nil {
				log.Println("worker: unfairPoll: remoteConn.Write:", writeErr)
				return syscall.EIO // input/output error, we cannot retry async write yet
			}

			if readN != writeN {
				log.Println("worker: unfairPoll: readN != writeN")
				return syscall.EIO // input/output error
			}

			return nil
		}); !errors.Is(err, syscall.EAGAIN) {
			if errors.Is(err, io.EOF) {
				log.Println("worker: unfairPoll: sourceConn is closed")
				return wasip1.EncodeWATERError(syscall.EPIPE) // broken pipe
			}
			log.Println("worker: unfairPoll: sourceConn.Read:", err)
			return wasip1.EncodeWATERError(syscall.EIO) // input/output error
		}

		// third priority: remoteConn
		if err := untilError(func() error {
			readN, readErr := remoteConn.Read(readBuf)
			if readErr != nil {
				return readErr
			}

			writeN, writeErr := sourceConn.Write(readBuf[:readN])
			if writeErr != nil {
				log.Println("worker: unfairPoll: sourceConn.Write:", writeErr)
				return syscall.EIO // input/output error, we cannot retry async write yet
			}

			if readN != writeN {
				log.Println("worker: unfairPoll: readN != writeN")
				return syscall.EIO // input/output error
			}

			return nil
		}); !errors.Is(err, syscall.EAGAIN) {
			if errors.Is(err, io.EOF) {
				log.Println("worker: unfairPoll: remoteConn is closed")
				return wasip1.EncodeWATERError(syscall.EPIPE) // broken pipe
			}
			log.Println("worker: unfairPoll: remoteConn.Read:", err)
			return wasip1.EncodeWATERError(syscall.EIO) // input/output error
		}
	}
}

// like unfairPoll, fairPoll also works on all three connections with a priority order
// of cancelConn > sourceConn > remoteConn.
//
// But different from unfairPoll, fairPoll spend equal amount of turns on each connection
// for calling Read. Therefore it has a better fairness than unfairPoll, which may still
// make progress if one of the connection is not properly set to non-blocking mode.
func fairPoll() int32 {
	var readBuf []byte = make([]byte, 65536)
	for {
		// first priority: cancelConn
		_, err := cancelConn.Read(readBuf)
		if !errors.Is(err, syscall.EAGAIN) {
			if errors.Is(err, io.EOF) || err == nil {
				log.Println("worker: unfairPoll: cancelConn is closed")
				return wasip1.EncodeWATERError(syscall.ECANCELED) // operation canceled
			}
			log.Println("worker: unfairPoll: cancelConn.Read:", err)
			return wasip1.EncodeWATERError(syscall.EIO) // input/output error
		}

		// second priority: sourceConn -> remoteConn
		if err := copyOnce(
			"remoteConn", // dstName
			"sourceConn", // srcName
			remoteConn,   // dst
			sourceConn,   // src
			readBuf); err != nil {
			return wasip1.EncodeWATERError(err.(syscall.Errno))
		}

		// third priority: remoteConn -> sourceConn
		if err := copyOnce(
			"sourceConn", // dstName
			"remoteConn", // srcName
			sourceConn,   // dst
			remoteConn,   // src
			readBuf); err != nil {
			return wasip1.EncodeWATERError(err.(syscall.Errno))
		}
	}
}

func copyOnce(dstName, srcName string, dst, src net.Conn, buf []byte) error {
	if buf == nil {
		buf = make([]byte, 65536)
	}

	readN, readErr := src.Read(buf)
	if !errors.Is(readErr, syscall.EAGAIN) { // if EAGAIN, do nothing and return
		if errors.Is(readErr, io.EOF) {
			return syscall.EPIPE // broken pipe
		} else if readErr != nil {
			log.Printf("worker: copyOnce: %s.Read: %v", srcName, readErr)
			return syscall.EIO // input/output error
		}

		writeN, writeErr := dst.Write(buf[:readN])
		if writeErr != nil {
			log.Printf("worker: copyOnce: %s.Write: %v", dstName, writeErr)
			return syscall.EIO // no matter input/output error or EAGAIN we cannot retry async write yet
		}

		if readN != writeN {
			log.Printf("worker: copyOnce: %s.read != %s.written", srcName, dstName)
			return syscall.EIO // input/output error
		}
	}

	return nil
}
