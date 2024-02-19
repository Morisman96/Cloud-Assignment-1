package utils

import (
	"Assignment1/book_json"
)

func CountAuthors(result book_json.GuntendexResult) int {
	occurrence := make(map[book_json.Author]bool)

	// Iterate through the list and mark each element as true in the map
	for _, author := range result.Books {
		for _, name := range author.Authors {
			occurrence[name] = true
		}

		// Return the count of unique elements return len(occurrence)

	}
	return len(occurrence)
}

func FracBooks(result book_json.GuntendexResult, totalBooks int) float32 {
	return float32(result.BookCount) / float32(totalBooks)
}
