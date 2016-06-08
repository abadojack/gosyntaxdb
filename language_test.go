package gosyntaxdb

import (
	"net/url"
	"testing"
)

// TODO: Write better tests.

func Test_GetLanguages(t *testing.T) {
	v := url.Values{}
	v.Set("sort", "-1")
	l, err := GetLanguages(v)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(l)
	}
}

func Test_GetLanguage(t *testing.T) {
	l, err := GetLanguage("go", nil)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(l)
	}
}

func Test_GetLanguageCategories(t *testing.T) {
	l, err := GetLanguageCategories("go", nil)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(l)
	}
}

func Test_GetLanguageCategoryConcepts(t *testing.T) {
	c, err := GetLanguageCategoryConcepts("go", 50, nil)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(c)
	}
}

func Test_GetLanguageConcepts(t *testing.T) {
	c, err := GetLanguageConcepts("go", nil)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(c)
	}
}

func Test_SearchLanguageConcepts(t *testing.T) {
	v := url.Values{}
	v.Set("q", "array")
	c, err := SearchLanguageConcepts("go", v)

	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(c)
	}
}
