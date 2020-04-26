package util

import "time"

var DelayTime = int64(0)

func NowTimestamp() int64 {
	return time.Now().Unix() + DelayTime
}

func NowTimestampMillsecond() int64 {
	return time.Now().UnixNano()/int64(time.Millisecond) + DelayTime*1000
}
