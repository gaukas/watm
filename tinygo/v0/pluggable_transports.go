package v0

import (
	v0net "github.com/gaukas/watm/tinygo/v0/net"
)

// WrappingTransport is the most basic transport type. It wraps
// a [v0net.Conn] into another [net.Conn] by providing some
// high-level application layer protocol.
type WrappingTransport interface {
	// Wrap wraps a [Conn] into another [net.Conn].
	//
	// The transport SHOULD provide non-blocking [net.Conn.Read]
	// operation on the returned [net.Conn] unless it is required
	// to block otherwise, e.g., by application protocol specifications
	// such as during a TLS handshake.
	//
	// The input [v0net.Conn] is not by default non-blocking. It is
	// the responsibility of the transport to make it non-blocking
	// if required by calling [v0net.Conn.SetNonblock].
	//
	// For the returned [v0net.Conn], it is highly recommended to
	// pass all funtions other than [v0net.Conn.Read] and [v0net.Conn.Write]
	// to the input [v0net.Conn] to reduce the complexity.
	Wrap(v0net.Conn) (v0net.Conn, error)
}

// DialingTransport is a transport type that can be used to dial
// a remote address and provide high-level application layer
// protocol over the dialed connection.
type DialingTransport interface {
	// SetDialer sets the dialer function that is used to dial
	// a remote address.
	//
	// In v0, the input parameter of the dialer function is
	// unused inside the WATM, given the connection is always
	// managed by the host application.
	//
	// The returned [v0net.Conn] is not by default non-blocking.
	// It is the responsibility of the transport to make it
	// non-blocking if required by calling [v0net.Conn.SetNonblock].
	SetDialer(dialer func(network, address string) (v0net.Conn, error))

	// Dial dials a remote address and returns a [net.Conn] that
	// provides high-level application layer protocol over the
	// dialed connection.
	//
	// The transport SHOULD provide non-blocking [v0net.Conn.Read]
	// operation on the returned [v0net.Conn] unless it is required
	// to block otherwise, e.g., by application protocol specifications
	// such as during a TLS handshake.
	//
	// For the returned [v0net.Conn], it is highly recommended to
	// pass all funtions other than [v0net.Conn.Read] and [v0net.Conn.Write]
	// to the [v0net.Conn] created by the underlying dialer function.
	Dial(network, address string) (v0net.Conn, error)
}

// ListeningTransport is a transport type that can be used to
// accept incoming connections on a local address and provide
// high-level application layer protocol over the accepted
// connection.
type ListeningTransport interface {
	// SetListener sets the listener that is used to accept
	// incoming connections.
	//
	// The returned [v0net.Conn] is not by default non-blocking.
	// It is the responsibility of the transport to make it
	// non-blocking if required by calling [v0net.Conn.SetNonblock].
	SetListener(listener v0net.Listener)

	// Accept accepts an incoming connection and returns a
	// [net.Conn] that provides high-level application layer
	// protocol over the accepted connection.
	//
	// The transport SHOULD provide non-blocking [v0net.Conn.Read]
	// operation on the returned [v0net.Conn] unless it is required
	// to block otherwise, e.g., by application protocol specifications
	// such as during a TLS handshake.
	//
	// For the returned [v0net.Conn], it is highly recommended to
	// pass all funtions other than [v0net.Conn.Read] and [v0net.Conn.Write]
	// to the [v0net.Conn] created by the underlying listener.
	Accept() (v0net.Conn, error)
}

type ConfigurableTransport interface {
	Configure(config []byte) error
}
