package network

import (
	"encoding/binary"
	"io"
	"net"
)

// -----------------------------------
// |flag(1)| len(3) |      data      |
// -----------------------------------
const _TcpPackMask = uint32(0xFFFFFF)
const _TcpPackHeadLength = int(4)

type StreamFilterTcpPack struct {
	LittleEndian bool
}

func (f *StreamFilterTcpPack) Read(c net.Conn, b []byte) (int, error) {
	const MsgMask = uint32(0xFFFFFF)

	headerBuf := make([]byte, _TcpPackHeadLength)
	if _, err := io.ReadFull(c, headerBuf); err != nil {
		return 0, err
	}
	header := uint32(0)
	if f.LittleEndian {
		header = binary.LittleEndian.Uint32(headerBuf)
	} else {
		header = binary.BigEndian.Uint32(headerBuf)
	}
	length := int(header & _TcpPackMask)
	if length >= MaxPackSize {
		return 0, ErrMessageOverBuf
	}
	if _, err := io.ReadFull(c, b[:length]); err != nil {
		return 0, err
	}
	return length, nil
}
func (f *StreamFilterTcpPack) Write(c net.Conn, b []byte) (int, error) {
	length := uint32(len(b))
	if length >= MaxPackSize {
		return 0, ErrMessageOverBuf
	}
	buf := make([]byte, int(length)+_TcpPackHeadLength)
	if f.LittleEndian {
		binary.LittleEndian.PutUint32(buf, length)
	} else {
		binary.BigEndian.PutUint32(buf, length)
	}
	copy(buf[_TcpPackHeadLength:], b)
	return c.Write(buf)
}
