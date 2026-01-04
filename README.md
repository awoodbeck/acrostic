# Acrostic

A Go library that generates acrostical phrases and secure passphrases from words or random letter combinations.

An acrostic phrase uses the first letter of each word to spell out a word or message. This library generates memorable phrases using curated word lists of over **4,600 adjectives** and **9,300 nouns**, providing strong entropy for passphrase generation.

## Features

- **Large word lists** with excellent letter distribution
- **Flexible options pattern** for customization
- **Secure passphrase generation** with configurable entropy
- **Multiple output formats**: custom separators, capitalization, number suffixes
- **Cryptographically secure** random number generation

## Entropy

The library provides strong entropy for secure passphrases:

| Phrase Length | Base Entropy | With Number (0-9999) |
|--------------|--------------|----------------------|
| 3 words      | 36.0 bits    | 49.3 bits           |
| 4 words      | 48.5 bits    | 61.8 bits (strong)  |
| 5 words      | 61.0 bits    | 74.3 bits (very strong) |
| 6 words      | 73.5 bits    | 86.8 bits (cryptographic) |

## Installation

```bash
go get github.com/awoodbeck/acrostic
```

## Usage

### Basic Example

```go
package main

import (
  "fmt"
  "log"

  "github.com/awoodbeck/acrostic"
)

func main() {
  // Create a new acrostic generator with default word lists
  a, err := acrostic.NewAcrostic(nil, nil)
  if err != nil {
    log.Fatal(err)
  }

  // Generate a single acrostic for a word
  phrase, err := a.GenerateAcrostic("golang")
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(phrase)
  // Example output: "graceful optimistic lively adorable noble gorilla"
}
```

### Generate Multiple Acrostics

```go
// Generate 5 different acrostics for the same word
phrases, err := a.GenerateAcrostics("test", 5)
if err != nil {
  log.Fatal(err)
}

for i, phrase := range phrases {
  fmt.Printf("%d: %s\n", i+1, phrase)
}
```

### Generate Random Acrostics

```go
// Generate a random 4-word acrostic
phrase, err := a.GenerateRandomAcrostic(4)
if err != nil {
  log.Fatal(err)
}

fmt.Println(phrase)
// Example output: "brave elegant creative dragon"

// Generate multiple random phrases
phrases, err := a.GenerateRandomPhrases(5, 3)
if err != nil {
  log.Fatal(err)
}
```

### Secure Passphrases with Options

```go
// Generate a secure passphrase with custom formatting
phrase, err := a.GenerateAcrostic("secure",
  acrostic.WithSeparator("-"),
  acrostic.WithNumber(0, 9999),
  acrostic.WithCapitalization(acrostic.CapitalizeFirst),
)
// Example output: "Silly-Elegant-Careful-Unusual-Rapid-Elephant-7342"

// Random passphrase with all options
phrase, err = a.GenerateRandomPhrase(5,
  acrostic.WithSeparator("_"),
  acrostic.WithNumber(100, 999),
  acrostic.WithCapitalization(acrostic.CapitalizeFirst),
)
// Example output: "Brave_Quick_Gentle_Modern_Tiger_456"
```

## Options

All generation methods accept optional configuration:

### `WithSeparator(sep string)`
Set a custom separator between words (default: space).

```go
phrase, _ := a.GenerateAcrostic("test", acrostic.WithSeparator("-"))
// Output: "tiny-eager-smart-turtle"
```

### `WithNumber(min, max int)`
Add a random number to the phrase (adds ~13.3 bits of entropy for 0-9999).

```go
phrase, _ := a.GenerateAcrostic("test", acrostic.WithNumber(0, 9999))
// Output: "tiny eager smart turtle 4782"
```

### `WithNumberPosition(pos NumberPosition)`
Control where the number appears:
- `NumberPositionEnd` (default)
- `NumberPositionBeginning`
- `NumberPositionRandom`

```go
phrase, _ := a.GenerateAcrostic("test",
  acrostic.WithNumber(1000, 9999),
  acrostic.WithNumberPosition(acrostic.NumberPositionBeginning),
)
// Output: "5432 tiny eager smart turtle"
```

### `WithCapitalization(mode CapitalizeMode)`
Set capitalization style:
- `CapitalizeNone` (default)
- `CapitalizeFirst` - capitalize first letter of each word
- `CapitalizeAll` - all uppercase
- `CapitalizeRandom` - randomly capitalize letters

```go
phrase, _ := a.GenerateAcrostic("test",
  acrostic.WithCapitalization(acrostic.CapitalizeFirst),
)
// Output: "Tiny Eager Smart Turtle"
```

## API

### Functions

- `NewAcrostic(adjWords, nounWords *words.Words) (*Acrostic, error)`
  - Creates a new acrostic generator. Pass `nil` for both arguments to use default word lists (recommended).

### Methods

All methods accept optional `Option` parameters for customization.

- `GenerateAcrostic(acro string, opts ...Option) (string, error)`
  - Generate a single acrostic phrase for the given word.

- `GenerateAcrostics(acro string, num int, opts ...Option) ([]string, error)`
  - Generate multiple acrostic phrases for the given word.

- `GenerateRandomAcrostic(length int, opts ...Option) (string, error)`
  - Generate a single random acrostic of the specified length.

- `GenerateRandomAcrostics(length, num int, opts ...Option) ([]string, error)`
  - Generate multiple random acrostics of the specified length.

- `GenerateRandomPhrase(phraseLen int, opts ...Option) (string, error)`
  - Generate a single random phrase with the specified number of words.

- `GenerateRandomPhrases(phraseLen, phraseNum int, opts ...Option) ([]string, error)`
  - Generate multiple random phrases with the specified number of words each.

### Option Types

- `WithSeparator(sep string) Option`
- `WithNumber(min, max int) Option`
- `WithNumberPosition(pos NumberPosition) Option`
- `WithCapitalization(mode CapitalizeMode) Option`

## Testing

Run tests with race detection:

```bash
go test -race ./...
```

Run benchmarks:

```bash
go test -bench=. ./...
```

## Word Lists

This library uses curated word lists sourced from:
- **Adjectives**: [5,069 English adjectives](https://gist.github.com/JTRNS/db8fa4aab86c8f91dae713c05534ce5f) (filtered to 4,615)
- **Nouns**: [Top 10,000 English nouns](https://github.com/david47k/top-english-wordlists) from Google ngrams (filtered to 9,380)

Word lists are filtered to include only 3-12 character words with letters only, ensuring memorable and typeable passphrases.

## License

See LICENSE file for details.

---

**Sources:**
- [EFF's New Wordlists for Random Passphrases](https://www.eff.org/deeplinks/2016/07/new-wordlists-random-passphrases)
- [EFF Large Wordlist](https://www.eff.org/files/2016/07/18/eff_large_wordlist.txt)
- [Comprehensive adjectives list on GitHub](https://gist.github.com/JTRNS/db8fa4aab86c8f91dae713c05534ce5f)
- [Top English word lists by frequency](https://github.com/david47k/top-english-wordlists)
