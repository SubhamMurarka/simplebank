package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

var randomGenerator *rand.Rand // Declare the randomGenerator as a package-level variable

func init() {
	newSource := rand.NewSource(time.Now().UnixNano())
	randomGenerator = rand.New(newSource)
}

func RandomInt(min, max int64) int64 {
	return randomGenerator.Int63n(max-min+1) + min
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[randomGenerator.Intn(k)] // Use randomGenerator
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"RS", "USD", "EUR"}
	n := len(currencies)
	return currencies[randomGenerator.Intn(n)] // Use randomGenerator
}
