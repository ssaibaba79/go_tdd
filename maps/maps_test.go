package maps

import "testing"

func TestLookup(t *testing.T) {

	t.Run("Known word", func(t *testing.T) {
		dict := Dictionary{"archer": "A person who can use a bow and arrow proficiently"}
		want := "A person who can use a bow and arrow proficiently"
		got, _ := dict.Lookup("archer")
		assertString(t, dict, got, want)
	})

	t.Run("Unknown word", func(t *testing.T) {
		dict := Dictionary{"archer": "A person who can use a bow and arrow proficiently"}
		_, err := dict.Lookup("arch")
		assertError(t, dict, err, ErrWordNotFound)
	})

	t.Run("Add word", func(t *testing.T) {
		dict := Dictionary{"archer": "A person who can use a bow and arrow proficiently"}
		want := "An instrument for documenting information"
		err := dict.Add("Book", "An instrument for documenting information")
		assertError(t, dict, err, nil)
		got, _ := dict.Lookup("Book")
		assertString(t, dict,  got, want)
	})


	t.Run("Add existing word", func(t *testing.T) {
		dict := Dictionary{"archer": "A person who can use a bow and arrow proficiently"}
		err := dict.Add("archer", "A warrior who uses a bow and arrow in a war")
		assertError(t, dict, err, ErrWordExists)
	})


	t.Run("Update definition", func(t *testing.T) {
		dict := Dictionary{"archer": "A person who can use a bow and arrow proficiently"}
		want := "A warrior who uses a bow and arrow in a war"
		err := dict.Update("archer", "A warrior who uses a bow and arrow in a war")
		assertError(t, dict, err, nil)
		got, _ := dict.Lookup("archer")
		assertString(t, dict, got, want)
	})

		t.Run("Update definition for unknown word", func(t *testing.T) {
		dict := Dictionary{"archer": "A person who can use a bow and arrow proficiently"}
		err := dict.Update("Book", "An instrument for documenting information")
		assertError(t, dict, err, ErrWordNotFound)
	})


	t.Run("Delete word", func(t *testing.T) {
		dict := Dictionary{"archer": "A person who can use a bow and arrow proficiently"}
		err := dict.Delete("archer")
		assertError(t, dict, err, nil)
		_, err = dict.Lookup("archer")
		assertError(t, dict, err, ErrWordNotFound)
	})

		t.Run("Delete unknown word", func(t *testing.T) {
		dict := Dictionary{"archer": "A person who can use a bow and arrow proficiently"}
		err := dict.Delete("Book")
		assertError(t, dict,  err, ErrWordNotFound)
	})

}

func assertString(t testing.TB, dict Dictionary,  got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("%#v want:%s, got:%s",dict,  want, got)
	}
}

func assertError(t testing.TB, dict Dictionary, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("%#v got error :%q, want:%q",dict,  want, got)
	}
}
