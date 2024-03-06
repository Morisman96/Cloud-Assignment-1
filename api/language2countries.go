package api

import (
	"Assignment1/structs"
	"Assignment1/utils"
	"encoding/json"
	"net/http"
	_ "net/url"
	_ "path"
	"regexp"
	"strconv"
)

// GetLanguage2countries gets the countries that speaks the given language code from the Language2Countries API.
func GetLanguage2countries(rw http.ResponseWriter, languageCode string) *[]*structs.Language2Countries {
	resp, err := http.Get(utils.Language2CountriesAPI + languageCode)
	if err != nil {
		http.Error(rw, "Failed to get request from language2CountriesAPI", http.StatusBadRequest)
	}
	var language2countriesList *[]*structs.Language2Countries
	err = json.NewDecoder(resp.Body).Decode(&language2countriesList)
	if err != nil {
		http.Error(rw, "Failed to decode JSON", http.StatusBadRequest)
	}

	return language2countriesList
}

// GetRestCountries gets the population of the countries from the RestCountries API
func GetRestCountries(rw http.ResponseWriter, list *[]*structs.Language2Countries, limit int) {
	for i := 0; i < limit && i < len(*list); i++ {
		country := (*list)[i]
		resp, err := http.Get(utils.RESTCountriesAPI + country.Country)
		if err != nil {
			http.Error(rw, "Failed to get request from RestCountries API", http.StatusBadRequest)
		}
		var population []structs.RestCountry
		err = json.NewDecoder(resp.Body).Decode(&population)
		if err != nil {
			http.Error(rw, "Failed to decode JSON", http.StatusBadRequest)
		}
		structs.UpdateReadership(country, population[0].Population)
	}
}

// HandlerGetLanguage2countries handles the /language2countries endpoint
func HandlerGetLanguage2countries(rw http.ResponseWriter, r *http.Request) {
	languageCode := utils.LanguageCode(r.URL.String())
	if languageCode == "" && languageCode == "v1" {
		http.Error(rw, "no language code found in query", http.StatusBadRequest)
	} else {
		language2countriesList := new([]*structs.Language2Countries)
		language2countriesList = GetLanguage2countries(rw, languageCode)
		var readershipResults []structs.BookLangReaders
		var book = StructResults(rw, languageCode)

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

// StructResults creates a new BookLanguage struct
func CreateReadership(language2countriesList *[]*structs.Language2Countries, book structs.BookLanguage, limit int) []structs.BookLangReaders {
	var readershipResults []structs.BookLangReaders
	for i := 0; i < limit && i < len(*language2countriesList); i++ {
		country := (*language2countriesList)[i]
		readershipResults = append(readershipResults, structs.BookLangReaders{Country: country.Country, IsoCode: country.IsoCode, Books: book.Books, Authors: book.TotalAuthors, Readership: country.Readership})
	}
	return readershipResults
}
