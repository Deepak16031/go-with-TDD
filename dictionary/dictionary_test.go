package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary(map[string]string{"test": "this is just a test"})
	t.Run("known word", func(t *testing.T) {
		actual, _ := dictionary.Search("test")
		expected := "this is just a test"
		assertStrings(t, actual, expected)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		expected := "this is an unknown key"
		if err == nil {
			t.Fatal("expected to get an error")
		}
		assertStrings(t, err.Error(), expected)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "this is a test")

	actual, _ := dictionary.Search("test")
	assertStrings(t, actual, "this is a test")
}

func assertStrings(t testing.TB, actual string, expected string) {
	t.Helper()
	if actual != expected {
		t.Errorf("expected %q actual %q", actual, expected)
	}
}
