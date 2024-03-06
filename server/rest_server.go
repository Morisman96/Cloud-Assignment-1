package server

import (
	"fmt"
	_ "github.com/russross/blackfriday/v2"
	"net/http"
	"os"
)

// HomePage takes the HomePage.html file and writes it to homepage
func HomePage(w http.ResponseWriter, r *http.Request) {

	content, err := os.ReadFile("HomePage.html")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	_, err = w.Write(content)
	if err != nil {
		http.Error(w, "Error failed to write the home page", http.StatusInternalServerError)
	}
}
