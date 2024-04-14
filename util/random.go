package util

import (
	"math/rand"
	"strings"
	"time"
)

const (
	alphabets = "abcdefghijklmnopqrstuvwxyz"
)

var rnd *rand.Rand

func init() {
	source := rand.NewSource(time.Now().UnixNano())
	rnd = rand.New(source)
}

func RandomOwner() string {
	return randomString(6)
}

func RandomBalance() int64 {
	return randomInt(1000, 10000000)
}

func RandomAmount() int64 {
	return randomInt(-100000, 100000)
}

func RandomPositiveAmount() int64 {
	return randomInt(1, 100000)
}

func RandomCurrency() string {
	currencies := []string{"INR", "USD", "EUR", "AED"}
	k := len(currencies)
	return currencies[rnd.Intn(k)]
}

func randomString(n int) string {
	var sb strings.Builder

	k := len(alphabets)

	for i := 0; i < n; i++ {
		c := alphabets[rnd.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func randomInt(min int64, max int64) int64 {
	return min + rnd.Int63n(max-(min+1))
}
