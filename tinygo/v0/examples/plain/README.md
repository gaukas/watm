# Example: `plain.wasm` 

This example shows how to build a minimal WATM with TinyGo which passes through the received/sent data without any processing.

## Build

Go 1.20/1.21 is required to build this example.

```bash
tinygo build -o plain.wasm -target=wasi .
```
