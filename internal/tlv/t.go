package tlv

import "github.com/sealdice/MiraiGo/binary"

func T(tag uint16, value []byte) []byte {
	return binary.NewWriterF(func(w *binary.Writer) {
		w.WriteUInt16(tag)
		w.WriteBytesShort(value)
	})
}
