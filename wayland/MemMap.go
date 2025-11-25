package wayland

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/unix"
)

type MemMapInfo struct {
	Bytes          []byte
	Addr           unsafe.Pointer
	Size           int
	FileDescriptor int
	UnMapped       bool
}

func NewMemMapInfo(fd int, size uint64) (MemMapInfo, error) {
	if size == 0 {
		return MemMapInfo{FileDescriptor: fd, UnMapped: true}, fmt.Errorf("size must be > 0")
	}

	sizeInt := int(size)
	data, err := unix.Mmap(fd, 0, sizeInt, unix.PROT_READ|unix.PROT_WRITE, unix.MAP_SHARED)
	if err != nil {
		return MemMapInfo{
			Bytes:          nil,
			Addr:           nil,
			Size:           sizeInt,
			FileDescriptor: fd,
			UnMapped:       true,
		}, fmt.Errorf("failed to mmap fd %d: %w", fd, err)
	}

	var addr unsafe.Pointer
	if len(data) > 0 {
		addr = unsafe.Pointer(&data[0])
	}

	info := MemMapInfo{
		Bytes:          data,
		Addr:           addr,
		Size:           sizeInt,
		FileDescriptor: fd,
		UnMapped:       false,
	}
	return info, nil
}

func (m *MemMapInfo) Unmap() {
	if m.UnMapped {
		return
	}
	_ = unix.Munmap(m.Bytes)
	m.UnMapped = true
	m.Bytes = nil
	m.Addr = nil
}
