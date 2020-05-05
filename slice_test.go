package util

import (
	"fmt"
	c "github.com/smartystreets/goconvey/convey"
	"testing"
)

type _mockData struct {
	id   int
	name string
}

func buildSlice() []*_mockData {
	var v []*_mockData
	for i := 0; i < 10; i++ {
		v = append(v, &_mockData{
			id:   i,
			name: fmt.Sprintf("Name_%d", i),
		})
	}
	return v
}
func TestDelSlice(t *testing.T) {
	c.Convey("测试DelSlice", t, func() {
		s := buildSlice()
		s[0].id = 3
		s[0].id = 3

		DelSlice(&s, func(d interface{}) bool {
			return d.(*_mockData).id == 3
		})
		for _, v := range s {
			c.So(v.id, c.ShouldNotEqual, 3)
		}
		c.So(len(s), c.ShouldNotEqual, 7)
	})
}
func TestDelSliceIndex(t *testing.T) {
	c.Convey("测试DelSliceIndex", t, func() {
		s := buildSlice()
		s[0].id = 3
		s[0].id = 3

		for k, v := range s {
			if v.id == 3 {
				DelSliceIndex(&s, k)
			}
		}

		for _, v := range s {
			c.So(v.id, c.ShouldNotEqual, 3)
		}
		c.So(len(s), c.ShouldNotEqual, 7)
	})
}

func TestSubSlice(t *testing.T) {
	s1 := []int{0, 1, 2, 3, 4, 5}
	c.Convey("测试SubSlice", t, func() {
		c.So(SubSlice(s1, 0, 2).([]int), ShouldSameIntSlice, []int{0, 1})
		c.So(SubSlice(s1, 3, 6).([]int), ShouldSameIntSlice, []int{3, 4, 5})
		c.So(SubSlice(s1, 3, 8).([]int), ShouldSameIntSlice, []int{3, 4, 5})
		c.So(SubSlice(s1, 3, 1).([]int), ShouldSameIntSlice, SubSlice(s1, 1, 3).([]int))
		c.So(SubSlice(s1, 5, 5).([]int), ShouldSameIntSlice, []int{})
	})
}
