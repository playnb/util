package util

import (
	c "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRandBetween(t *testing.T) {
	c.Convey("随机数", t, func() {
		bMin := false
		bMax := false
		min := 10
		max := 20
		for i := 0; i < 10000; i++ {
			n := RandBetweenInt(min, max)
			c.So(n, c.ShouldBeGreaterThanOrEqualTo, min)
			c.So(n, c.ShouldBeLessThanOrEqualTo, max)
			bMax = bMax || (n == max)
			bMin = bMin || (n == min)
		}
		c.So(bMax, c.ShouldBeTrue)
		c.So(bMin, c.ShouldBeTrue)
	})
}
