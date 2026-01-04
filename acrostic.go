// Package acrostic generates acrostical phrases from words or random letters.
// An acrostic phrase uses the first letter of each word to spell out a word
// or message. This package provides functionality to generate such phrases
// using word lists of adjectives and nouns.
package acrostic

import (
	"bytes"
	_ "embed"
	"errors"
	"io"
	"strings"

	"github.com/awoodbeck/acrostic/internal/words"
)

var (
	// ErrBlankAcrostic is returned when a given acrostic has a length of zero characters.
	ErrBlankAcrostic = errors.New("acrostic may not be blank")

	// ErrInvalidNumber is returned when the number of requested acrostics is less than 1.
	ErrInvalidNumber = errors.New("number of returned acrostics less than 1 is not valid")

	// ErrUninitialized is returned when acrostics are requested from an uninitialized Acrostic object.
	ErrUninitialized = errors.New("acrostic object has not been initialized")

	//go:embed files/adjectives.txt
	adjectives []byte

	//go:embed files/nouns.txt
	nouns []byte
)

// NewAcrostic accepts pointers to an adjective and noun word lists, and returns a pointer to
// a populated Acrostic object.
//
// If the adjectives word list pointer is nil, the default adjectives word list will be used.
// Likewise for the nouns word list pointer.
func NewAcrostic(adjWords, nounWords *words.Words) (*Acrostic, error) {
	var err error

	if adjWords == nil {
		adjWords, err = makeWordList(bytes.NewReader(adjectives))
		if err != nil {
			return nil, err
		}
	}
	if nounWords == nil {
		nounWords, err = makeWordList(bytes.NewReader(nouns))
		if err != nil {
			return nil, err
		}
	}

	acro := &Acrostic{
		adjectives: adjWords,
		nouns:      nounWords,
	}

	return acro, nil
}

// Acrostic maintains a list of adjectives and nouns, and returns an acrostical
// phrase for a given word.
type Acrostic struct {
	adjectives *words.Words
	nouns      *words.Words
}

// GenerateAcrostic returns a single acrostical phrase for the given word.
// Options can be provided to customize the output format.
func (a *Acrostic) GenerateAcrostic(acro string, opts ...Option) (string, error) {
	results, err := a.GenerateAcrostics(acro, 1, opts...)
	if err != nil {
		return "", err
	}
	return results[0], nil
}

// GenerateAcrostics accepts an integer indicating the number of phrases
// to return, and returns a string slice with the results.
// Options can be provided to customize the output format.
func (a *Acrostic) GenerateAcrostics(acro string, num int, opts ...Option) ([]string, error) {
	acroLen := len(acro)

	switch {
	case a.adjectives == nil || a.nouns == nil:
		return nil, ErrUninitialized
	case acroLen == 0:
		return nil, ErrBlankAcrostic
	case num < 1:
		return nil, ErrInvalidNumber
	}

	cfg := applyOptions(opts)
	results := make([]string, 0, num)
	var builder strings.Builder

	for range num {
		builder.Reset()

		for j := range acroLen {
			if j > 0 {
				builder.WriteString(cfg.separator)
			}

			var word string
			var err error

			if j == acroLen-1 {
				word, err = a.nouns.RandomWord(acro[j])
			} else {
				word, err = a.adjectives.RandomWord(acro[j])
			}

			if err != nil {
				return nil, err
			}

			builder.WriteString(word)
		}

		phrase, err := formatPhrase(builder.String(), cfg)
		if err != nil {
			return nil, err
		}

		results = append(results, phrase)
	}

	return results, nil
}

// GenerateRandomAcrostic generates a single random acrostic of the specified length.
// Options can be provided to customize the output format.
func (a *Acrostic) GenerateRandomAcrostic(length int, opts ...Option) (string, error) {
	results, err := a.GenerateRandomAcrostics(length, 1, opts...)
	if err != nil {
		return "", err
	}
	return results[0], nil
}

// GenerateRandomAcrostics accepts an acrostic length and an integer indicating the number
// of acrostics to return.
// Options can be provided to customize the output format.
func (a *Acrostic) GenerateRandomAcrostics(length, num int, opts ...Option) ([]string, error) {
	switch {
	case a.adjectives == nil || a.nouns == nil:
		return nil, ErrUninitialized
	case length < 1:
		return nil, ErrInvalidNumber
	case num < 1:
		return nil, ErrInvalidNumber
	}

	acro := make([]byte, length)

	for i := range length - 1 {
		key, err := a.adjectives.RandomKey()
		if err != nil {
			return nil, err
		}
		acro[i] = key
	}

	key, err := a.nouns.RandomKey()
	if err != nil {
		return nil, err
	}
	acro[length-1] = key

	return a.GenerateAcrostics(string(acro), num, opts...)
}

// GenerateRandomPhrase generates a single random phrase of the specified length.
// Options can be provided to customize the output format.
func (a *Acrostic) GenerateRandomPhrase(phraseLen int, opts ...Option) (string, error) {
	results, err := a.GenerateRandomPhrases(phraseLen, 1, opts...)
	if err != nil {
		return "", err
	}
	return results[0], nil
}

// GenerateRandomPhrases accepts two integers: words per phrase, and number of phrases.
// It returns a string slice matching the number of phrases.
// Options can be provided to customize the output format.
func (a *Acrostic) GenerateRandomPhrases(phraseLen, phraseNum int, opts ...Option) ([]string, error) {
	switch {
	case a.adjectives == nil || a.nouns == nil:
		return nil, ErrUninitialized
	case phraseLen < 1:
		return nil, ErrInvalidNumber
	case phraseNum < 1:
		return nil, ErrInvalidNumber
	}

	results := make([]string, 0, phraseNum)
	acro := make([]byte, phraseLen)

	for range phraseNum {
		// Generate random acrostic
		for i := range phraseLen - 1 {
			key, err := a.adjectives.RandomKey()
			if err != nil {
				return nil, err
			}
			acro[i] = key
		}

		key, err := a.nouns.RandomKey()
		if err != nil {
			return nil, err
		}
		acro[phraseLen-1] = key

		// Generate single phrase
		phrases, err := a.GenerateAcrostics(string(acro), 1, opts...)
		if err != nil {
			return nil, err
		}

		results = append(results, phrases[0])
	}

	return results, nil
}

// makeWordList accepts a file name and returns a pointer to a populated Words object.
func makeWordList(r io.Reader) (*words.Words, error) {
	buf := &bytes.Buffer{}
	if _, err := buf.ReadFrom(r); err != nil {
		return nil, err
	}
	return words.NewWords(buf)
}
