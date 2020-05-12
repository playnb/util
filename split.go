package util

import (
	"strings"
)

func StringToIntVector(str string, sep string) []int {
	ss := strings.Split(str, sep)
	var ret []int
	for _, s := range ss {
		ret = append(ret, StringToInt(s))
	}
	return ret
}

func StringToUint16Vector(str string, sep string) []uint16 {
	ss := strings.Split(str, sep)
	var ret []uint16
	for _, s := range ss {
		ret = append(ret, StringToUint16(s))
	}
	return ret
}

func StringToUint32Vector(str string, sep string) []uint32 {
	ss := strings.Split(str, sep)
	var ret []uint32
	for _, s := range ss {
		ret = append(ret, StringToUint32(s))
	}
	return ret
}

func StringToUint64Vector(str string, sep string) []uint64 {
	ss := strings.Split(str, sep)
	var ret []uint64
	for _, s := range ss {
		ret = append(ret, StringToUint64(s))
	}
	return ret
}
