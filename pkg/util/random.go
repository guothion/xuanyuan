package util

import (
	"math/rand"
	"time"
)

const (
	HexString = "0123456789abcdef"
)

func RandString(n int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, n)
	for i := range b {
		b[i] = HexString[r.Intn(len(HexString))]
	}
	return string(b)
}

func RandBool() bool {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Int()%2 == 0
}

func RandInt63n(n int64) int64 {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(n)
}
