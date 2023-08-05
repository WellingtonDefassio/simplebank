package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghtjklmnopqrstuvxzyw"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString returns a random string of length stringSize
func RandomString(stringSize int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < stringSize; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomOwner generate a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generate a random currency
func RandomCurrency() string {
	currencies := []string{"EUR", "BRL", "USD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
