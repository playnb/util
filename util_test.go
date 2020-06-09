package util

import "fmt"

func ShouldSameByteSlice(actual interface{}, expected ...interface{}) string {
	d1 := (actual).([]byte)
	d2 := (expected[0]).([]byte)
	if len(d1) != len(d2) {
		return fmt.Sprintf("需要:%v 却得到:%v", actual, expected[0])
	}
	for i := 0; i < len(d1) && i < len(d2); i++ {
		if d1[i] != d2[i] {
			return fmt.Sprintf("需要:%v 却得到:%v", actual, expected[0])
		}
	}
	return ""
}

func ShouldSameIntSlice(actual interface{}, expected ...interface{}) string {
	d1 := (actual).([]int)
	d2 := (expected[0]).([]int)
	if len(d1) != len(d2) {
		return fmt.Sprintf("需要:%v 却得到:%v", actual, expected[0])
	}
	for i := 0; i < len(d1) && i < len(d2); i++ {
		if d1[i] != d2[i] {
			return fmt.Sprintf("需要:%v 却得到:%v", actual, expected[0])
		}
	}
	return ""
}
