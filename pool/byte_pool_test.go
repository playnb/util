package pool

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDefaultPool(t *testing.T) {
	pool := pool2.DefaultPool().Get(1024)
	t.Log(pool)

	convey.Convey("将两数相加", t, func() {
		convey.So(1, convey.ShouldEqual, 1)
	})
}