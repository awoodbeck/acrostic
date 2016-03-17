package acrostic

import (
	"bytes"
	"testing"
)

func TestNewWords(t *testing.T) {
	t.Parallel()

	var buf *bytes.Buffer

	// Test for ErrNilBuffer.
	_, err := NewWords(buf)
	switch {
	case err == nil:
		t.Error("expected nil buffer error; actual error was nil")
	case err != ErrNilBuffer:
		t.Error("expected nil buffer error; actual error was: ", err)
	}

	// Test for ErrEmptyBuffer.
	buf = &bytes.Buffer{}
	_, err = NewWords(buf)
	switch {
	case err == nil:
		t.Error("expected empty buffer error; actual error was nil")
	case err != ErrEmptyBuffer:
		t.Error("expected empty buffer error; actual error was: ", err)
	}

	// Test for proper initialization.
	buf = bytes.NewBuffer(testWords)
	_, err = NewWords(buf)
	if err != nil {
		t.Error("expected nil error; actual error was:", err)
	}
}

func TestWordCount(t *testing.T) {
	t.Parallel()

	buf := bytes.NewBuffer(testWords)
	w, err := NewWords(buf)
	if err != nil {
		t.Fatal(err)
	}

	if actual := w.WordCount(); actual != testWordsLen {
		t.Errorf("expected word count = %d; actual word count = %d", testWordsLen, actual)
	}
}

func TestGetRandomWord(t *testing.T) {
	t.Parallel()

	buf := bytes.NewBuffer(testWords)
	w, err := NewWords(buf)
	if err != nil {
		t.Fatal(err)
	}

	// Test for ErrNonexistentKey.
	_, err = w.GetRandomWord('z')
	switch {
	case err == nil:
		t.Error("expected empty buffer error; actual error was nil")
	case err != ErrNonexistentKey:
		t.Error("expected empty buffer error; actual error was:", err)
	}

	r, err := w.GetRandomWord('a')
	if err != nil {
		t.Fatal(err)
	}

	if bytes.Contains(testWords, []byte(r)) == false {
		t.Errorf("unable to find random word %q in the test word list", r)
	}
}

func TestGetRandomKey(t *testing.T) {
	t.Parallel()

	buf := bytes.NewBuffer(testWords)
	w, err := NewWords(buf)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := w.GetRandomKey(); err != nil {
		t.Error("expected nil error; actual error was:", err)
	}

	// Test for empty keys slice.
	w.keys = []byte{}
	if _, err := w.GetRandomKey(); err != nil {
		t.Error("expected nil error; actual error was:", err)
	}
	if len(w.keys) == 0 {
		t.Error("expected length of keys > 0; actual length = 0")
	}

	// Test for ErrNoKeys.
	w.words = make(map[byte][]string)
	w.keys = []byte{}
	_, err = w.GetRandomKey()
	switch {
	case err == nil:
		t.Error("expected no keys error; actual error was nil")
	case err != ErrNoKeys:
		t.Error("expected no keys error; actual error was:", err)
	}
}

// Test word list with various potential format edge cases mixed in.
var testWordsLen = 60
var testWords = []byte(`
aardvark
aardwolf
abalone
abyssiniancat
abyssiniangroundhornbill
acaciarat
achillestang
acornbarnacle
acornweevil
acornwoodpecker
acouchi
blowfish
bluebird
bluebottle
bluebottlejellyfish
        bluebreastedkookaburra
bluefintuna
bluefish
bluegill
bluejay
cow
cowbird
cowrie
coyote
coypu
crab


crane
cranefly
crayfish
creature
cricket
crocodile
crossbill
crow
dove
dowitcher
drafthorse
dragon
dragonfly
Drake
drever
    dromaeosaur
dromedary
drongo
duck
duckbillcat
duckbillplatypus
duckling

elephant
elephantbeetle
elephantseal
ELK
elkhound
elver
emeraldtreeskink
emperorpenguin
emperorshrimp
emU
equestrian
equine`)

// Slice of keys we expect to find in the above word list.
var testKeys = []byte{'a', 'b', 'c', 'd', 'e'}
