// Package words provides functionality for managing word lists indexed by
// their first letter, with support for random word selection and entropy
// calculation.
package words

import (
	"bytes"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"math"
	"math/big"
)

var (
	// ErrEmptyWordList is returned when an operation is attempted on an empty word list.
	ErrEmptyWordList = errors.New("word list is empty")

	// ErrNilBuffer is returned when a nil buffer is provided to NewWords.
	ErrNilBuffer = errors.New("nil buffer")

	// ErrEmptyBuffer is returned when an empty buffer is provided to NewWords.
	ErrEmptyBuffer = errors.New("empty buffer")

	// ErrInvalidMax is returned when max value is less than 1.
	ErrInvalidMax = errors.New("max cannot be less than 1")

	// ErrIntegerOverflow is returned when word count exceeds int64 capacity.
	ErrIntegerOverflow = errors.New("integer overflow")
)

// NewWords accepts a pointer to a bytes buffer and uses its contents
// to populate a new Words object before returning its pointer.
func NewWords(buf *bytes.Buffer) (*Words, error) {
	w := &Words{words: make(map[byte][]string)}

	err := w.compileWords(buf)

	return w, err
}

// Words maintains a word list, facilitates calculation of
// entropy, and can return a random word upon request.
type Words struct {
	words map[byte][]string
	keys  []byte
}

// Entropy returns the bits of entropy for "n" number of words.
//
// For example, if the word list was 2048 words long and 4 words were
// chosen at random for a passphrase, the entropy would be 44 bits.
//
// This is only taking into consideration the number of words. Capitalizing
// words, or choosing different separators will increase the entropy on paper.
func (w *Words) Entropy(n int) (float64, error) {
	i, err := w.wordCount()
	if err != nil {
		return 0.0, err
	}

	return math.Log2(math.Pow(float64(i), float64(n))), nil
}

// RandomKey returns a random key from the map.
func (w *Words) RandomKey() (byte, error) {
	if len(w.keys) == 0 {
		return 0, ErrEmptyWordList
	}

	n, err := randomInt(len(w.keys))
	if err != nil {
		return 0, err
	}

	return w.keys[n], nil
}

// RandomWord returns a random word from the map for the given key.
func (w *Words) RandomWord(key byte) (string, error) {
	words, ok := w.words[key]
	if !ok || len(words) == 0 {
		return "", fmt.Errorf("no words found starting with %q", key)
	}

	n, err := randomInt(len(words))
	if err != nil {
		return "", err
	}

	return words[n], err
}

// compileWords parses the given bytes buffer and populates
// the Words object.
//
// All words are lower cased and any extraneous spaces are trimmed.
func (w *Words) compileWords(buf *bytes.Buffer) error {
	var k byte

	switch {
	case buf == nil:
		return ErrNilBuffer
	case buf.Len() == 0:
		return ErrEmptyBuffer
	}

Outer:
	for {
		b, err := buf.ReadBytes('\n')
		switch {
		case err != io.EOF && err != nil:
			return err
		case err == io.EOF && len(b) == 0:
			break Outer
		}

		b = bytes.ToLower(bytes.TrimSpace(b))
		if len(b) == 0 {
			continue
		}

		k = b[0]

		if _, ok := w.words[k]; !ok {
			w.words[k] = []string{string(b)}
			w.keys = append(w.keys, k)
			continue
		}

		w.words[k] = append(w.words[k], string(b))
	}

	return nil
}

// wordCount returns the total number of words in the current object.
func (w *Words) wordCount() (int64, error) {
	var c int64

	for _, v := range w.words {
		c += int64(len(v))
	}

	if c < 0 {
		return 0, ErrIntegerOverflow
	}

	return c, nil
}

// randomInt generates a random integer from 0 to max.
func randomInt(maxInt int) (int, error) {
	if maxInt < 1 {
		return 0, ErrInvalidMax
	}

	i, err := rand.Int(rand.Reader, big.NewInt(int64(maxInt)))
	if err != nil {
		return 0, err
	}

	return int(i.Int64()), nil
}
