package maps

import (
	asserts "hello-world"
	"testing"
)

func TestSearch(t *testing.T) {
	t.Run("known word", func(t *testing.T) {
		// how to create a map: map[string]string{"test": "this is just a test"}
		dictionary := Dictionary{"test": "this is just a test"}

		actual, _ := dictionary.Search("test")
		expected := "this is just a test"

		asserts.AssertStringEquals(actual, expected, t)
	})

	t.Run("unknown word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		_, actualError := dictionary.Search("unicorns")
		expectedError := ErrorWordNotFound

		asserts.AssertErrorEquals(actualError, expectedError, t)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add("test", "this is just a test")

		actual, err := dictionary.Search("test")
		expected := "this is just a test"

		if err != nil {
			t.Fatal("should have added the word, but cannot find it")
		}

		asserts.AssertStringEquals(actual, expected, t)
	})

	t.Run("existing word", func(t *testing.T) {
		expectedDefinition := "this is just a test"
		dictionary := Dictionary{"test": expectedDefinition}

		actualDefinition, _ := dictionary.Search("test")
		actualError := dictionary.Add("test", "this is a v2 test definition")
		expectedError := ErrorWordExists

		asserts.AssertErrorEquals(actualError, expectedError, t)
		asserts.AssertStringEquals(actualDefinition, expectedDefinition, t)
	})
}

func TestUpdate(t *testing.T) {

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{word: "this is just a test"}
		expectedDefinition := "test definition v2"

		updateErr := dictionary.Update(word, expectedDefinition)
		actualDefinition, searchErr := dictionary.Search(word)

		asserts.AssertNoError(updateErr, t)
		asserts.AssertNoError(searchErr, t)
		asserts.AssertStringEquals(actualDefinition, expectedDefinition, t)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}
		expectedDefinition := "test definition v2"

		actualUpdateError := dictionary.Update(word, expectedDefinition)
		expectedUpdateError := ErrorWordDoesNotExist
		_, actualSearchErr := dictionary.Search(word)
		expectedSearchError := ErrorWordNotFound

		asserts.AssertErrorEquals(actualUpdateError, expectedUpdateError, t)
		asserts.AssertErrorEquals(actualSearchErr, expectedSearchError, t)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test definition"}

	dictionary.Delete(word)
	_, actualError := dictionary.Search(word)

	asserts.AssertErrorEquals(actualError, ErrorWordNotFound, t)
}
