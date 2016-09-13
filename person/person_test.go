package person

import "testing"

func TestNothingFound(t *testing.T) {
	expected := Person{}

	actualPerson := Person{}
	actual := actualPerson.GetPerson(200)

	if expected != actual {
		t.Errorf("Expected '%s' but got '%s'", expected, actual)
	}
}

func TestPersonFound(t *testing.T) {
	expected := Person{"Luke Skywalker", "male"}

	actualPerson := Person{}
	actual := actualPerson.GetPerson(1)

	if expected != actual {
		t.Errorf("Expected 's' but got '%s'", expected, actual)
	}
}
