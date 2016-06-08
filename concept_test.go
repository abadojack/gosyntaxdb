package gosyntaxdb

import (
	"net/url"
	"testing"
)

// TODO: Write better tests.

func Test_GetConcepts(t *testing.T) {
	c, err := GetConcepts(nil)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(c)
	}
}

func Test_SearchConcepts(t *testing.T) {
	c, err := SearchConcepts(url.Values{"q": {"arrays"}})
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(c)
	}
}

func Test_GetConcept(t *testing.T) {
	c, err := GetConcept(52, nil)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(c)
	}
}

func Test_GetLanguageConcept(t *testing.T) {
	c, err := GetLanguageConcept("go", "arrays", nil)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(c)
	}
}
