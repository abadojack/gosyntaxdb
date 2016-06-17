package gosyntaxdb

import (
	"math/rand"
	"net/url"
	"strconv"
	"testing"
	"time"
)

func Test_GetConcepts(t *testing.T) {
	sort := []string{"id", "concept_name", "category_id", "position", "language_id", "language_permalink",
		"concept_search", "concept_permalink", "description", "syntax", "notes", "examples", "keywords", "related", "documentation",
		"-id", "-concept_name", "-category_id", "-position", "-language_id", "-language_permalink",
		"-concept_search", "-concept_permalink", "-description", "-syntax", "-notes", "-examples", "-keywords", "-related", "-documentation"}

	r := rand.New(rand.NewSource(time.Now().Unix()))

	v := url.Values{}
	v.Set("limit", strconv.Itoa(1+r.Intn(20)))
	v.Set("sort", sort[r.Intn(len(sort)-1)])

	concepts, err := GetConcepts(v)
	if err != nil {
		t.Fatal("GetConcepts", err)
	}

	concept, err := GetConcept(concepts[0].ID, nil)
	if err != nil {
		t.Fatal("GetConcept", err)
	}

	_, err = GetLanguageConcept(concept.LanguagePermalink, concept.Permalink, nil)
	if err != nil {
		t.Fatal("GetLanguageConcept", err)
	}
}

func Test_SearchConcepts(t *testing.T) {
	q := []string{"arrays", "variables", "string"}
	r := rand.New(rand.NewSource(time.Now().Unix()))

	v := url.Values{}
	v.Set("q", q[r.Intn(len(q)-1)])

	_, err := SearchConcepts(url.Values{"q": {"arrays"}})
	if err != nil {
		t.Fatal(err)
	}
}
