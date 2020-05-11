package util

import "math/rand"

//返回随机数 [min, max]
func RandBetweenInt(min, max int) int {
	if min > max {
		t := max
		max = min
		min = t
	}
	if min == max {
		return max
	}
	return min + rand.Intn(max-min+1)
}
