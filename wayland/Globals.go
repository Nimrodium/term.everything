package wayland

import "github.com/mmulet/term.everything/wayland/protocols"

var Global_WlDisplay = MakeWLDisplay()
var Global_WlOutput = MakeWlOutput()
var Global_WlSeat = MakeWLSeat()
var Global_WlShm = MakeWlShm()
var Global_WlCompositor = MakeWlCompositor()
var Global_WlSubcompositor = MakeWlSubcompositor()
var Global_XdgWmBase = MakeXdgWmBase()
var Global_WlDataDeviceManager = MakeWlDataDeviceManager()
var Global_WlKeyboard = MakeWlKeyboard()
var Global_WlPointer = &protocols.WlPointer{
	Delegate: &Pointer,
}
var Global_ZwpXwaylandKeyboardGrabManagerV1 = MakeZwpXwaylandKeyboardGrabManagerV1()
var Global_XwaylandShellV1 = MakeXwaylandShellV1()
var Global_WlDataDevice = Global_WlSeat

var Global_WlTouch = MakeWlTouch()

var Global_ZxdgDecorationManagerV1 = MakeZxdgDecorationManagerV1()
