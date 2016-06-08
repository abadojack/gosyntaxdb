package gosyntaxdb

import (
	"net/url"
	"strconv"
)

// Concept represents a concept belonging to a category/language.
type Concept struct {
	ID                int    `json:"id"`
	Name              string `json:"concept_name"`
	CategoryID        int    `json:"category_id"`
	Position          int    `json:"position"`
	LanguageID        int    `json:"language_id"`
	Search            string `json:"concept_search"`
	Permalink         string `json:"concept_permalink"`
	Description       string `json:"description"`
	Syntax            string `json:"syntax"`
	Notes             string `json:"notes"`
	Example           string `json:"example"`
	Keywords          string `json:"keywords"`
	Related           string `json:"related"`
	Documentation     string `json:"documentation"`
	LanguagePermalink string `json:"language_permalink"`
}

// GetConcepts returns an array of concepts contained in SyntaxCenter.
// By default, the array is limited to 20 concepts, sorted by ID in ascending order.
func GetConcepts(v url.Values) ([]Concept, error) {
	var concepts []Concept
	err := apiGet("/concepts", v, &concepts)

	return concepts, err
}

//SearchConcepts performs a full text search for matching concepts across the entire database.
func SearchConcepts(v url.Values) ([]Concept, error) {
	var concepts []Concept
	err := apiGet("/concepts/search", v, &concepts)

	return concepts, err
}

//GetConcept returns the concept corresponding to the unique identifier.
func GetConcept(conceptID int, v url.Values) (Concept, error) {
	var concept Concept
	err := apiGet("/concepts/"+strconv.Itoa(conceptID), v, &concept)

	return concept, err
}

//GetLanguageConcept returns the concept corresponding to the language permalink and concept permalink combination.
func GetLanguageConcept(languagePermalink, conceptPermalink string, v url.Values) (Concept, error) {
	var concept Concept
	err := apiGet("/languages/"+languagePermalink+"/concepts/"+conceptPermalink, v, &concept)

	return concept, err
}
