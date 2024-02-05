module github.com/gaukas/watm

go 1.20

replace golang.org/x/sys v0.16.0 => ./replace/x/sys

replace github.com/tetratelabs/wazero v1.6.0 => github.com/gaukas/wazero v1.6.5-w

require (
	github.com/refraction-networking/utls v1.6.2-wasm
	github.com/tetratelabs/wazero v1.6.0
)

require (
	github.com/andybalholm/brotli v1.0.6 // indirect
	github.com/cloudflare/circl v1.3.7 // indirect
	github.com/klauspost/compress v1.17.4 // indirect
	github.com/quic-go/quic-go v0.40.1 // indirect
	golang.org/x/crypto v0.18.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
)
