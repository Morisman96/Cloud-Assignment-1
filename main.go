package main

import (
	"Assignment1/api"
	"Assignment1/server"
	_ "Assignment1/server"
	"Assignment1/utils"
	_ "Assignment1/utils"
	"log"
	"net/http"
)

func main() {
	utils.InitTimer()
	http.HandleFunc("/", server.HomePage)
	http.HandleFunc(utils.STATUSPATH, api.PostStatus)
	http.HandleFunc(utils.READERSHIPPATH, api.HandlerGetLanguage2countries)
	http.HandleFunc(utils.BOOKCOUNTPATH, api.HandlerGetGutendex)
	log.Fatal(http.ListenAndServe(utils.PORT, nil))
}
