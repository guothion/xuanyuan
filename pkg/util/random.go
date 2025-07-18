package util

import (
	"math/rand"
	"time"
)

func RandBool() bool {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Int()%2 == 0
}

func RandInt63n(n int64) int64 {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(n)
}
