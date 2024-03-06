package utils

import (
	"Assignment1/structs"
	"strings"
	"time"
)

// CountAuthors returns the number of unique authors
func CountAuthors(result *structs.AuthorsList) int {
	occurrence := make(map[structs.Author]bool) // map to store authors as keys and boolean as value
	for _, author := range result.AuthorsList {
		occurrence[author] = true
	}
	return len(occurrence)
}

// FracBooks returns the fraction of books in a language
func FracBooks(result structs.GutendexResult, totalBooks int) float32 {
	return float32(result.BookCount) / float32(totalBooks)
}

// LanguageCode returns the language code from the URL path
func LanguageCode(url string) string {
	pathSegments := strings.Split(url, "/")
	var languageCode string
	for i, segment := range pathSegments {
		if segment == "readership" && i+1 < len(pathSegments) {
			languageCode = pathSegments[i+1]
			break
		}
	}
	return languageCode
}

// InitTimer initializes the Timer variable
func InitTimer() {
	Timer = time.Now()
}
