package main

// Import from host.
//
//export hello
func hello()

// main is required for the `wasi` target, even if it isn't used.
func main() {
	hello()
}
