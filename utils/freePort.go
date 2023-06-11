package utils

import "net"

// GetFreeRandomPort returns a free port number on the host machine
func GetFreeRandomPort() int {
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	return ln.Addr().(*net.TCPAddr).Port
}
