package api

import (
	"Assignment1/book_json"
	"Assignment1/utils"
	"encoding/json"
	"net/http"
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

func GetRestCountries(rw http.ResponseWriter, list *[]*book_json.Language2Countries) {
	for _, country := range *list {

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
	languageCode := r.URL.Query().Get("languages")

	if languageCode == "" {
		http.Error(rw, "no language code found in query", http.StatusBadRequest)
	} else {
		language2countriesList := new([]*book_json.Language2Countries)
		language2countriesList = GetLanguage2countries(rw, languageCode)
		var readershipResults []book_json.BookLangReaders
		GetRestCountries(rw, language2countriesList)
		var book = CreateResults(rw, languageCode)
		for _, country := range *language2countriesList {
			readershipResults = append(readershipResults, book_json.BookLangReaders{Country: country.Country, IsoCode: country.IsoCode, Books: book.Books, Authors: book.TotalAuthors, Readership: country.Readership})
		}
		err := json.NewEncoder(rw).Encode(readershipResults)
		if err != nil {
			http.Error(rw, "Error during JSON encoding (Error: "+err.Error()+")", http.StatusInternalServerError)
		}
	}
}
