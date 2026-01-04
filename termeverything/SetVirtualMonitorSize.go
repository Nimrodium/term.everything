package termeverything

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mmulet/term.everything/wayland"
)

func SetVirtualMonitorSize(newVirtualMonitorSize string) error {
	mkError := func(msg string, a ...any) error {
		errorMessage := fmt.Sprintf("Invalid virtual monitor size %s, expected <width>x<height>", newVirtualMonitorSize)
		return fmt.Errorf("%v; %v", errorMessage, fmt.Sprintf(msg, a...))
	}

	if newVirtualMonitorSize == "" {
		return nil
	}
	parts := strings.Split(newVirtualMonitorSize, "x")
	if len(parts) != 2 {
		return mkError("found less than 2 dimensions %v", parts)
	}
	width, widthErr := strconv.Atoi(parts[0])
	height, heightErr := strconv.Atoi(parts[1])
	if widthErr != nil {
		return mkError("invalid width %v; must be an integer", widthErr)
	}
	if heightErr != nil {
		return mkError("invalid height %v; must be an integer", heightErr)
	}
	if width <= 0 {
		return mkError("invalid width %v; must be greater than zero", width)
	}
	if height <= 0 {
		return mkError("invalid height %v; must be greater than zero", height)
	}
	wayland.VirtualMonitorSize.Width = wayland.Pixels(width)
	wayland.VirtualMonitorSize.Height = wayland.Pixels(height)
	return nil
}
