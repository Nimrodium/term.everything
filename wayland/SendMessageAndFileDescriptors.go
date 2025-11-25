package wayland

import (
	"fmt"
	"net"

	"golang.org/x/sys/unix"
)

func SendMessageAndFileDescriptors(conn *net.UnixConn, buf []byte, fds []int) error {
	if len(buf) == 0 {
		return nil
	}

	total := 0
	var oobFirst []byte
	if len(fds) > 0 {
		oobFirst = unix.UnixRights(fds...)
	}

	for total < len(buf) {
		chunk := buf[total:]
		var oob []byte
		if total == 0 {
			oob = oobFirst // send FDs only once
		}

		n, _, err := conn.WriteMsgUnix(chunk, oob, nil)
		if err != nil {
			return err
		}
		if n <= 0 {
			return fmt.Errorf("WriteMsgUnix wrote %d bytes", n)
		}
		total += n
	}

	return nil
}
