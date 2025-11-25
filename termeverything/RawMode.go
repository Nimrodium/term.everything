package termeverything

import (
	"fmt"

	"golang.org/x/sys/unix"
)

func EnableRawModeFD(fd int) (func() error, error) {
	orig, err := unix.IoctlGetTermios(fd, unix.TCGETS)
	if err != nil {
		return func() error { return nil }, nil
	}

	raw := *orig

	raw.Iflag &^= (unix.IGNBRK | unix.BRKINT | unix.PARMRK | unix.ISTRIP | unix.INLCR | unix.IGNCR | unix.ICRNL | unix.IXON)
	raw.Oflag &^= unix.OPOST
	raw.Lflag &^= (unix.ECHO | unix.ECHONL | unix.ICANON | unix.ISIG | unix.IEXTEN)
	raw.Cflag &^= (unix.CSIZE | unix.PARENB)
	raw.Cflag |= unix.CS8

	raw.Cc[unix.VMIN] = 1
	raw.Cc[unix.VTIME] = 0

	// Preserve output post-processing (enable OPOST and ONLCR like the shell default).
	raw.Oflag |= unix.OPOST
	// // ONLCR exists on Linux; set it to preserve NL -> CRNL translation.
	// raw.Oflag |= unix.ONLCR

	if err := unix.IoctlSetTermios(fd, unix.TCSETS, &raw); err != nil {
		return nil, fmt.Errorf("tcsetattr (raw) failed: %w", err)
	}

	restored := false
	restore := func() error {
		if restored {
			return nil
		}
		restored = true
		if err := unix.IoctlSetTermios(fd, unix.TCSETS, orig); err != nil {
			return fmt.Errorf("tcsetattr (restore) failed: %w", err)
		}
		return nil
	}

	return restore, nil
}
