//go:build !wasip1 && !wasi

package v0

import (
	"syscall"

	"github.com/gaukas/watm/wasip1"
)

func _import_host_defer() {
	// just do nothing, since nothing really matters if not
	// commanded by the host.
}

// emulate the behavior when no config is provided on
// the host side.
func _import_pull_config() (fd int32) {
	return wasip1.EncodeWATERError(syscall.ENOENT)
}
