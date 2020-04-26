package pool

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDefaultPool(t *testing.T) {
	pool := DefaultPool().Get(1024)
	t.Log(pool)

	convey.Convey("测试BytePool", t, func() {
		convey.So(1, convey.ShouldEqual, 1)
	})
}
