//go:generate go run -tags=dev assets_generate.go

package acrostic

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
)

// Acrostic maintains a list of adjectives and nouns, and returns an acrostical
// phrase for a given word.
type Acrostic struct {
}

// calcEntropy returns the calculated entropy where "t" is the total number of
// items available for random selection, and "n" is the number of random
// items selected.
func calcEntropy(t, n float64) float64 {
	return math.Log2(math.Pow(t, n))
}

// randomInt64 generates a random 64-bit integer from 0 to "max."
func randomInt64(max int64) (n int64, err error) {
	if max < 1 {
		err = fmt.Errorf("The given max (%d) cannot be less than 1.", max)
		return
	}

	var i *big.Int
	i, err = rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		return
	}

	n = i.Int64()

	return
}
