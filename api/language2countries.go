package api

import (
	"Assignment1/structs"
	"Assignment1/utils"
	"net/http"
	_ "net/url"
	_ "path"
	"regexp"
	"strconv"
)

// GetRestCountries gets the population of the countries from the RestCountries API
func GetRestCountries(rw http.ResponseWriter, list *[]structs.Language2Countries, limit int) {
	for i := 0; i < limit && i < len(*list); i++ {
		country := (*list)[i]
		var population = utils.GetResponse(rw, utils.RESTCountriesAPI+country.Country).([]structs.RestCountry)
		structs.UpdateReadership(*list, population[0].Population, i)
	}
}

// HandlerGetLanguage2countries handles the /language2countries endpoint
func HandlerGetLanguage2countries(rw http.ResponseWriter, r *http.Request) {
	languageCode := utils.GetUrlPath(r.URL.String(), "readership")
	var language2countriesList = utils.GetResponse(rw, utils.Language2CountriesAPI+languageCode)
	if language2countriesList != nil {
		countries := language2countriesList.([]structs.Language2Countries)
		var readershipResults []structs.BookLangReaders
		var book = StructResults(rw, languageCode)
		if r.URL.Query().Get("limit") != "" && regexp.MustCompile(utils.REGEXPOSETIVINTEGER).MatchString(r.URL.Query().Get("limit")) {
			limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
			GetRestCountries(rw, &countries, limit)
			readershipResults = CreateReadership(&countries, book, limit)
		} else {
			GetRestCountries(rw, &countries, len(countries))
			readershipResults = CreateReadership(&countries, book, len(countries))
		}
		utils.PostResponse(rw, readershipResults, "Readership")
	}
}

// CreateReadership creates a new BookLanguage struct
func CreateReadership(language2countriesList *[]structs.Language2Countries, book structs.BookLanguage, limit int) []structs.BookLangReaders {
	var readershipResults []structs.BookLangReaders
	for i := 0; i < limit && i < len(*language2countriesList); i++ {
		country := (*language2countriesList)[i]
		readershipResults = append(readershipResults, structs.BookLangReaders{Country: country.Country, IsoCode: country.IsoCode, Books: book.Books, Authors: book.TotalAuthors, Readership: country.Readership})
	}
	return readershipResults
}
