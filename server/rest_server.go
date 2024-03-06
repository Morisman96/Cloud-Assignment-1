package server

import (
	"fmt"
	"net/http"
	"os"
)

// HomePage takes the README.md file and writes it to homepage
func HomePage(w http.ResponseWriter, r *http.Request) {
	markdown, err := os.ReadFile("README.md")
	if err != nil {
		fmt.Println("Error reading README.md:", err)
		return
	}
	_, err = w.Write(markdown)
	if err != nil {
		return
	}
}
