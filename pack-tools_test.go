package util

import (
	c "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestPackTool(t *testing.T) {
	u8 := byte(0)
	_ = u8
	u16 := uint16(0)
	_ = u16
	u32 := uint32(0)
	_ = u32
	u64 := uint64(0)
	i32 := int32(-32)
	str := ""
	_ = str
	boo := false
	_ = boo
	sampleBuf := []byte{1, 23, 4, 5, 6, 7, 98, 0, 0, 77}
	buf := make([]byte, len(sampleBuf))
	sampleArray := [5]byte{10, 2, 3, 40, 5}
	var array [5]byte

	c.Convey("测试PackTool", t, func() {

		p := NewPackTool(false)
		data := make([]byte, 1000)
		{
			//Pack
			offset := 0
			offset += p.PackByte(data[offset:], 65)
			offset += p.PackUint32(data[offset:], 1101)
			offset += p.PackUint64(data[offset:], 99998)
			offset += p.PackUint16(data[offset:], 98)
			offset += p.PackString(data[offset:], "hello", 10)
			offset += p.PackSlice(data[offset:], sampleBuf, len(sampleBuf))
			offset += p.PackUint32(data[offset:], 999)
			offset += p.PackArrayByte(data[offset:], sampleArray[0:])
			offset += p.PackBool(data[offset:], true)
			offset += p.PackBool(data[offset:], false)
			offset += p.PackByte(data[offset:], 66)
			offset += p.PackInt32(data[offset:], -66)
		}
		{
			//Unpack
			offset := 0
			offset += p.UnpackByte(&u8, data[offset:])
			c.So(u8, c.ShouldEqual, 65)

			offset += p.UnpackUint32(&u32, data[offset:])
			c.So(u32, c.ShouldEqual, 1101)

			offset += p.UnpackUint64(&u64, data[offset:])
			c.So(u64, c.ShouldEqual, 99998)

			offset += p.UnpackUint16(&u16, data[offset:])
			c.So(u16, c.ShouldEqual, 98)

			offset += p.UnpackString(&str, data[offset:], 10)
			c.So(str, c.ShouldEqual, "hello")

			offset += p.UnpackSlice(&buf, data[offset:], len(sampleBuf))
			c.So(buf, ShouldSameByteSlice, sampleBuf)

			offset += p.UnpackUint32(&u32, data[offset:])
			c.So(u32, c.ShouldEqual, 999)

			offset += p.UnpackArrayByte(array[0:], data[offset:])
			c.So(array, c.ShouldEqual, sampleArray)

			offset += p.UnpackBool(&boo, data[offset:])
			c.So(boo, c.ShouldEqual, true)

			offset += p.UnpackBool(&boo, data[offset:])
			c.So(boo, c.ShouldEqual, false)

			offset += p.UnpackByte(&u8, data[offset:])
			c.So(u8, c.ShouldEqual, 66)

			offset += p.UnpackInt32(&i32, data[offset:])
			c.So(i32, c.ShouldEqual, -66)
		}
	})
}

func TestPackTool2(t *testing.T) {

}
