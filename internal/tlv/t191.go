package tlv

import "github.com/sealdice/MiraiGo/binary"

func T191(k byte) []byte {
	return binary.NewWriterF(func(w *binary.Writer) {
		w.WriteUInt16(0x191)
		pos := w.FillUInt16()
		w.WriteByte(k)
		w.WriteUInt16At(pos, uint16(w.Len()-4))
	})
}
