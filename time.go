package util

import "time"

var DelaySeconds = int64(0)

func NowTimestamp() int64 {
	return time.Now().Unix() + DelaySeconds
}

func NowTimestampMillisecond() int64 {
	return time.Now().UnixNano()/int64(time.Millisecond) + DelaySeconds*1000
}
