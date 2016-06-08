package gosyntaxdb

// Category represents a category corresponding to a programming language.
type Category struct {
	ID                int    `json:"id"`
	Name              string `json:"category_name"`
	Search            string `json:"category_search"`
	LanguageID        int    `json:"language_id"`
	Position          int    `json:"position"`
	LanguagePermalink string `json:"language_permalink"`
}
