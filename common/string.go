package common

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func StringRand(length int) string {
	return StringWithCharset(length, charset)
}

func StringReverse(s string) string {
    rs := []rune(s)
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        rs[i], rs[j] = rs[j], rs[i]
    }
    return string(rs)
}