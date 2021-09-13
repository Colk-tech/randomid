package util

import (
	"math/rand"
)

func SetRandSeed(seed int64) {
	rand.Seed(seed)
}
