package api

import (
	"Assignment1/structs"
	"Assignment1/utils"
	"net/http"
	"strings"
)

// GetGutendexTotalBooks gets the total amount of books in gutendex from the Gutendex API
func GetGutendexTotalBooks(rw http.ResponseWriter) int {
	var totalBooksCount = utils.GetResponse(rw, utils.GutendexAPI).(structs.GutendexResult)
	return totalBooksCount.BookCount
}

// HandlerGetGutendex handles the /bookcount endpoint
func HandlerGetGutendex(rw http.ResponseWriter, r *http.Request) {
	languageCode := r.URL.Query().Get("languages")
	if languageCode == "" {
		http.Error(rw, "no language code found in query", http.StatusBadRequest)
	} else {
		for _, languageCode := range strings.Split(languageCode, ",") {
			utils.PostResponse(rw, StructResults(rw, languageCode), "Bookcount")
		}
	}

}

// StructResults creates a new BookLanguage struct
func StructResults(rw http.ResponseWriter, languageCode string) structs.BookLanguage {
	resultList := new(structs.GutendexListResult)
	authorList := new(structs.AuthorsList)
	GetGutendex(rw, utils.GutendexAPI+"?languages="+languageCode, resultList, authorList)
	var books = structs.BookLanguage{Language: languageCode, Books: resultList.Results[0].BookCount, TotalAuthors: utils.CountAuthors(authorList), Fraction: utils.FracBooks(resultList.Results[0], GetGutendexTotalBooks(rw))}
	return books
}

// GetGutendex gets the books from the Gutendex API
func GetGutendex(rw http.ResponseWriter, query string, resultList *structs.GutendexListResult, authorList *structs.AuthorsList) {
	var guntendexResult = utils.GetResponse(rw, query).(structs.GutendexResult)
	resultList.Results = append(resultList.Results, guntendexResult)
	for _, book := range guntendexResult.Books {
		authorList.AuthorsList = append(authorList.AuthorsList, book.AuthorsList...)
	}
	if guntendexResult.Next != "" {
		GetGutendex(rw, guntendexResult.Next, resultList, authorList)
	}
}
