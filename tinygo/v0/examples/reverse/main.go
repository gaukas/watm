package main

import v0 "github.com/gaukas/watm/tinygo/v0"

func init() {
	v0.BuildDialerWithWrappingTransport(&ReverseWrappingTransport{})
	v0.BuildListenerWithWrappingTransport(&ReverseWrappingTransport{})
	v0.BuildRelayWithWrappingTransport(&ReverseWrappingTransport{}, false)
}

func main() {}
