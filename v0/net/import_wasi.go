//go:build wasip1

package net

// Import the host-imported dialer function.
//
//export host_dial
func _import_host_dial() (fd int32)
