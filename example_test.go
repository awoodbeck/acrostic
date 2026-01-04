package acrostic_test

import (
	"fmt"
	"log"

	"github.com/awoodbeck/acrostic"
)

func ExampleAcrostic_GenerateAcrostic() {
	a, err := acrostic.NewAcrostic(nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	phrase, err := a.GenerateAcrostic("test")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Generated acrostic for 'test': %s\n", phrase)
}

func ExampleAcrostic_GenerateAcrostics() {
	a, err := acrostic.NewAcrostic(nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	phrases, err := a.GenerateAcrostics("go", 3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Generated %d acrostics for 'go'\n", len(phrases))
	for i, phrase := range phrases {
		fmt.Printf("%d: %s\n", i+1, phrase)
	}
}

func ExampleAcrostic_GenerateRandomAcrostic() {
	a, err := acrostic.NewAcrostic(nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	phrase, err := a.GenerateRandomAcrostic(4)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Generated random 4-word acrostic: %s\n", phrase)
}

func ExampleAcrostic_GenerateRandomPhrase() {
	a, err := acrostic.NewAcrostic(nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	phrase, err := a.GenerateRandomPhrase(5)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Generated random 5-word phrase: %s\n", phrase)
}

func ExampleNewAcrostic() {
	// Create a new acrostic generator with default word lists
	a, err := acrostic.NewAcrostic(nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Generate a simple acrostic
	phrase, err := a.GenerateAcrostic("hello")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Acrostic for 'hello': %s\n", phrase)
}

func ExampleWithSeparator() {
	a, err := acrostic.NewAcrostic(nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Generate with dash separator
	phrase, err := a.GenerateAcrostic("test", acrostic.WithSeparator("-"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("With dashes: %s\n", phrase)
}

func ExampleWithNumber() {
	a, err := acrostic.NewAcrostic(nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Generate with a number suffix (0-9999)
	phrase, err := a.GenerateAcrostic("test", acrostic.WithNumber(0, 9999))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("With number: %s\n", phrase)
}

func ExampleWithCapitalization() {
	a, err := acrostic.NewAcrostic(nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Generate with first letter capitalized
	phrase, err := a.GenerateAcrostic("test", acrostic.WithCapitalization(acrostic.CapitalizeFirst))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Capitalized: %s\n", phrase)
}

func ExampleAcrostic_GenerateAcrostic_multipleOptions() {
	a, err := acrostic.NewAcrostic(nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Combine multiple options for a secure passphrase
	phrase, err := a.GenerateAcrostic("golang",
		acrostic.WithSeparator("-"),
		acrostic.WithNumber(100, 999),
		acrostic.WithCapitalization(acrostic.CapitalizeFirst),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Secure passphrase: %s\n", phrase)
}
