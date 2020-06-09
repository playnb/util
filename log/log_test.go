package log

import (
	c "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDefaultLogger(t *testing.T) {
	Init(DefaultLogger(`C:\ltp\code\log`, "test"))
	InitPanic(`C:\ltp\code\panic`)
	defer Flush()
	c.Convey("测试Log", t, func() {
		defer PrintPanicStack()

		Debug("debug........")
		Info("info........")
		Trace("trace........")
		Error("error........")
		Fatal("fatal........")
		c.So(MakePanic, c.ShouldPanic)
	})
}
