package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// random int between min max value
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// Random string
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)] //Intn(n int) random [0,n)
		sb.WriteByte(c)
	}

	return sb.String()
}

// Account Random
func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 10000)
}

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "VND", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
