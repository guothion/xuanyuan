package util

import "time"

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
