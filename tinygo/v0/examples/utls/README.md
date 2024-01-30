# Example: `uTLS.wasm`

This example shows how to build a fully functional TLS client with TinyGo from [uTLS](https://github.com/refraction-networking/utls/tree/wasm). 

## Build

Go 1.20/1.21 is required to build this example.

```bash
tinygo build -o utls.wasm -target=wasi -tags=purego .
```

## Dependencies

In `go.mod`, you must replace `utls` with the branch `wasm` of [uTLS](https://github.com/refraction-networking/utls/tree/wasm). The best way is clone uTLS and checkout to the branch `wasm`, then replace the `utls` in `go.mod` with path to the cloned uTLS.