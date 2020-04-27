package util

import "strings"

func Capitalize(str string) string {
	if len(str) == 0 {
		return ""
	}
	var upperStr string
	vv := []rune(str) // 后文有介绍
	need := true
	for i := 0; i < len(vv); i++ {
		if vv[i] == 95 {
			need = true
		} else if need {
			need = false
			if vv[i] >= 97 && vv[i] <= 122 { // 后文有介绍
				vv[i] -= 32 // string的码表相差32位
				upperStr += string(vv[i])
			} else {
				upperStr += string(vv[i])
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	if need {
		upperStr += "_"
	}
	return upperStr
}

func StringReplaceAll(src string, idet ...string) string {
	if len(idet)%2 != 0 {
		return src
	}
	for i := 0; i < len(idet); i += 2 {
		src = strings.ReplaceAll(src, idet[i], idet[i+1])
	}
	return src
}
