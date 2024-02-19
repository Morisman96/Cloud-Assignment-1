package book_json

import (
	_ "encoding/json"
	_ "fmt"
	_ "log"
	_ "net/http"
	_ "os"
	_ "strconv"
)

type BookLanguage struct {
	Language     string  `json:"languages"`
	Books        int     `json:"count"`
	TotalAuthors int     `json:"authors"`
	Fraction     float32 `json:"fraction"`
}

type BookLangReaders struct {
	Country    string
	IsoCode    int
	Books      int
	Authors    int
	Readership int
}
type GutendexListResult struct {
	Results []GutendexResult `json:"results"`
}

func newGutendexListResult() *GutendexListResult {
	gutendexListResult := new(GutendexListResult)
	return gutendexListResult
}

type GutendexResult struct {
	Next      string        `json:"next"`
	Previous  string        `json:"previous"`
	Books     []AuthorsList `json:"results"`
	BookCount int           `json:"count"`
}

type AuthorsList struct {
	AuthorsList []Author `json:"authors"`
}

func newAllAuthorsList() *AuthorsList {
	authorsList := new(AuthorsList)
	return authorsList
}

type TotalBooksCount struct {
	TotalBooks int `json:"count"`
}

type Author struct {
	Author string `json:"name"`
}
