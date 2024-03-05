package utils

import (
	"Assignment1/book_json"
	"fmt"
	"strings"
)

func CountAuthors(result *book_json.AuthorsList) int {
	occurrence := make(map[book_json.Author]bool)

	// Iterate through the list and mark each element as true in the map
	for _, author := range result.AuthorsList {

		occurrence[author] = true

		// Return the count of unique elements return len(occurrence)

	}
	fmt.Println(len(occurrence))
	for k, v := range occurrence {
		fmt.Println(k, v)

	}
	return len(occurrence)
}

func FracBooks(result book_json.GutendexResult, totalBooks int) float32 {
	return float32(result.BookCount) / float32(totalBooks)
}

func LanguageCode(url string) string {

	pathSegments := strings.Split(url, "/")

	// Iterate over path segments to find the language code
	var languageCode string
	for i, segment := range pathSegments {
		if segment == "readership" && i+1 < len(pathSegments) {
			// The language code is the next segment after "librarystats"
			languageCode = pathSegments[i+1]
			break
		}
	}
	return languageCode
}
