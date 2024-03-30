package pkg

import (
	"math/rand"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var digits = []rune("0123456789")
var symbols = []rune("-._~@!$%&*")

func Generate(n int, l, s, d bool) string {
	var charsets []rune

	if l {
		charsets = append(charsets, letters...)
	}
	if s {
		charsets = append(charsets, symbols...)
	}
	if d {
		charsets = append(charsets, digits...)
	}

	buf := make([]rune, n)

	for i := 0; i < n; i++ {
		buf[i] = randomChar(charsets)
	}

	shuffle(buf)
	return string(buf)
}

func randomChar(chars []rune) rune {
	return chars[rand.Intn(len(chars))]
}

func shuffle(r []rune) {
	rand.Shuffle(len(r), func(i, j int) {
		r[i], r[j] = r[j], r[i]
	})
}
