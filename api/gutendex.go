package api

import (
	"Assignment1/book_json"
	"Assignment1/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

func GetGutendex(rw http.ResponseWriter, r *http.Request) {
	languageCode := r.URL.Query().Get("languages")
	if languageCode == "" {
		fmt.Println("no language code found in query")
	}
	resp, err := http.Get(utils.GutendexAPI + "?languages=" + languageCode)
	if err != nil {
		http.Error(rw, "No launguage found in gutendex.", http.StatusBadRequest)
	}
	var guntendexResult book_json.GuntendexResult
	if err := json.NewDecoder(resp.Body).Decode(&guntendexResult); err != nil {
		log.Fatalf("Error decoding JSON response: %v", err)
	}
	var books book_json.BookLanguage = book_json.BookLanguage{Language: languageCode, Books: guntendexResult.BookCount, TotalAuthors: utils.CountAuthors(guntendexResult), Fraction: utils.FracBooks(guntendexResult, GetGutendexTotalBooks(rw))}
	rw.Header().Set("Books", utils.BOOKCOUNTPATH)
	err = json.NewEncoder(rw).Encode(books)
	if err != nil {
		http.Error(rw, "Error during JSON encoding (Error: "+err.Error()+")", http.StatusInternalServerError)
	}
}
