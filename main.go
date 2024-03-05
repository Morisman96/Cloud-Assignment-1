package main

import (
	"Assignment1/api"
	_ "Assignment1/server"
	"Assignment1/utils"
	_ "Assignment1/utils"
	"log"
	"net/http"
)

// Variable to store the start time

func main() {
	utils.InitTimer()
	http.HandleFunc(utils.STATUSPATH, api.PostStatus)
	http.HandleFunc(utils.READERSHIPPATH, api.HandlerGetLanguage2countries)
	http.HandleFunc(utils.BOOKCOUNTPATH, api.HandlerGetGutendex)
	log.Fatal(http.ListenAndServe(utils.PORT, nil))
}
