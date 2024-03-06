package api

import (
	"Assignment1/structs"
	"Assignment1/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// GetGutendexTotalBooks gets the total amount of books in gutendex from the Gutendex API
func GetGutendexTotalBooks(rw http.ResponseWriter) int {
	resp, err := http.Get(utils.GutendexAPI)
	if err != nil {
		http.Error(rw, "Failed to get bookcount from Gutendex API", http.StatusBadRequest)
	}
	decoder := json.NewDecoder(resp.Body)
	if err != nil {
		fmt.Println("Error: could not read json file", err)
	}
	var totalBooksCount structs.TotalBooksCount
	if err := decoder.Decode(&totalBooksCount); err != nil {
		fmt.Println("Error: could not encode json file", err)
	}
	return totalBooksCount.TotalBooks
}

// HandlerGetGutendex handles the /gutendex endpoint
func HandlerGetGutendex(rw http.ResponseWriter, r *http.Request) {
	languageCode := r.URL.Query().Get("languages")

	if languageCode == "" {
		fmt.Println("no language code found in query")
	} else {
		for _, languageCode := range strings.Split(languageCode, ",") {
			PostGutenDex(rw, StructResults(rw, languageCode))
		}
	}

}

// PostGutenDex posts the results of the books in the given language code
func PostGutenDex(rw http.ResponseWriter, resultBookLanguage structs.BookLanguage) {
	rw.Header().Set("Books", utils.BOOKCOUNTPATH)
	err := json.NewEncoder(rw).Encode(resultBookLanguage)
	if err != nil {
		http.Error(rw, "Error during JSON encoding (Error: "+err.Error()+")", http.StatusInternalServerError)
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
func GetGutendex(rw http.ResponseWriter, next string, resultList *structs.GutendexListResult, authorList *structs.AuthorsList) {
	resp, err := http.Get(next)
	if err != nil {
		http.Error(rw, "Failed to get request from Gutendex API", http.StatusBadRequest)
	}
	var guntendexResult structs.GutendexResult
	if err := json.NewDecoder(resp.Body).Decode(&guntendexResult); err != nil {
		log.Fatalf("Error decoding JSON response: %v", err)
	}
	resultList.Results = append(resultList.Results, guntendexResult)
	for _, book := range guntendexResult.Books {
		authorList.AuthorsList = append(authorList.AuthorsList, book.AuthorsList...)
	}
	if guntendexResult.Next != "" {
		GetGutendex(rw, guntendexResult.Next, resultList, authorList)
	}
}
