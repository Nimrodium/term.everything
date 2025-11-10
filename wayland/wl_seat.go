package wayland

import (
	"github.com/mmulet/term.everything/wayland/protocols"
)

type WlSeat struct {
	Version uint32
}

func (w *WlSeat) WlSeat_get_pointer(
	s protocols.ClientState,
	_ protocols.ObjectID[protocols.WlSeat],
	id protocols.ObjectID[protocols.WlPointer],
) {
	s.AddGlobalWlPointerBind(id, protocols.Version(w.Version))
	AddObject(s, id, Global_WlPointer)
}

func (w *WlSeat) WlSeat_get_keyboard(
	s protocols.ClientState,
	_ protocols.ObjectID[protocols.WlSeat],
	id protocols.ObjectID[protocols.WlKeyboard],
) {
	s.AddGlobalWlKeyboardBind(id, protocols.Version(w.Version))
	AddObject(s, id, Global_WlKeyboard)
	Global_WlKeyboard.Delegate.AfterGetKeyboard(s, id)
}

func (w *WlSeat) WlSeat_get_touch(
	s protocols.ClientState,
	object_id protocols.ObjectID[protocols.WlSeat],
	_ protocols.ObjectID[protocols.WlTouch],
) {
	SendError(s, object_id, protocols.WlSeatError_enum_missing_capability, "no touch")
}

func (w *WlSeat) WlSeat_release(
	_ protocols.ClientState,
	_ protocols.ObjectID[protocols.WlSeat],
) bool {
	return true
}

func (w *WlSeat) OnBind(
	s protocols.ClientState,
	_ protocols.AnyObjectID,
	_ string,
	newIdAny protocols.AnyObjectID,
	version uint32,
) {
	w.Version = version
	newID := protocols.ObjectID[protocols.WlSeat](newIdAny)

	// w.capabilities(s, new_id, wl_seat_capability.pointer | wl_seat_capability.keyboard);
	protocols.WlSeat_capabilities(
		s,
		newID,
		protocols.WlSeatCapability_enum_pointer|protocols.WlSeatCapability_enum_keyboard,
	)
	protocols.WlSeat_name(s, version, newID, "seat0")
}

func MakeWLSeat() *protocols.WlSeat {
	return &protocols.WlSeat{
		Delegate: &WlSeat{
			Version: 1,
		},
	}
}
