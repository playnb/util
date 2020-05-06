package util

import "encoding/binary"

func StringFromC(d []byte) string {
	i := 0
	for ; i < len(d); i++ {
		if d[i] == 0 {
			break
		}
	}
	if i == 0 {
		return ""
	} else {
		return string(d[0:i])
	}
}

func NewPackTool(bigEndian bool) *PackTool {
	p := &PackTool{
		IsBigEndian: bigEndian,
	}
	return p
}

type CanPack interface {
	Pack([]byte) int
	Unpack([]byte) int
}

type PackTool struct {
	IsBigEndian bool
}

func (p *PackTool) UnpackByte(val *byte, data []byte) int {
	*val = data[0]
	return 1
}
func (p *PackTool) UnpackUint16(val *uint16, data []byte) int {
	if p.IsBigEndian {
		*val = binary.BigEndian.Uint16(data)
	} else {
		*val = binary.LittleEndian.Uint16(data)
	}
	return 2
}
func (p *PackTool) UnpackInt16(val *int16, data []byte) int {
	if p.IsBigEndian {
		*val = int16(binary.BigEndian.Uint16(data))
	} else {
		*val = int16(binary.LittleEndian.Uint16(data))
	}
	return 2
}
func (p *PackTool) UnpackUint32(val *uint32, data []byte) int {
	if p.IsBigEndian {
		*val = binary.BigEndian.Uint32(data)
	} else {
		*val = binary.LittleEndian.Uint32(data)
	}
	return 4
}
func (p *PackTool) UnpackUint64(val *uint64, data []byte) int {
	if p.IsBigEndian {
		*val = binary.BigEndian.Uint64(data)
	} else {
		*val = binary.LittleEndian.Uint64(data)
	}
	return 8
}
func (p *PackTool) UnpackBool(val *bool, data []byte) int {
	*val = data[0] != 0
	return 1
}
func (p *PackTool) UnpackString(val *string, data []byte, size int) int {
	l := 0
	for ; l < size; l++ {
		if data[l] == 0 {
			break
		}
	}
	*val = string(data[:l])
	return size
}
func (p *PackTool) UnpackSlice(val *[]byte, data []byte, size int) int {
	if *val == nil {
		*val = make([]byte, size)
	}
	copy(*val, data[:size])
	return size
}
func (p *PackTool) UnpackArrayByte(val []byte, data []byte) int {
	copy(val, data)
	return len(val)
}
func (p *PackTool) UnpackSliceUint16(val *[]uint16, data []byte, size int) int {
	if *val == nil {
		*val = make([]uint16, size)
	}
	offset := 0
	for i := 0; i < size; i++ {
		p.UnpackUint16(&((*val)[i]), data[offset:])
	}
	return offset
}
func (p *PackTool) UnpackArrayUint16(val []uint16, data []byte) int {
	offset := 0
	for i := 0; i < len(val); i++ {
		p.UnpackUint16(&(val[i]), data[offset:])
	}
	return offset
}
func (p *PackTool) UnpackSliceUint64(val *[]uint64, data []byte, size int) int {
	if *val == nil {
		*val = make([]uint64, size)
	}
	offset := 0
	for i := 0; i < size; i++ {
		p.UnpackUint64(&((*val)[i]), data[offset:])
	}
	return offset
}
func (p *PackTool) UnpackArrayUint64(val []uint64, data []byte) int {
	offset := 0
	for i := 0; i < len(val); i++ {
		p.UnpackUint64(&(val[i]), data[offset:])
	}
	return offset
}

func (p *PackTool) PackByte(data []byte, val byte) int {
	data[0] = val
	return 1
}
func (p *PackTool) PackUint16(data []byte, val uint16) int {
	if p.IsBigEndian {
		binary.BigEndian.PutUint16(data, val)
	} else {
		binary.LittleEndian.PutUint16(data, val)
	}
	return 2
}
func (p *PackTool) PackUint32(data []byte, val uint32) int {
	if p.IsBigEndian {
		binary.BigEndian.PutUint32(data, val)
	} else {
		binary.LittleEndian.PutUint32(data, val)
	}
	return 4
}
func (p *PackTool) PackUint64(data []byte, val uint64) int {
	if p.IsBigEndian {
		binary.BigEndian.PutUint64(data, val)
	} else {
		binary.LittleEndian.PutUint64(data, val)
	}
	return 8
}
func (p *PackTool) PackBool(data []byte, val bool) int {
	if val {
		data[0] = 1
	} else {
		data[0] = 0
	}
	return 1
}
func (p *PackTool) PackString(data []byte, val string, size int) int {
	copy(data[0:size], []byte(val))
	return size
}
func (p *PackTool) PackSlice(data []byte, val []byte, size int) int {
	if val != nil && size > 0 {
		if len(val) > size {
			val = val[0:size]
		}
		copy(data, val)
	}
	return size
}
func (p *PackTool) PackArrayByte(data []byte, val []byte) int {
	copy(data, val)
	return len(val)
}
func (p *PackTool) PackSliceUint16(data []byte, val []uint16, size int) int {
	offset := 0
	for i := 0; i < size; i++ {
		offset += p.PackUint16(data[offset:], val[i])
	}
	return offset
}
func (p *PackTool) PackArrayUint16(data []byte, val []uint16) int {
	offset := 0
	for _, v := range val {
		offset += p.PackUint16(data[offset:], v)
	}
	return offset
}
func (p *PackTool) PackSliceUint64(data []byte, val []uint64, size int) int {
	offset := 0
	for i := 0; i < size; i++ {
		offset += p.PackUint64(data[offset:], val[i])
	}
	return offset
}
func (p *PackTool) PackArrayUint64(data []byte, val []uint64) int {
	offset := 0
	for _, v := range val {
		offset += p.PackUint64(data[offset:], v)
	}
	return offset
}

var DefaultPack = NewPackTool(false)

//----------------------
type BYTE byte

func (val *BYTE) Pack(data []byte) int {
	return DefaultPack.PackByte(data, byte(*val))
}
func (val *BYTE) Unpack(data []byte) int {
	return DefaultPack.UnpackByte((*byte)(val), data)
}
func (val *BYTE) Val() byte {
	return byte(*val)
}

//----------------------
type WORD uint16

func (val *WORD) Pack(data []byte) int {
	return DefaultPack.PackUint16(data, uint16(*val))
}
func (val *WORD) Unpack(data []byte) int {
	return DefaultPack.UnpackUint16((*uint16)(val), data)
}
func (val *WORD) Val() uint16 {
	return uint16(*val)
}

//----------------------
type DWORD uint32

func (val *DWORD) Pack(data []byte) int {
	return DefaultPack.PackUint32(data, uint32(*val))
}
func (val *DWORD) Unpack(data []byte) int {
	return DefaultPack.UnpackUint32((*uint32)(val), data)
}
func (val *DWORD) Val() uint32 {
	return uint32(*val)
}

//----------------------
type QWORD uint64

func (val *QWORD) Pack(data []byte) int {
	return DefaultPack.PackUint64(data, uint64(*val))
}
func (val *QWORD) Unpack(data []byte) int {
	return DefaultPack.UnpackUint64((*uint64)(val), data)
}
func (val *QWORD) Val() uint64 {
	return uint64(*val)
}

//----------------------
type String string

func (val *String) Pack(data []byte, size int) int {
	return DefaultPack.PackString(data, string(*val), size)
}
func (val *String) Unpack(data []byte, size int) int {
	return DefaultPack.UnpackString((*string)(val), data, size)
}
func (val *String) Val() string {
	return string(*val)
}
func (val *String) Set(str string) {
	*val = String(str)
}
