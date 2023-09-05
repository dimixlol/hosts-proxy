package utils

import (
	"math/rand"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyz")

func RandomStringWithLength(n uint) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
