package maps

import "testing"

/* * Assertion utilities * */

// assertStrings checks if both the "got" and "want" string matches
func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()

	if got != nil {
		t.Fatal("should not return an error:", got)
	}
}

// assertError checks if it is the wanted error
func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got an error %q want %q", got, want)
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if got != definition {
		t.Errorf("got %q want %q given", got, definition)
	}
}
