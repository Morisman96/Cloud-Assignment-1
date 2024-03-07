package utils

import (
	"Assignment1/structs"
	"net"
	"net/url"
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

// GetUrlPath returns the next path after the given path from the URL
func GetUrlPath(url string, path string) string {
	pathSegments := strings.Split(url, "/")
	var languageCode string
	for i, segment := range pathSegments {
		if segment == path && i+1 < len(pathSegments) {
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

// ExtractHostAndPort extracts the host and port from a URL, returns the host and port as a string
func ExtractHostAndPort(urlStr string) (string, error) {
	// Parse the URL
	u, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}
	host, port, err := net.SplitHostPort(u.Host)
	if err != nil {
		return "", err
	}
	hostPort := net.JoinHostPort(host, port)
	return hostPort, nil
}
