package maps

import "testing"

func TestSearch(t *testing.T) {

	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, err, ErrNotFound)
	})

}

// assertStrings checks if both the "got" and "want" string matches
func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got an error %q want %q", got, want)
	}
}
