module github.com/gaukas/watm/tinygo/v0/examples/utls

go 1.20

replace github.com/tetratelabs/wazero v1.6.0 => github.com/gaukas/wazero v1.6.5-w

replace github.com/refraction-networking/utls v1.6.2 => ../../../../../../refraction/utls

replace golang.org/x/sys v0.16.0 => ../../../../x/sys

require (
	github.com/gaukas/watm v0.0.0-20240129210620-96dc8ca14ab5
	github.com/refraction-networking/utls v1.6.2
)

require (
	github.com/andybalholm/brotli v1.0.6 // indirect
	github.com/cloudflare/circl v1.3.7 // indirect
	github.com/klauspost/compress v1.17.4 // indirect
	github.com/quic-go/quic-go v0.40.1 // indirect
	golang.org/x/crypto v0.18.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
)
