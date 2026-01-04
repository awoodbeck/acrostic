package acrostic

import (
	"strings"
	"testing"
)

func TestWithSeparator(t *testing.T) {
	t.Parallel()

	a, err := NewAcrostic(nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Test with dash separator
	phrase, err := a.GenerateAcrostic("test", WithSeparator("-"))
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(phrase, "-") {
		t.Errorf("expected phrase to contain '-', got: %s", phrase)
	}

	words := strings.Split(phrase, "-")
	if len(words) != 4 {
		t.Errorf("expected 4 words, got %d", len(words))
	}
}

func TestWithNumber(t *testing.T) {
	t.Parallel()

	a, err := NewAcrostic(nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Test with number 0-99
	phrase, err := a.GenerateAcrostic("test", WithNumber(0, 99))
	if err != nil {
		t.Fatal(err)
	}

	words := strings.Split(phrase, " ")
	if len(words) != 5 {
		t.Errorf("expected 5 parts (4 words + 1 number), got %d: %s", len(words), phrase)
	}
}

func TestWithNumberPosition(t *testing.T) {
	t.Parallel()

	a, err := NewAcrostic(nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Test number at beginning
	phrase, err := a.GenerateAcrostic("test",
		WithNumber(1000, 1000),
		WithNumberPosition(NumberPositionBeginning))
	if err != nil {
		t.Fatal(err)
	}

	if !strings.HasPrefix(phrase, "1000 ") {
		t.Errorf("expected phrase to start with '1000 ', got: %s", phrase)
	}

	// Test number at end
	phrase, err = a.GenerateAcrostic("test",
		WithNumber(2000, 2000),
		WithNumberPosition(NumberPositionEnd))
	if err != nil {
		t.Fatal(err)
	}

	if !strings.HasSuffix(phrase, " 2000") {
		t.Errorf("expected phrase to end with ' 2000', got: %s", phrase)
	}
}

func TestWithCapitalization(t *testing.T) {
	t.Parallel()

	a, err := NewAcrostic(nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Test capitalize first
	phrase, err := a.GenerateAcrostic("test", WithCapitalization(CapitalizeFirst))
	if err != nil {
		t.Fatal(err)
	}

	words := strings.Split(phrase, " ")
	for i, word := range words {
		if len(word) > 0 && word[0] < 'A' || word[0] > 'Z' {
			t.Errorf("word %d (%s) should start with uppercase letter", i, word)
		}
	}

	// Test capitalize all
	phrase, err = a.GenerateAcrostic("test", WithCapitalization(CapitalizeAll))
	if err != nil {
		t.Fatal(err)
	}

	if phrase != strings.ToUpper(phrase) {
		t.Errorf("expected all uppercase, got: %s", phrase)
	}
}

func TestMultipleOptions(t *testing.T) {
	t.Parallel()

	a, err := NewAcrostic(nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Test combining multiple options
	phrase, err := a.GenerateAcrostic("test",
		WithSeparator("-"),
		WithNumber(100, 999),
		WithCapitalization(CapitalizeFirst))
	if err != nil {
		t.Fatal(err)
	}

	// Should have dashes
	if !strings.Contains(phrase, "-") {
		t.Errorf("expected phrase to contain '-', got: %s", phrase)
	}

	// Should have 5 parts
	parts := strings.Split(phrase, "-")
	if len(parts) != 5 {
		t.Errorf("expected 5 parts, got %d: %s", len(parts), phrase)
	}

	t.Logf("Generated phrase: %s", phrase)
}

func TestRandomPhraseWithOptions(t *testing.T) {
	t.Parallel()

	a, err := NewAcrostic(nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	phrases, err := a.GenerateRandomPhrases(4, 3,
		WithSeparator("_"),
		WithNumber(0, 9999),
		WithCapitalization(CapitalizeFirst))
	if err != nil {
		t.Fatal(err)
	}

	if len(phrases) != 3 {
		t.Fatalf("expected 3 phrases, got %d", len(phrases))
	}

	for i, phrase := range phrases {
		t.Logf("Phrase %d: %s", i+1, phrase)

		if !strings.Contains(phrase, "_") {
			t.Errorf("phrase %d should contain '_', got: %s", i, phrase)
		}
	}
}
