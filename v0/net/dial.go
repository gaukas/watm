package net

import (
	"github.com/gaukas/watm/wasip1"
)

// Dial dials a remote host for a network connection.
//
// In v0, network and address parameters are ignored, as the
// connection is essentially fully managed by the host.
func Dial(_, _ string) (Conn, error) {
	fd, err := wasip1.FromWATERErrCode(_import_host_dial())
	if err != nil {
		return nil, err
	}

	return &TCPConn{
		rawConn: &rawTCPConn{
			fd: fd,
		},
	}, nil
}
