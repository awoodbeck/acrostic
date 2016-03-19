package acrostic

import (
	"strings"
	"testing"
)

func TestNewAcrostic(t *testing.T) {
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

func TestGenerateAcrostic(t *testing.T) {
	x := Acrostic{}
	if _, err := x.GenerateAcrostic("test", 1); err != ErrUninitialized {
		t.Error("expected ErrUninitialized error; actual error:", err)
	}

	a, err := NewAcrostic(nil, nil)
	if err != nil {
		t.Fatal("expected nil error; actual error:", err)
	}

	if _, err := a.GenerateAcrostic("", 1); err != ErrBlankAcrostic {
		t.Error("expected ErrBlankAcrostic error; actual error:", err)
	}

	if _, err := a.GenerateAcrostic("test", 0); err != ErrInvalidNumber {
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

	acros, err := a.GenerateAcrostic(string(keys), 10)

	for _, acro := range acros {
		words := strings.Split(acro, " ")
		for i, word := range words {
			if word[0] != keys[i] {
				t.Errorf("letter %q does not match the first letter of the word %q", keys[i], word)
			}
		}
	}
}
