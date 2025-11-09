package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mmulet/term.everything/server"
	"github.com/mmulet/term.everything/wayland"
)

//go:generate go generate ./wayland

// #cgo CFLAGS: -I./wayland/src -I./wayland/include
// #include "wayland.c"
// #include "xdg-shell.c"
// #include "xdg-decoration-unstable-v1.c"
// #include "xwayland-keyboard-grab-unstable-v1.c"
// #include "xwayland-shell-v1.c"
// #include "wayland.h"
// #include "xdg-shell.h"
// #include "xdg-decoration-unstable-v1.h"
// #include "xwayland-keyboard-grab-unstable-v1.h"
// #include "xwayland-shell-v1.h"

func main() {
	args := server.ParseArgs()
	server.SetVirtualMonitorSize(args.VirtualMonitorSize)
	listener, err := wayland.MakeSocketListener(&args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create socket listener: %v\n", err)
		os.Exit(1)
	}

	go func() {
		if err := listener.MainLoop(); err != nil {
			log.Printf("listener stopped: %v\n", err)
		}
	}()

	go func() {
		for conn := range listener.Connections {
			fmt.Println("New connection:", conn.RemoteAddr())
			// Example: immediately close via channel.
			listener.CloseConnections <- conn
		}
	}()

	// // Wait for SigInt, TODO something different
	// sig := make(chan os.Signal, 1)
	// signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	// <-sig
	// _ = listener.Close()
	// fmt.Println("Shutdown complete")
}
