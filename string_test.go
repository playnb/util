package util

import (
	c "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCapitalize(t *testing.T) {
	c.Convey("测试转化首字母大写", t, func() {
		c.So(Capitalize("a123"), c.ShouldEqual, "A123")
		c.So(Capitalize("1123"), c.ShouldEqual, "1123")
		c.So(Capitalize("Avv"), c.ShouldEqual, "Avv")
		c.So(Capitalize("af_ff"), c.ShouldEqual, "AfFf")
		c.So(Capitalize("acf__fcf_"), c.ShouldEqual, "AcfFcf_")
		c.So(Capitalize(""), c.ShouldEqual, "")
		c.So(Capitalize("___"), c.ShouldEqual, "_")
		c.So(Capitalize("___a"), c.ShouldEqual, "A")
	})
}
