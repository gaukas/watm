//go:build !wasip1

package net

var hostDialedFD int32 = -1

func setHostDialedFD(fd int32) {
	hostDialedFD = fd
}

// This function should be imported from the host in WASI.
// On non-WASI platforms, it mimicks the behavior of the host
// by returning a file descriptor of preset value.
func _import_host_dial() (fd int32) {
	return hostDialedFD
}
