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
	TotalAuthors int     `json:"totalAuthors"`
	Fraction     float32 `json:"fraction"`
}

type BookLangReaders struct {
	Country    string
	IsoCode    int
	Books      int
	Authors    int
	Readership int
}

type GuntendexResult struct {
	Next      string   `json:"next"`
	Previous  string   `json:"previous"`
	BookCount int      `json:"count"`
	Books     []Result `json:"results"`
}

type Result struct {
	Authors   []Author `json:"authors"`
	Languages []string `json:"languages"`
}

type TotalBooksCount struct {
	TotalBooks int `json:"count"`
}

type Author struct {
	Author string `json:"name"`
}
