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

// BookLanguage struct for the bookcount
type BookLanguage struct {
	Language     string  `json:"languages"`
	Books        int     `json:"count"`
	TotalAuthors int     `json:"authors"`
	Fraction     float32 `json:"fraction"`
}

// BookLangReaders	struct for the readership
type BookLangReaders struct {
	Country    string `json:"Official_Name"`
	IsoCode    string `json:"ISO3166_1_Alpha_2"`
	Books      int    `json:"count"`
	Authors    int    `json:"authors"`
	Readership int    `json:"population"`
}

// GutendexResult struct for the Gutendex API
type GutendexResult struct {
	Next      string        `json:"next"`
	Previous  string        `json:"previous"`
	Books     []AuthorsList `json:"results"`
	BookCount int           `json:"count"`
}

// AuthorsList struct for the GutendexResult
type AuthorsList struct {
	AuthorsList []Author `json:"authors"`
}

// Author struct for the AuthorsList
type Author struct {
	Author string `json:"name"`
}

// Language2Countries struct for the Language2Countries API
type Language2Countries struct {
	Country    string `json:"Official_Name"`
	IsoCode    string `json:"ISO3166_1_Alpha_2"`
	Readership int    `json:"population"`
}

// RestCountry struct for the RestCountries API to get the population
type RestCountry struct {
	Population int `json:"population"`
}

// UpdateReadership updates the readership variable
func UpdateReadership(Language2Countries []Language2Countries, population int, index int) {
	Language2Countries[index].Readership = population
}

// Status struct for status of the service
type Status struct {
	GutendexApi           string        `json:"gutendexapi"`
	Language2countriesApi string        `json:"languageapi"`
	RestCountriesApi      string        `json:"countriesapi"`
	Version               string        `json:"version"`
	Uptime                time.Duration `json:"uptime"`
}
