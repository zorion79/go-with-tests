package maps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assert.Equal(t, want, got)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		assert.Error(t, err)
		assert.Equal(t, ErrNotFound, err)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}

		_ = dictionary.Add("test", "this is just a test")

		want := "this is just a test"
		got, err := dictionary.Search("test")

		assert.NoError(t, err)
		assert.Equal(t, want, got)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, "new test")
		assert.Error(t, err)
		assert.Equal(t, ErrWordExist, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		newDefinition := "new definition"
		err := dictionary.Update(word, newDefinition)
		assert.NoError(t, err)

		search, err := dictionary.Search(word)

		assert.NoError(t, err)
		assert.Equal(t, newDefinition, search)
	})

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		definition := "new definition"

		err := dictionary.Update("test", definition)

		assert.Error(t, err)
		assert.Equal(t, ErrWordDoesNotExist, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("word exists", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{word: "test definition"}

		_ = dictionary.Delete(word)

		_, err := dictionary.Search(word)
		assert.Equal(t, ErrNotFound, err)
	})

	t.Run("word does not exist", func(t *testing.T) {
		dictionary := Dictionary{}

		err := dictionary.Delete("test")
		assert.Equal(t, ErrWordDoesNotExist, err)
	})
}
