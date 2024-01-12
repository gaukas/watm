module github.com/gaukas/watm

go 1.21

// replace github.com/tetratelabs/wazero => ../wazero

replace github.com/tetratelabs/wazero v1.6.0 => github.com/gaukas/wazero v1.6.0-w

require github.com/tetratelabs/wazero v1.6.0
