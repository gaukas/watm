//go:build wasip1 || wasi

package main

import "github.com/gaukas/watm/tinygo/snippets/multi_package/addon"

func init() {
	addon.Login("gaukas")
}

func main() {
}
