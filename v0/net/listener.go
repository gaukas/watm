package net

type Listener interface {
	Accept() (Conn, error)
	SetNonBlock(nonblocking bool) error
}
