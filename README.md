# gosyntaxdb

[![Build Status](https://travis-ci.org/abadojack/gosyntaxdb.svg?branch=master)](https://travis-ci.org/abadojack/gosyntaxdb)  [![GoDoc](https://godoc.org/github.com/abadojack/gosyntaxdb?status.png)](http://godoc.org/github.com/abadojack/gosyntaxdb)

gosyntaxdb is a simple, transparent Go package for accessing the [SyntaxDB](https://syntaxdb.com/api/v1) API.

Successful API queries return native Go structs that can be used immediately, with no need for type assertions.

gosyntaxdb implements the endpoints defined in the documentation: https://syntaxdb.com/api/v1.
More detailed information about the behavior of each particular endpoint can be found at the official documentation of the API.

Examples
-------------

## Installation

	$ go get -u github.com/abadojack/gosyntaxdb

## Usage

```Go
  import "github.com/abadojack/gosyntaxdb"
```

## Authentication

No Authentication is required for the usage of the API.

## Queries

Executing queries is simple.
```Go
  languages, _ := gosyntaxdb.GetLanguages(nil)
  for _, language := range languages {
    fmt.Println(language.Name)
  }
```

Cetain endpoints allow separate optional parameter; if desired, these can be passed as the final parameter.
```Go
  v := url.Values{}
  v.Set("sort", "-language_version")
  language, err := gosyntaxdb.GetLanguages(v)
```

## Licence
gosyntaxdb is free software licensed under the WTFPL – Do What the Fuck You Want to Public License.
Details provided in the LICENCE file.
