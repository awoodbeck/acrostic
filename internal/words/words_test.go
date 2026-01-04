package words

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewWords(t *testing.T) {
	t.Parallel()

	t.Run("nil buffer", func(t *testing.T) {
		t.Parallel()

		_, err := NewWords(nil)
		require.ErrorContains(t, err, "nil buffer")
	})

	t.Run("empty buffer", func(t *testing.T) {
		t.Parallel()

		_, err := NewWords(&bytes.Buffer{})
		require.ErrorContains(t, err, "empty buffer")
	})

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		_, err := NewWords(bytes.NewBuffer(testWords))
		require.NoError(t, err)
	})
}

func TestWordCount(t *testing.T) {
	t.Parallel()

	w, err := NewWords(bytes.NewBuffer(testWords))
	require.NoError(t, err)

	actual, err := w.wordCount()
	require.NoError(t, err)
	assert.Equal(t, int64(testWordsLen), actual)
}

func TestGetRandomWord(t *testing.T) {
	t.Parallel()

	w, err := NewWords(bytes.NewBuffer(testWords))
	require.NoError(t, err)

	t.Run("nonexistent key", func(t *testing.T) {
		t.Parallel()

		_, err := w.RandomWord('z')
		require.ErrorContains(t, err, "no words found")
	})

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		r, err := w.RandomWord('a')
		require.NoError(t, err)
		assert.True(t, bytes.Contains(testWords, []byte(r)))
	})
}

func TestGetRandomKey(t *testing.T) {
	t.Parallel()

	t.Run("key found", func(t *testing.T) {
		t.Parallel()

		w, err := NewWords(bytes.NewBuffer(testWords))
		require.NoError(t, err)

		r, err := w.RandomKey()
		require.NoError(t, err)
		assert.Contains(t, testKeys, r)
	})

	t.Run("keys initialized on construction", func(t *testing.T) {
		t.Parallel()

		w, err := NewWords(bytes.NewBuffer(testWords))
		require.NoError(t, err)

		// Keys should be initialized during construction
		assert.NotEmpty(t, w.keys)
		assert.Len(t, w.keys, len(testKeys))
	})

	t.Run("no keys", func(t *testing.T) {
		t.Parallel()

		w, err := NewWords(bytes.NewBuffer(testWords))
		require.NoError(t, err)

		w.words = make(map[byte][]string)
		w.keys = []byte{}
		_, err = w.RandomKey()
		require.ErrorContains(t, err, "word list is empty")
	})
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

func BenchmarkNewWords(b *testing.B) {
	for range b.N {
		_, err := NewWords(bytes.NewBuffer(testWords))
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkRandomKey(b *testing.B) {
	w, err := NewWords(bytes.NewBuffer(testWords))
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for range b.N {
		_, err := w.RandomKey()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkRandomWord(b *testing.B) {
	w, err := NewWords(bytes.NewBuffer(testWords))
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for range b.N {
		_, err := w.RandomWord('a')
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkEntropy(b *testing.B) {
	w, err := NewWords(bytes.NewBuffer(testWords))
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for range b.N {
		_, err := w.Entropy(4)
		if err != nil {
			b.Fatal(err)
		}
	}
}
