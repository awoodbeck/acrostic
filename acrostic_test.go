package acrostic

import (
	"strings"
	"testing"
)

func TestNewAcrostic(t *testing.T) {
	t.Parallel()

	a, err := NewAcrostic(nil, nil)
	if err != nil {
		t.Fatal("expected nil error; actual error:", err)
	}
	if a.adjectives == nil {
		t.Error("expected populated adjectives; actual adjectives pointer = nil")
	}
	if a.nouns == nil {
		t.Error("expected populated nouns; actual nouns pointer = nil")
	}
}

func TestGenerateAcrostics(t *testing.T) {
	t.Parallel()

	var err error
	a := &Acrostic{}
	if _, err = a.GenerateAcrostics("test", 1); err != ErrUninitialized {
		t.Error("expected ErrUninitialized error; actual error:", err)
	}

	a, err = NewAcrostic(nil, nil)
	if err != nil {
		t.Fatal("expected nil error; actual error:", err)
	}

	if _, err = a.GenerateAcrostics("", 1); err != ErrBlankAcrostic {
		t.Error("expected ErrBlankAcrostic error; actual error:", err)
	}

	if _, err = a.GenerateAcrostics("test", 0); err != ErrInvalidNumber {
		t.Error("expected ErrInvalidNumber error; actual error:", err)
	}

	keys := make([]byte, 4)

	keys[0], err = a.adjectives.GetRandomKey()
	if err != nil {
		t.Fatal("expected nil error; actual error:", err)
	}
	keys[1], err = a.adjectives.GetRandomKey()
	if err != nil {
		t.Fatal("expected nil error; actual error:", err)
	}
	keys[2], err = a.adjectives.GetRandomKey()
	if err != nil {
		t.Fatal("expected nil error; actual error:", err)
	}
	keys[3], err = a.nouns.GetRandomKey()
	if err != nil {
		t.Fatal("expected nil error; actual error:", err)
	}

	var acros []string
	acros, err = a.GenerateAcrostics(string(keys), 10)
	if err != nil {
		t.Fatal("expected nil error; actual error:", err)
	}

	for _, acro := range acros {
		t.Log(acro)
		words := strings.Split(acro, " ")
		for i, word := range words {
			if word[0] != keys[i] {
				t.Errorf("letter %q does not match the first letter of the word %q", keys[i], word)
			}
		}
	}
}

func TestGenerateRandomAcrostics(t *testing.T) {
	t.Parallel()

	var err error
	a := &Acrostic{}
	if _, err = a.GenerateRandomAcrostics(4, 1); err != ErrUninitialized {
		t.Error("expected ErrUninitialized error; actual error:", err)
	}

	a, err = NewAcrostic(nil, nil)
	if err != nil {
		t.Fatal("expected nil error; actual error:", err)
	}

	if _, err = a.GenerateRandomAcrostics(0, 1); err != ErrInvalidNumber {
		t.Error("expected ErrInvalidNumber error; actual error:", err)
	}

	if _, err = a.GenerateRandomAcrostics(1, 0); err != ErrInvalidNumber {
		t.Error("expected ErrInvalidNumber error; actual error:", err)
	}

	var acros []string
	acros, err = a.GenerateRandomAcrostics(4, 10)
	if err != nil {
		t.Fatal(err)
	}
	if l := len(acros); l != 10 {
		t.Error("expected 10 acrostics; actual =", l)
	}
	for _, acro := range acros {
		t.Log(acro)
	}
}

func TestGenerateRandomPhrases(t *testing.T) {
	t.Parallel()

	var err error
	a := &Acrostic{}
	if _, err = a.GenerateRandomPhrases(4, 1); err != ErrUninitialized {
		t.Error("expected ErrUninitialized error; actual error:", err)
	}

	a, err = NewAcrostic(nil, nil)
	if err != nil {
		t.Fatal("expected nil error; actual error:", err)
	}

	if _, err = a.GenerateRandomPhrases(0, 1); err != ErrInvalidNumber {
		t.Error("expected ErrInvalidNumber error; actual error:", err)
	}

	if _, err = a.GenerateRandomPhrases(1, 0); err != ErrInvalidNumber {
		t.Error("expected ErrInvalidNumber error; actual error:", err)
	}

	var p []string
	p, err = a.GenerateRandomPhrases(4, 10)
	if err != nil {
		t.Fatal(err)
	}
	if l := len(p); l != 10 {
		t.Error("expected 10 phrases; actual =", l)
	}
}
