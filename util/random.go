package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmniopqrstuvwxyz"

// it generates a new source for random numbers (rand.NewSource()), and then creates a local random rand.New().
func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of lenght n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random money number
func RandomMoney() int64 {
	return RandomInt(0, 1000)

}

//RandomCurrency generates a random type of currency

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "ARS"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
