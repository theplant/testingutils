package random

import (
	"math/rand"
	"time"
)

const (
	LettersAndNumbersBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	LettersBytes           = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	NumbersBytes           = "0123456789"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func StringWithBaseBytes(length int, baseBytes string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = baseBytes[rand.Intn(len(baseBytes))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithBaseBytes(length, LettersAndNumbersBytes)
}

func Letters(length int) string {
	return StringWithBaseBytes(length, LettersBytes)
}

func Numbers(length int) string {
	return StringWithBaseBytes(length, NumbersBytes)
}

func EmailWithLengthAndDomainName(length int, domainName string) string {
	return String(length) + "@" + domainName
}

func Email() string {
	return EmailWithLengthAndDomainName(20, "example.com")
}

func Bool() bool {
	return rand.Int63n(2) == 0
}
