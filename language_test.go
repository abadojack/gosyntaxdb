package gosyntaxdb

import (
	"math/rand"
	"net/url"
	"strconv"
	"testing"
	"time"
)

var permalinks = []string{"java", "c", "cpp", "csharp", "python", "ruby", "javascript", "swift", "go"}

func Test_GetLanguages(t *testing.T) {
	sort := []string{"language_name", "position", "language_permalink", "language_version",
		"-language_name", "-position", "-language_permalink", "-language_version"}
	v := url.Values{}

	r := rand.New(rand.NewSource(time.Now().Unix()))

	v.Set("sort", sort[r.Intn(len(sort)-1)])

	_, err := GetLanguages(v)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_GetLanguage(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	_, err := GetLanguage(permalinks[r.Intn(len(permalinks)-1)], nil)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_GetLanguageCategories(t *testing.T) {
	sort := []string{"category_name", "category_search", "language_id", "language_permalink", "position",
		"-category_name", "-category_search", "-language_id", "-language_permalink", "-position"}
	r := rand.New(rand.NewSource(time.Now().Unix()))
	v := url.Values{}

	v.Set("limit", strconv.Itoa(1+r.Intn(20)))
	v.Set("sort", sort[r.Intn(len(sort)-1)])

	_, err := GetLanguageCategories(permalinks[r.Intn(len(permalinks)-1)], v)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_GetLanguageCategoryConcepts(t *testing.T) {
	sort := []string{"id", "concept_name", "category_id", "position", "language_id", "language_permalink",
		"concept_search", "concept_permalink", "description", "syntax", "notes", "examples", "keywords", "related", "documentation",
		"-id", "-concept_name", "-category_id", "-position", "-language_id", "-language_permalink",
		"-concept_search", "-concept_permalink", "-description", "-syntax", "-notes", "-examples", "-keywords", "-related", "-documentation"}

	categoryID := []int{50, 51, 52, 53}

	r := rand.New(rand.NewSource(time.Now().Unix()))

	v := url.Values{}
	v.Set("limit", strconv.Itoa(1+r.Intn(20)))
	v.Set("sort", sort[r.Intn(len(sort)-1)])

	_, err := GetLanguageCategoryConcepts("go", categoryID[r.Intn(len(categoryID))], v)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_GetLanguageConcepts(t *testing.T) {
	sort := []string{"id", "concept_name", "category_id", "position", "language_id", "language_permalink",
		"concept_search", "concept_permalink", "description", "syntax", "notes", "examples", "keywords", "related", "documentation",
		"-id", "-concept_name", "-category_id", "-position", "-language_id", "-language_permalink",
		"-concept_search", "-concept_permalink", "-description", "-syntax", "-notes", "-examples", "-keywords", "-related", "-documentation"}

	r := rand.New(rand.NewSource(time.Now().Unix()))

	v := url.Values{}
	v.Set("limit", strconv.Itoa(1+r.Intn(20)))
	v.Set("sort", sort[r.Intn(len(sort)-1)])

	_, err := GetLanguageConcepts(permalinks[r.Intn(len(permalinks)-1)], v)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_SearchLanguageConcepts(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	v := url.Values{}
	v.Set("q", "array")
	_, err := SearchLanguageConcepts(permalinks[r.Intn(len(permalinks)-1)], v)

	if err != nil {
		t.Fatal(err)
	}
}
