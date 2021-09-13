package logic

import (
	"math/rand"
)

var HexCharacters = []rune("013456789ABCDEF")

func randomString(length uint16, characters []rune) (retval string) {
	creating := make([]rune, length)
	for i := range creating {
		creating[i] = characters[rand.Intn(len(characters))]
	}

	return string(creating)
}

func GenerateId(length uint16) (result string) {
	return randomString(length, HexCharacters)
}
