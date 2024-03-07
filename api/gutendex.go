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
	languageCode := r.URL.Query().Get("language")
	if languageCode == "" {
		http.Error(rw, "no language code found in query", http.StatusBadRequest)
	} else {
		var bookList []structs.BookLanguage
		for _, languageCode := range strings.Split(languageCode, ",") {
			var book = StructResults(rw, languageCode)
			if book.Books != 0 && book.TotalAuthors != 0 && book.Fraction != 0 {
				bookList = append(bookList, book)
			}
		}
		utils.PostResponse(rw, bookList, "Bookcount")
	}

}

// StructResults creates a new BookLanguage struct
func StructResults(rw http.ResponseWriter, languageCode string) structs.BookLanguage {
	resultList := new([]structs.GutendexResult)
	authorList := new(structs.AuthorsList)
	GetGutendex(rw, utils.GutendexAPI+"?languages="+languageCode, resultList, authorList)
	var books = structs.BookLanguage{
		Language:     languageCode,
		Books:        (*resultList)[0].BookCount,
		TotalAuthors: utils.CountAuthors(authorList),
		Fraction:     utils.FracBooks((*resultList)[0], GetGutendexTotalBooks(rw))}
	return books

}

// GetGutendex gets the books from the Gutendex API
func GetGutendex(rw http.ResponseWriter, query string, resultList *[]structs.GutendexResult, authorList *structs.AuthorsList) {
	var gutendexResult = utils.GetResponse(rw, query)
	if gutendexResult != nil || gutendexResult != "" {
		result := gutendexResult.(structs.GutendexResult)
		*resultList = append(*resultList, result)
		for _, book := range result.Books {
			authorList.AuthorsList = append(authorList.AuthorsList, book.AuthorsList...)
		}
		if result.Next != "" {
			GetGutendex(rw, result.Next, resultList, authorList)
		}
	}
}
