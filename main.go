package main

import (
	"crypto/rand"
	"fmt"
	"log/slog"
	"math"
	"math/big"
	mranf "math/rand"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var digits = []rune("0123456789")
var symbols = []rune("-._~@!$%&*")

func main() {
	pLen := 20
	pBuf := []rune{}
	it := math.Floor(float64(pLen) / 3)

	for range int(it) {
		pBuf = append(pBuf, randomLetter())
	}

	for range int(it) {
		pBuf = append(pBuf, randomDigit())
	}

	for range int(it) {
		pBuf = append(pBuf, randomSymbol())
	}

	for range pLen - len(pBuf) {
		pBuf = append(pBuf, randomLetter())
	}

	suffle(pBuf)

	pstr := string(pBuf)
	fmt.Println(pstr)
}

func suffle(r []rune) {
	mranf.Shuffle(len(r), func(i, j int) {
		r[i], r[j] = r[j], r[i]
	})
}

func randomSymbol() rune {
	ri, err := rand.Int(rand.Reader, big.NewInt(int64(len(symbols))))
	if err != nil {
		slog.Error(err.Error())
	}
	return symbols[ri.Int64()]
}

func randomDigit() rune {
	ri, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
	if err != nil {
		slog.Error(err.Error())
	}
	return digits[ri.Int64()]
}

func randomLetter() rune {
	ri, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
	if err != nil {
		slog.Error(err.Error())
	}
	return letters[ri.Int64()]
}
