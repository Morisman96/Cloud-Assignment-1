package api

import (
	"Assignment1/book_json"
	"Assignment1/utils"
	"encoding/json"
	"net/http"
	_ "net/url"
	_ "path"
	"regexp"
	"strconv"
)

func GetLanguage2countries(rw http.ResponseWriter, languageCode string) *[]*book_json.Language2Countries {
	resp, err := http.Get(utils.Language2CountriesAPI + languageCode)
	if err != nil {
		http.Error(rw, "Failed to get request from Gutendex API", http.StatusBadRequest)
	}
	var language2countriesList *[]*book_json.Language2Countries
	err = json.NewDecoder(resp.Body).Decode(&language2countriesList)
	if err != nil {
		http.Error(rw, "Failed to decode JSON", http.StatusBadRequest)
	}

	return language2countriesList
}

func GetRestCountries(rw http.ResponseWriter, list *[]*book_json.Language2Countries, limit int) {
	for i := 0; i < limit && i < len(*list); i++ {
		country := (*list)[i]
		resp, err := http.Get(utils.RESTCountriesAPI + country.Country)
		if err != nil {
			http.Error(rw, "Failed to get request from RestCountries API", http.StatusBadRequest)
		}
		var population []book_json.RestCountry
		err = json.NewDecoder(resp.Body).Decode(&population)
		if err != nil {
			http.Error(rw, "Failed to decode JSON", http.StatusBadRequest)
		}
		book_json.UpdateReadership(country, population[0].Population)
	}
}

func HandlerGetLanguage2countries(rw http.ResponseWriter, r *http.Request) {
	languageCode := utils.LanguageCode(r.URL.String())
	if languageCode == "" && languageCode == "v1" {
		http.Error(rw, "no language code found in query", http.StatusBadRequest)
	} else {
		language2countriesList := new([]*book_json.Language2Countries)
		language2countriesList = GetLanguage2countries(rw, languageCode)
		var readershipResults []book_json.BookLangReaders
		var book = CreateResults(rw, languageCode)

		if r.URL.Query().Get("limit") != "" && regexp.MustCompile(utils.REGEXPOSETIVINTEGER).MatchString(r.URL.Query().Get("limit")) {
			limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
			GetRestCountries(rw, language2countriesList, limit)
			readershipResults = CreateReadership(language2countriesList, book, limit)
		} else {
			GetRestCountries(rw, language2countriesList, len(*language2countriesList))
			readershipResults = CreateReadership(language2countriesList, book, len(*language2countriesList))
		}

		err := json.NewEncoder(rw).Encode(readershipResults)
		if err != nil {
			http.Error(rw, "Error during JSON encoding (Error: "+err.Error()+")", http.StatusInternalServerError)
		}
	}
}
func CreateReadership(language2countriesList *[]*book_json.Language2Countries, book book_json.BookLanguage, limit int) []book_json.BookLangReaders {
	var readershipResults []book_json.BookLangReaders
	for i := 0; i < limit && i < len(*language2countriesList); i++ {
		country := (*language2countriesList)[i]
		readershipResults = append(readershipResults, book_json.BookLangReaders{Country: country.Country, IsoCode: country.IsoCode, Books: book.Books, Authors: book.TotalAuthors, Readership: country.Readership})
	}
	return readershipResults
}
