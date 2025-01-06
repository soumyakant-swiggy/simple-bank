package util

import (
	"math/rand"
	"strings"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyz"

var currency = [2]string{"USD", "INR"}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random number between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1) // min->max
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(letters)

	for i := 0; i < n; i++ {
		c := letters[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomOwnerName returns a random owner name
func RandomOwnerName() string {
	return RandomString(6)
}

// RandomAmount returns a random amount
func RandomAmount() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency returns a random currency
func RandomCurrency() string {
	return currency[rand.Intn(len(currency))]
}
