package util

import "time"

const (
	SimpleTimeFormat = "2006-01-02 15:04:05"
)

func JitterDuration(baseDuration, maxJitter time.Duration) time.Duration {
	var sign int64 = 1
	if RandBool() {
		sign = -1
	}
	return baseDuration + time.Duration(sign*RandInt63n(int64(maxJitter)))
}

func RandomSleep(baseDuration, maxJitter time.Duration) {
	duration := JitterDuration(baseDuration, maxJitter)
	time.Sleep(duration)
}

func Timestamp() int64 {
	return time.Now().UnixNano() / 1e6
}
