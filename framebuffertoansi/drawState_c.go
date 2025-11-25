package framebuffertoansi

// #cgo pkg-config: chafa glib-2.0
// #include "chafa.h"
import "C"

import (
	"fmt"
	"os"
	"strings"
	"unsafe"

	"github.com/mmulet/term.everything/escapecodes"
)

type drawState struct {
	SessionTypeIsX11 bool
	ChafaInfo        *chafaInfo
}

func (ds *drawState) isDrawState() {}

func MakeDrawState(sessionTypeIsX11 bool) DrawState {
	return &drawState{
		SessionTypeIsX11: sessionTypeIsX11,
	}
}

func (ds *drawState) ResizeChafaInfoIfNeeded(WidthCells int, HeightCells int, termSize TermSize) {

	if ds.ChafaInfo != nil && !(ds.ChafaInfo.WidthCells == WidthCells &&
		ds.ChafaInfo.HeightCells == HeightCells &&
		ds.ChafaInfo.WidthOfACellInPixels == termSize.WidthOfACellInPixels &&
		ds.ChafaInfo.HeightOfACellInPixels == termSize.HeightOfACellInPixels) {
		ds.ChafaInfo.destroy()
		ds.ChafaInfo = nil
	}
	if ds.ChafaInfo != nil {
		return
	}

	ds.ChafaInfo = makeChafaInfo(WidthCells,
		HeightCells,
		termSize.WidthOfACellInPixels,
		termSize.HeightOfACellInPixels,
		ds.SessionTypeIsX11)
}

func (ds *drawState) Destroy() {
	if ds.ChafaInfo != nil {
		ds.ChafaInfo.destroy()
		ds.ChafaInfo = nil
	}
}

func (ds *drawState) DrawDesktop(texturePixels []byte, width, height uint32, statusLine *string) (int, int) {
	haveStatusLine := statusLine != nil && len(*statusLine) > 0
	termSize := MakeTermSize()

	widthCells := termSize.WidthCells

	statusLineHeight := 0
	if haveStatusLine {
		statusLineHeight = 1
	}

	heightCells := termSize.HeightCells - statusLineHeight

	// Adjust geometry preserving aspect ratio.
	C.chafa_calc_canvas_geometry(
		C.int(width),
		C.int(height),
		(*C.int)(unsafe.Pointer(&widthCells)),
		(*C.int)(unsafe.Pointer(&heightCells)),
		C.gfloat(termSize.FontRatio),
		C.gboolean(1), // preserve aspect
		C.gboolean(0), // do not upscale
	)

	ds.ResizeChafaInfoIfNeeded(widthCells, heightCells, termSize)

	printable := ds.ChafaInfo.ConvertImage(texturePixels, width, height, width*4)

	var sb strings.Builder
	if haveStatusLine {
		sb.WriteString(escapecodes.MoveCursorToHome)
		sb.WriteString(*statusLine)
		sb.WriteString(escapecodes.ClearLineAfterCursor)
		sb.WriteString("\n")

	}
	sb.WriteString(printable)

	fmt.Fprint(os.Stdout, sb.String())
	_ = os.Stdout.Sync()

	return widthCells, heightCells
}
