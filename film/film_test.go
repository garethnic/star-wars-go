package film

import (
	"testing"
)

func TestNothingIsFound(t *testing.T) {
	expected := Film{}

	actualFilm := Film{}
	actual := actualFilm.GetFilm(20)

	if expected != actual {
		t.Errorf("Expected '%s' but got '%s'", expected, actual)
	}
}

func TestFilmFound(t *testing.T) {
	expected := Film{Title: "A New Hope"}

	actualFilm := Film{}
	actual := actualFilm.GetFilm(1)

	if expected.GetTitle() != actual.GetTitle() {
		t.Errorf("Expected '%s' but got '%s'", expected.GetTitle(), actual.GetTitle())
	}
}
