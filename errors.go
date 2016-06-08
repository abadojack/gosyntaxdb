package gosyntaxdb

import (
	"fmt"
	"net/url"
)

//APIErrorResponse represents model schema of an error returned by the API.
type APIErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

//APIError has an APIErrorResponse and the URL associated with the error.
type APIError struct {
	APIErrorResponse APIErrorResponse
	StatusCode       int
	URL              *url.URL
}

func (e APIError) Error() string {
	return fmt.Sprintf("Get %s returned status %d : %s", e.URL, e.StatusCode, e.APIErrorResponse.Message)
}
