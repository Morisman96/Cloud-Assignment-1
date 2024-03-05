package api

import (
	"Assignment1/book_json"
	"Assignment1/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func GetGutendexTotalBooks(rw http.ResponseWriter) int {
	resp, err := http.Get(utils.GutendexAPI)
	if err != nil {
		http.Error(rw, "Failed to get bookcount from Gutendex API", http.StatusBadRequest)
	}
	decoder := json.NewDecoder(resp.Body)
	if err != nil {
		fmt.Println("Error: could not read json file", err)
	}
	var totalBooksCount book_json.TotalBooksCount
	if err := decoder.Decode(&totalBooksCount); err != nil {
		fmt.Println("Error: could not encode json file", err)
	}
	return totalBooksCount.TotalBooks
}

func HandlerGetGutendex(rw http.ResponseWriter, r *http.Request) {
	languageCode := r.URL.Query().Get("languages")

	if languageCode == "" {
		fmt.Println("no language code found in query")
	} else {
		PostGutenDex(rw, CreateResults(rw, languageCode))
	}

}

func PostGutenDex(rw http.ResponseWriter, resultBookLanguage book_json.BookLanguage) {
	rw.Header().Set("Books", utils.BOOKCOUNTPATH)
	err := json.NewEncoder(rw).Encode(resultBookLanguage)
	if err != nil {
		http.Error(rw, "Error during JSON encoding (Error: "+err.Error()+")", http.StatusInternalServerError)
	}
}

func CreateResults(rw http.ResponseWriter, languageCode string) book_json.BookLanguage {
	resultList := new(book_json.GutendexListResult)
	authorList := new(book_json.AuthorsList)
	substrings := strings.Split(languageCode, ",")
	for _, languageCode := range substrings {
		GetGutendex(rw, utils.GutendexAPI+"?languages="+languageCode, resultList, authorList)
	}
	var books = book_json.BookLanguage{Language: languageCode, Books: resultList.Results[0].BookCount, TotalAuthors: utils.CountAuthors(authorList), Fraction: utils.FracBooks(resultList.Results[0], GetGutendexTotalBooks(rw))}
	return books
}

func GetGutendex(rw http.ResponseWriter, next string, resultList *book_json.GutendexListResult, authorList *book_json.AuthorsList) {
	resp, err := http.Get(next)
	if err != nil {
		http.Error(rw, "Failed to get request from Gutendex API", http.StatusBadRequest)
	}
	var guntendexResult book_json.GutendexResult
	if err := json.NewDecoder(resp.Body).Decode(&guntendexResult); err != nil {
		log.Fatalf("Error decoding JSON response: %v", err)
	}
	resultList.Results = append(resultList.Results, guntendexResult)
	for _, book := range guntendexResult.Books {
		authorList.AuthorsList = append(authorList.AuthorsList, book.AuthorsList...)
	}
	if guntendexResult.Next != "" {
		fmt.Println("next: ", guntendexResult.Next)
		GetGutendex(rw, guntendexResult.Next, resultList, authorList)
	}
}
