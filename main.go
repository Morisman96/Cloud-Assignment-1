package main

import (
	"Assignment1/api"
	_ "Assignment1/server"
	"Assignment1/utils"
	_ "Assignment1/utils"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/librarystats", api.GetGutendex)
	log.Fatal(http.ListenAndServe(":"+utils.PORT, nil))
}
