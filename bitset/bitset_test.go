package bitset

import (
	c "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBitset(t *testing.T) {
	c.Convey("TestBitset", t, func() {
		b := New(1024)

		b.Set(0)
		c.So(b.Test(0), c.ShouldBeTrue)
		b.Set(10)
		c.So(b.Test(10), c.ShouldBeTrue)
		b.Set(64)
		c.So(b.Test(64), c.ShouldBeTrue)
		b.Set(80)
		c.So(b.Test(80), c.ShouldBeTrue)
		b.Clear(80)
		c.So(b.Test(80), c.ShouldBeFalse)
		b.Set(1023)
		c.So(b.Test(1023), c.ShouldBeTrue)

		c.So(b.Test(11), c.ShouldBeFalse)
		c.So(b.Test(65), c.ShouldBeFalse)
		c.So(b.Test(70), c.ShouldBeFalse)

		c.So(func() { b.Set(1024) }, c.ShouldPanic)
		c.So(func() { b.Clear(1024) }, c.ShouldPanic)
	})
}
