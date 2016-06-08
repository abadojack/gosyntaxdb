package gosyntaxdb

import (
	"encoding/json"
	"net/http"
	"net/url"
)

// BaseURL is the API URL for version 1 of the syntaxdb API.
const BaseURL = "https://syntaxdb.com/api/v1"

// apiGet issues a GET request to the SyntaxDB API and decodes the JSON response to data.
func apiGet(path string, form url.Values, data interface{}) error {
	var urlStr string
	if form != nil {
		urlStr = BaseURL + path + "?" + form.Encode()
	} else {
		urlStr = BaseURL + path
	}

	resp, err := http.Get(urlStr)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return decodeResponse(resp, data)
}

// decodeResponse decodes the JSON response from the SyntaxDB API.
func decodeResponse(resp *http.Response, data interface{}) error {
	if resp.StatusCode == 200 {
		return json.NewDecoder(resp.Body).Decode(data)
	}

	var aer APIErrorResponse

	//TODO: Not ignore the error that might be returned by this.
	json.NewDecoder(resp.Body).Decode(&aer)

	return APIError{
		URL:              resp.Request.URL,
		StatusCode:       resp.StatusCode,
		APIErrorResponse: aer,
	}
}
