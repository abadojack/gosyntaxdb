//Package gosyntaxdb provides structs and functions for accessing version 1
//of SyntaxDB API.
//
//Successful API queries return native Go structs that can be used immediately,
//with no need for type assertions.
//
//No authentification is required for accessing the API.
//
//Queries
//
//Executing queries is simple.
//
//  languages, _ := gosyntaxdb.GetLanguages(nil)
//  for _, language := range languages {
//    fmt.Println(language.Name)
//  }
//
//Cetain endpoints allow separate optional parameter; if desired, these can be passed as the final parameter.
//
//  v := url.Values{}
//  v.Set("sort", "-1")
//  language, err := gosyntaxdb.GetLanguage("go", v)
//
//Endpoints
//
//gosyntaxdb implements all the endpoints defined in the SyntaxDB API documentation : https://syntaxdb.com/api/v1
//
//More detailed information about the behavior of each particular endpoint can be found at the official SyntaxDB API documentation.
package gosyntaxdb
