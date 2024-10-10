package tlv

import "github.com/sealdice/MiraiGo/binary"

func T516() []byte {
	return binary.NewWriterF(func(w *binary.Writer) {
		w.WriteUInt16(0x516)
		w.WriteUInt16(4) // len of uint32
		w.WriteUInt32(0)
	})
}
