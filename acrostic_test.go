package acrostic

import "testing"

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
