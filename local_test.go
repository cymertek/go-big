package big

import (
	"math/rand"
)

var isRaceBuilder = false
var support_adx = false

// Always the same seed for reproducible results.
var rnd = rand.New(rand.NewSource(0))

func rndW() Word {
	return Word(rnd.Int63()<<1 | rnd.Int63n(2))
}

func rndV(n int) []Word {
	v := make([]Word, n)
	for i := range v {
		v[i] = rndW()
	}
	return v
}
