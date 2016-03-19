//go:generate go run -tags=dev assets_generate.go

package acrostic

import (
	"bytes"
	"errors"
	"net/http"

	"github.com/awoodbeck/acrostic/assets"
	"github.com/awoodbeck/acrostic/words"
)

// ErrBlankAcrostic is returned when a given acrostic has a length of zero characters.
var ErrBlankAcrostic = errors.New("acrostic may not be blank")

// ErrInvalidNumber is returned when the number of requested acrostics is less than 1.
var ErrInvalidNumber = errors.New("number of returned acrostics less than 1 is not valid")

// ErrUninitialized is returned when acrostics are requested from an uninitalized Acrostic object.
var ErrUninitialized = errors.New("acrostic object has not been initialized")

// NewAcrostic accepts pointers to an adjective and noun word lists, and returns a pointer to
// a populated Acrostic object.
//
// If the adjectives word list pointer is nil, the default adjectives word list will be used.
// Likewise for the nouns word list pointer.
func NewAcrostic(adjs, nouns *words.Words) (acro *Acrostic, err error) {
	if adjs == nil {
		var adj http.File
		adj, err = assets.Assets.Open("adjectives.txt")
		if err != nil {
			return
		}
		buf := &bytes.Buffer{}
		if _, err = buf.ReadFrom(adj); err != nil {
			return
		}
		adjs, err = words.NewWords(buf)
		if err != nil {
			return
		}
	}
	if nouns == nil {
		var adj http.File
		adj, err = assets.Assets.Open("nouns.txt")
		if err != nil {
			return
		}
		buf := &bytes.Buffer{}
		if _, err = buf.ReadFrom(adj); err != nil {
			return
		}
		nouns, err = words.NewWords(buf)
		if err != nil {
			return
		}
	}
	acro = &Acrostic{
		adjectives: adjs,
		nouns:      nouns,
	}

	return
}

// Acrostic maintains a list of adjectives and nouns, and returns an acrostical
// phrase for a given word.
type Acrostic struct {
	adjectives *words.Words
	nouns      *words.Words
}

// GenerateAcrostic accepts an integer indicating the number of phrases
// to return, and returns a string slice with the results.
func (a *Acrostic) GenerateAcrostic(acro string, n int) (o []string, err error) {
	var w string
	acroLen := len(acro)

	switch {
	case a.adjectives == nil || a.nouns == nil:
		err = ErrUninitialized
	case acroLen == 0:
		err = ErrBlankAcrostic
	case n < 1:
		err = ErrInvalidNumber
	}

	if err != nil {
		return
	}

	for i := 0; i < n; i++ {
		var p []byte
		for j := 0; j < acroLen; j++ {
			switch {
			case j == acroLen-1:
				w, err = a.nouns.GetRandomWord(acro[j])
				if err != nil {
					return
				}
				p = append(p, []byte(" "+w)...)
			case j > 0:
				p = append(p, ' ')
				fallthrough
			default:
				w, err = a.nouns.GetRandomWord(acro[j])
				if err != nil {
					return
				}
				p = append(p, []byte(" "+w)...)
			}
		}
		o = append(o, string(p))
	}

	return
}
