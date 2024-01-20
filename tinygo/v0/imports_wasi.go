//go:build wasip1 || wasi

package v0

//go:wasmimport env host_defer
//go:noescape
func _import_host_defer()

//go:wasmimport env pull_config
//go:noescape
func _import_pull_config() (fd int32)
