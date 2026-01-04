package acrostic

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
	"unicode"
)

// Option is a functional option for configuring phrase generation.
type Option interface {
	apply(*phraseConfig)
}

// phraseConfig holds the configuration for phrase generation.
type phraseConfig struct {
	separator      string
	numberMin      int
	numberMax      int
	includeNumber  bool
	capitalize     CapitalizeMode
	numberPosition NumberPosition
}

// CapitalizeMode defines how words should be capitalized.
type CapitalizeMode int

const (
	// CapitalizeNone performs no capitalization (default).
	CapitalizeNone CapitalizeMode = iota
	// CapitalizeFirst capitalizes the first letter of each word.
	CapitalizeFirst
	// CapitalizeAll capitalizes all letters.
	CapitalizeAll
	// CapitalizeRandom randomly capitalizes letters.
	CapitalizeRandom
)

// NumberPosition defines where the number should be placed.
type NumberPosition int

const (
	// NumberPositionEnd appends the number at the end (default).
	NumberPositionEnd NumberPosition = iota
	// NumberPositionBeginning prepends the number at the beginning.
	NumberPositionBeginning
	// NumberPositionRandom places the number at a random position.
	NumberPositionRandom
)

// defaultConfig returns the default phrase configuration.
func defaultConfig() *phraseConfig {
	return &phraseConfig{
		separator:      " ",
		numberMin:      0,
		numberMax:      9999,
		includeNumber:  false,
		capitalize:     CapitalizeNone,
		numberPosition: NumberPositionEnd,
	}
}

// optionFunc wraps a function to implement the Option interface.
type optionFunc func(*phraseConfig)

func (f optionFunc) apply(c *phraseConfig) {
	f(c)
}

// WithSeparator sets a custom separator between words.
// Default is a single space " ".
func WithSeparator(sep string) Option {
	return optionFunc(func(c *phraseConfig) {
		c.separator = sep
	})
}

// WithNumber adds a random number to the phrase within the specified range.
// The number will be appended at the end by default.
// Use WithNumberPosition to change the placement.
func WithNumber(min, max int) Option {
	return optionFunc(func(c *phraseConfig) {
		c.includeNumber = true
		c.numberMin = min
		c.numberMax = max
	})
}

// WithNumberPosition sets where the number should be placed in the phrase.
// Only takes effect if WithNumber is also used.
func WithNumberPosition(pos NumberPosition) Option {
	return optionFunc(func(c *phraseConfig) {
		c.numberPosition = pos
	})
}

// WithCapitalization sets the capitalization mode for the phrase.
func WithCapitalization(mode CapitalizeMode) Option {
	return optionFunc(func(c *phraseConfig) {
		c.capitalize = mode
	})
}

// applyOptions applies all options to the configuration.
func applyOptions(opts []Option) *phraseConfig {
	cfg := defaultConfig()
	for _, opt := range opts {
		opt.apply(cfg)
	}
	return cfg
}

// formatPhrase applies formatting options to a phrase.
func formatPhrase(phrase string, cfg *phraseConfig) (string, error) {
	// Apply capitalization
	phrase = capitalizePhrase(phrase, cfg.capitalize)

	// Add number if requested
	if cfg.includeNumber {
		num, err := randomNumber(cfg.numberMin, cfg.numberMax)
		if err != nil {
			return "", err
		}

		numStr := fmt.Sprintf("%d", num)

		switch cfg.numberPosition {
		case NumberPositionBeginning:
			phrase = numStr + cfg.separator + phrase
		case NumberPositionEnd:
			phrase = phrase + cfg.separator + numStr
		case NumberPositionRandom:
			// Split by separator and insert at random position
			words := strings.Split(phrase, cfg.separator)
			pos, err := randomInt(len(words) + 1)
			if err != nil {
				return "", err
			}
			// Insert number at random position
			words = append(words[:pos], append([]string{numStr}, words[pos:]...)...)
			phrase = strings.Join(words, cfg.separator)
		}
	}

	return phrase, nil
}

// capitalizePhrase applies capitalization according to the mode.
func capitalizePhrase(phrase string, mode CapitalizeMode) string {
	switch mode {
	case CapitalizeNone:
		return phrase
	case CapitalizeFirst:
		return capitalizeFirstLetters(phrase)
	case CapitalizeAll:
		return strings.ToUpper(phrase)
	case CapitalizeRandom:
		return randomCapitalize(phrase)
	default:
		return phrase
	}
}

// capitalizeFirstLetters capitalizes the first letter of each word.
func capitalizeFirstLetters(s string) string {
	var result strings.Builder
	result.Grow(len(s))

	capitalizeNext := true
	for _, r := range s {
		if unicode.IsSpace(r) || r == '-' || r == '_' {
			capitalizeNext = true
			result.WriteRune(r)
		} else if capitalizeNext && unicode.IsLetter(r) {
			result.WriteRune(unicode.ToUpper(r))
			capitalizeNext = false
		} else {
			result.WriteRune(r)
		}
	}

	return result.String()
}

// randomCapitalize randomly capitalizes letters.
func randomCapitalize(s string) string {
	var result strings.Builder
	result.Grow(len(s))

	for _, r := range s {
		if unicode.IsLetter(r) {
			// 50% chance to capitalize
			if shouldCapitalize, _ := randomInt(2); shouldCapitalize == 1 {
				result.WriteRune(unicode.ToUpper(r))
			} else {
				result.WriteRune(r)
			}
		} else {
			result.WriteRune(r)
		}
	}

	return result.String()
}

// randomNumber generates a random number in the range [min, max].
func randomNumber(min, max int) (int, error) {
	if min > max {
		return 0, fmt.Errorf("min (%d) cannot be greater than max (%d)", min, max)
	}

	if min == max {
		return min, nil
	}

	rangeSize := max - min + 1
	n, err := rand.Int(rand.Reader, big.NewInt(int64(rangeSize)))
	if err != nil {
		return 0, err
	}

	return int(n.Int64()) + min, nil
}

// randomInt generates a random integer from 0 to max-1.
func randomInt(max int) (int, error) {
	if max < 1 {
		return 0, fmt.Errorf("max cannot be less than 1")
	}

	n, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, err
	}

	return int(n.Int64()), nil
}
