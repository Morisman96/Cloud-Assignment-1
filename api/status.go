package api

import (
	"Assignment1/book_json"
	"Assignment1/utils"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

func StructStatus() book_json.Status {
	return book_json.Status{GutendexApi: GetStatus(utils.GutendexAPI), Language2countriesApi: GetStatus(utils.Language2CountriesAPI + "no"), RestCountriesApi: GetStatus(utils.RESTCountriesAPI + "norway"), Version: utils.VERSION, Uptime: time.Duration(time.Since(utils.Timer).Seconds())}
}
func GetStatus(api string) string {

	resp, err := http.Get(api)
	if err != nil {
		// Handle connection errors
	}
	defer resp.Body.Close()

	// Convert status code to string
	statusCodeStr := strconv.Itoa(resp.StatusCode)

	// Return status code as string
	return statusCodeStr
}

func PostStatus(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Status", utils.STATUSPATH)
	err := json.NewEncoder(rw).Encode(StructStatus())
	if err != nil {
		http.Error(rw, "Error during JSON encoding (Error: "+err.Error()+")", http.StatusInternalServerError)
	}
}
