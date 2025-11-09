package wayland

import "github.com/mmulet/term.everything/wayland/protocols"

const (
	stateObjectID protocols.DecodeStateType = iota
	stateOpcode
	stateSize
	stateData
)

func initialDecodeState() protocols.DecodeState {
	return protocols.DecodeState{
		Phase:    stateObjectID,
		I:        0,
		ObjectID: 0,
		Opcode:   0,
		Size:     0,
		Data:     nil,
	}
}

type MessageDecoder struct {
	state protocols.DecodeState
}

func NewMessageDecoder() *MessageDecoder {
	return &MessageDecoder{
		state: initialDecodeState(),
	}
}

func (d *MessageDecoder) reset() {
	d.state = initialDecodeState()
}

func (d *MessageDecoder) nextState() {
	d.state.I = 0
	switch d.state.Phase {
	case stateObjectID:
		d.state.Phase = stateOpcode
		d.state.Opcode = 0
	case stateOpcode:
		d.state.Phase = stateSize
		d.state.Size = 0
	case stateSize:
		d.state.Phase = stateData
		d.state.Data = d.state.Data[:0]
		// 8 is the header size; no payload means return to initial state
		if d.state.Size == 8 {
			d.reset()
		}
	case stateData:
		d.reset()
	}
}

func (d *MessageDecoder) Consume(buf []byte) []protocols.Message {

	out := make([]protocols.Message, 0)
	for _, b := range buf {
		switch d.state.Phase {
		case stateObjectID:
			d.state.ObjectID |= protocols.AnyObjectID(b) << d.state.I
			d.state.I += 8
			if d.state.I == 32 {
				d.nextState()
			}
		case stateOpcode:
			d.state.Opcode |= uint16(b) << d.state.I
			d.state.I += 8
			if d.state.I == 16 {
				d.nextState()
			}
		case stateSize:
			d.state.Size |= uint16(b) << d.state.I
			d.state.I += 8
			if d.state.I == 16 {
				if d.state.Size == 8 {
					// zero-size payload message (header-only)
					out = append(out, protocols.Message{
						ObjectID: d.state.ObjectID,
						Opcode:   d.state.Opcode,
						Size:     d.state.Size,
						Data:     []byte{},
					})
				}
				d.nextState()
			}
		case stateData:
			d.state.Data = append(d.state.Data, b)
			// size includes 8-byte header, so payload length is size-8
			want := int(d.state.Size) - 8
			if len(d.state.Data) == want {
				// copy data to detach from internal buffer
				payload := make([]byte, want)
				copy(payload, d.state.Data)
				out = append(out, protocols.Message{
					ObjectID: d.state.ObjectID,
					Opcode:   d.state.Opcode,
					Size:     d.state.Size,
					Data:     payload,
				})
				d.nextState()
			}
		}
	}

	return out
}
