package addon

// Imports a few functions from host

// Import from host.
//
//export send_nuke
func sendNuke(count int32)

// Import from host.
//
//export cancel_nuke
func cancelNuke() (errno int32)
