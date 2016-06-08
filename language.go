package gosyntaxdb

import (
	"net/url"
	"strconv"
)

// Language represents a programming language.
type Language struct {
	ID        int    `json:"id"`
	Name      string `json:"language_name"`
	Position  int    `json:"position"`
	Version   string `json:"language_version"`
	Permalink string `json:"language_permalink"`
}

// GetLanguages returns all languages available on SyntaxCenter, sorted by position by default.
func GetLanguages(v url.Values) ([]Language, error) {
	var languages []Language
	err := apiGet("/languages", v, &languages)

	return languages, err
}

// GetLanguage returns the language corresponding to the permalink along with its information.
func GetLanguage(languagePermalink string, v url.Values) (Language, error) {
	var language Language
	err := apiGet("/languages/"+languagePermalink, v, &language)

	return language, err
}

//GetLanguageCategories returns all of the categories corresponding to the language.
func GetLanguageCategories(languagePermalink string, v url.Values) ([]Category, error) {
	var categories []Category
	err := apiGet("/languages/"+languagePermalink+"/categories", v, &categories)

	return categories, err
}

//GetLanguageCategoryConcepts returns all concepts belonging to the category, sorted by position by default.
func GetLanguageCategoryConcepts(languagePermalink string, categoryID int, v url.Values) ([]Concept, error) {
	var concept []Concept
	err := apiGet("/languages/"+languagePermalink+"/categories/"+strconv.Itoa(categoryID)+"/concepts", v, &concept)

	return concept, err
}

// GetLanguageConcepts returns all concepts belonging to the language, sorted by position by default.
func GetLanguageConcepts(languagePermalink string, v url.Values) ([]Concept, error) {
	var concept []Concept
	err := apiGet("/languages/"+languagePermalink+"/concepts", v, &concept)

	return concept, err
}

//SearchLanguageConcepts returns all concepts within the specified language that matches the query.
func SearchLanguageConcepts(languagePermalink string, v url.Values) ([]Concept, error) {
	var concept []Concept
	err := apiGet("/languages/"+languagePermalink+"/concepts/search", v, &concept)

	return concept, err
}
