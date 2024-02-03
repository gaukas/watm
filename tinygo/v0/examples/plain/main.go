package main

import v0 "github.com/gaukas/watm/tinygo/v0"

func init() {
	v0.BuildDialerWithWrappingTransport(&PlainWrappingTransport{})
	v0.BuildListenerWithWrappingTransport(&PlainWrappingTransport{})
	v0.BuildRelayWithWrappingTransport(&PlainWrappingTransport{}, v0.RelayWrapRemote)
}

func main() {}
