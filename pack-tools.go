package util

import "encoding/binary"

func NewPackTool(bigEndian bool) *PackTool {
	p := &PackTool{
		IsBigEndian: bigEndian,
	}
	return p
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
	copy(*val, data[:size])
	return size
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
func (p *PackTool) PackString(data []byte, val string, size int) int {
	copy(data[0:size], []byte(val))
	return size
}
func (p *PackTool) PackSlice(data []byte, val []byte, size int) int {
	if len(val) > size {
		val = val[0:size]
	}
	copy(data, val)
	return size
}
