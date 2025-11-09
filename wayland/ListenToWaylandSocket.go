package wayland

import (
	"fmt"
	"net"
)

func ListenToWaylandSocket(socketName string, socketPath string) (listner *net.UnixListener, fd int, e error) {

	if err := removeFileIfExists(socketPath); err != nil {
		return nil, -1, fmt.Errorf("remove existing socket: %w", err)
	}

	addr := &net.UnixAddr{
		Name: socketPath,
		Net:  "unix",
	}
	ln, err := net.ListenUnix("unix", addr)
	if err != nil {
		return nil, -1, fmt.Errorf("listen unix: %w", err)
	}

	file, err := ln.File()
	if err != nil {
		_ = ln.Close()
		return nil, -1, fmt.Errorf("get listener file: %w", err)
	}
	fd = int(file.Fd())
	_ = file.Close()

	return ln, fd, nil
}
