package utils

import (
	"Assignment1/book_json"
	"fmt"
)

func CountAuthors(result *book_json.AuthorsList) int {
	occurrence := make(map[book_json.Author]bool)

	// Iterate through the list and mark each element as true in the map
	for _, author := range result.AuthorsList {

		occurrence[author] = true

		// Return the count of unique elements return len(occurrence)

	}
	fmt.Println(len(occurrence))
	return len(occurrence)
}

func FracBooks(result book_json.GutendexResult, totalBooks int) float32 {
	return float32(result.BookCount) / float32(totalBooks)
}
