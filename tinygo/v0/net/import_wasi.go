//go:build wasip1

package net

// Import the host-imported dialer function.
//
//export host_dial
func _import_host_dial() (fd int32)

// Import the host-imported acceptor function.
//
//export host_accept
func _import_host_accept() (fd int32)
