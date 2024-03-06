package structs

import (
	_ "encoding/json"
	_ "fmt"
	_ "log"
	_ "net/http"
	_ "os"
	_ "strconv"
	"time"
)

// BookLang
type BookLanguage struct {
	Language     string  `json:"languages"`
	Books        int     `json:"count"`
	TotalAuthors int     `json:"authors"`
	Fraction     float32 `json:"fraction"`
}

type BookLangReaders struct {
	Country    string `json:"Official_Name"`
	IsoCode    string `json:"ISO3166_1_Alpha_2"`
	Books      int    `json:"count"`
	Authors    int    `json:"authors"`
	Readership int    `json:"population"`
}
type GutendexListResult struct {
	Results []GutendexResult `json:"results"`
}

/*
func newGutendexListResult() *GutendexListResult {
	gutendexListResult := new(GutendexListResult)
	return gutendexListResult
}
*/

type GutendexResult struct {
	Next      string        `json:"next"`
	Previous  string        `json:"previous"`
	Books     []AuthorsList `json:"results"`
	BookCount int           `json:"count"`
}

type AuthorsList struct {
	AuthorsList []Author `json:"authors"`
}

/*
func newAllAuthorsList() *AuthorsList {
	authorsList := new(AuthorsList)
	return authorsList
}
*/

type TotalBooksCount struct {
	TotalBooks int `json:"count"`
}

type Author struct {
	Author string `json:"name"`
}

type Language2Countries struct {
	Country    string `json:"Official_Name"`
	IsoCode    string `json:"ISO3166_1_Alpha_2"`
	Readership int    `json:"population"`
}

type RestCountry struct {
	Population int `json:"population"`
}

type Language2CountriesList struct {
	CountriesList []Language2Countries `json:"countries"`
}

type BookLangReadersList struct {
	ReaderList []BookLangReaders `json:"readrship"`
}

// UpdateReadership updates the readership variable
func UpdateReadership(Language2Countries *Language2Countries, population int) {
	Language2Countries.Readership = population
}

// Status struct for status of the service
type Status struct {
	GutendexApi           string        `json:"gutendexapi"`
	Language2countriesApi string        `json:"languageapi"`
	RestCountriesApi      string        `json:"countriesapi"`
	Version               string        `json:"version"`
	Uptime                time.Duration `json:"uptime"`
}
