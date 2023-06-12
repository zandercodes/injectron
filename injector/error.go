package injector

import "errors"

// Create Error Type ErrExecutableNotFound with message "executable not found"
var ErrExecutableNotFound = errors.New("executable not found")

// Create Error Type ErrNoSocketsFound with message "no sockets found"
var ErrNoSocketsFound = errors.New("no sockets found")
