// GENERATED BY events_encoding_gen.go
// DO NOT MODIFY MANUALLY
package trace

func NewEventByCode(code byte) Event {
	switch code {
	case 0x00:
		return &Invalid{}
	case 0x01:
		return &StreamStart{}
	case 0x02:
		return &StreamStop{}
	case 0x03:
		return &ThreadStart{}
	case 0x04:
		return &ThreadSleep{}
	case 0x05:
		return &ThreadWake{}
	case 0x06:
		return &ThreadStop{}
	case 0x07:
		return &SpanBegin{}
	case 0x08:
		return &SpanEnd{}
	case 0x09:
		return &Start{}
	case 0x0A:
		return &Finish{}
	case 0x0C:
		return &Snapshot{}
	case 0x0D:
		return &Info{}
	}
	panic("unknown code")
}

// Invalid 0x00
func (ev *Invalid) Code() byte { return 0x00 }
func (ev *Invalid) Decode(dec *Decoder) {

}
func (ev *Invalid) Encode(enc *Encoder) {

}

// StreamStart 0x01
func (ev *StreamStart) Code() byte { return 0x01 }
func (ev *StreamStart) Decode(dec *Decoder) {
	ev.ProcessID = dec.readID()
	ev.MachineID = dec.readID()
	ev.Time = dec.readTime()
	ev.Freq = dec.readFreq()
}
func (ev *StreamStart) Encode(enc *Encoder) {
	enc.writeID(ev.ProcessID)
	enc.writeID(ev.MachineID)
	enc.writeTime(ev.Time)
	enc.writeFreq(ev.Freq)
}

// StreamStop 0x02
func (ev *StreamStop) Code() byte { return 0x02 }
func (ev *StreamStop) Decode(dec *Decoder) {
	ev.Time = dec.readTime()
}
func (ev *StreamStop) Encode(enc *Encoder) {
	enc.writeTime(ev.Time)
}

// ThreadStart 0x03
func (ev *ThreadStart) Code() byte { return 0x03 }
func (ev *ThreadStart) Decode(dec *Decoder) {
	ev.Time = dec.readTime()
	ev.ThreadID = dec.readID()
	ev.StackID = dec.readID()
}
func (ev *ThreadStart) Encode(enc *Encoder) {
	enc.writeTime(ev.Time)
	enc.writeID(ev.ThreadID)
	enc.writeID(ev.StackID)
}

// ThreadSleep 0x04
func (ev *ThreadSleep) Code() byte { return 0x04 }
func (ev *ThreadSleep) Decode(dec *Decoder) {
	ev.Time = dec.readTime()
	ev.ThreadID = dec.readID()
	ev.StackID = dec.readID()
}
func (ev *ThreadSleep) Encode(enc *Encoder) {
	enc.writeTime(ev.Time)
	enc.writeID(ev.ThreadID)
	enc.writeID(ev.StackID)
}

// ThreadWake 0x05
func (ev *ThreadWake) Code() byte { return 0x05 }
func (ev *ThreadWake) Decode(dec *Decoder) {
	ev.Time = dec.readTime()
	ev.ThreadID = dec.readID()
	ev.StackID = dec.readID()
}
func (ev *ThreadWake) Encode(enc *Encoder) {
	enc.writeTime(ev.Time)
	enc.writeID(ev.ThreadID)
	enc.writeID(ev.StackID)
}

// ThreadStop 0x06
func (ev *ThreadStop) Code() byte { return 0x06 }
func (ev *ThreadStop) Decode(dec *Decoder) {
	ev.Time = dec.readTime()
	ev.ThreadID = dec.readID()
	ev.StackID = dec.readID()
}
func (ev *ThreadStop) Encode(enc *Encoder) {
	enc.writeTime(ev.Time)
	enc.writeID(ev.ThreadID)
	enc.writeID(ev.StackID)
}

// SpanBegin 0x07
func (ev *SpanBegin) Code() byte { return 0x07 }
func (ev *SpanBegin) Decode(dec *Decoder) {
	ev.Time = dec.readTime()
	ev.ThreadID = dec.readID()
	ev.StackID = dec.readID()
	ev.ID = dec.readID()
}
func (ev *SpanBegin) Encode(enc *Encoder) {
	enc.writeTime(ev.Time)
	enc.writeID(ev.ThreadID)
	enc.writeID(ev.StackID)
	enc.writeID(ev.ID)
}

// SpanEnd 0x08
func (ev *SpanEnd) Code() byte { return 0x08 }
func (ev *SpanEnd) Decode(dec *Decoder) {
	ev.Time = dec.readTime()
	ev.ThreadID = dec.readID()
	ev.StackID = dec.readID()
	ev.ID = dec.readID()
}
func (ev *SpanEnd) Encode(enc *Encoder) {
	enc.writeTime(ev.Time)
	enc.writeID(ev.ThreadID)
	enc.writeID(ev.StackID)
	enc.writeID(ev.ID)
}

// Start 0x09
func (ev *Start) Code() byte { return 0x09 }
func (ev *Start) Decode(dec *Decoder) {
	ev.Time = dec.readTime()
	ev.ThreadID = dec.readID()
	ev.StackID = dec.readID()
	ev.ID = dec.readID()
}
func (ev *Start) Encode(enc *Encoder) {
	enc.writeTime(ev.Time)
	enc.writeID(ev.ThreadID)
	enc.writeID(ev.StackID)
	enc.writeID(ev.ID)
}

// Finish 0x0A
func (ev *Finish) Code() byte { return 0x0A }
func (ev *Finish) Decode(dec *Decoder) {
	ev.Time = dec.readTime()
	ev.ThreadID = dec.readID()
	ev.StackID = dec.readID()
	ev.ID = dec.readID()
}
func (ev *Finish) Encode(enc *Encoder) {
	enc.writeTime(ev.Time)
	enc.writeID(ev.ThreadID)
	enc.writeID(ev.StackID)
	enc.writeID(ev.ID)
}

// Snapshot 0x0C
func (ev *Snapshot) Code() byte { return 0x0C }
func (ev *Snapshot) Decode(dec *Decoder) {
	ev.Time = dec.readTime()
	ev.ThreadID = dec.readID()
	ev.StackID = dec.readID()
	ev.ID = dec.readID()
	ev.Kind = dec.readKind()
	ev.Content = dec.readBytes()
}
func (ev *Snapshot) Encode(enc *Encoder) {
	enc.writeTime(ev.Time)
	enc.writeID(ev.ThreadID)
	enc.writeID(ev.StackID)
	enc.writeID(ev.ID)
	enc.writeKind(ev.Kind)
	enc.writeBytes(ev.Content)
}

// Info 0x0D
func (ev *Info) Code() byte { return 0x0D }
func (ev *Info) Decode(dec *Decoder) {
	ev.ID = dec.readID()
	ev.Name = dec.readString()
	ev.Kind = dec.readKind()
	ev.Content = dec.readBytes()
}
func (ev *Info) Encode(enc *Encoder) {
	enc.writeID(ev.ID)
	enc.writeString(ev.Name)
	enc.writeKind(ev.Kind)
	enc.writeBytes(ev.Content)
}