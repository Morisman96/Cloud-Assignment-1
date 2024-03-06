package api

import (
	"Assignment1/structs"
	"Assignment1/utils"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

// StructStatus creates a new status struct
func StructStatus() structs.Status {
	return structs.Status{
		GutendexApi:           GetStatus(utils.GutendexAPI),
		Language2countriesApi: GetStatus(utils.Language2CountriesAPI + utils.STATUSLANGUAGE2COUNTRIESQUERY),
		RestCountriesApi:      GetStatus(utils.RESTCountriesAPI + utils.STATUSRESTCOUNTRIESQUERY),
		Version:               utils.VERSION,
		Uptime:                time.Duration(time.Since(utils.Timer).Seconds())}
}

// GetStatus gets the status of an API
func GetStatus(api string) string {
	resp, err := http.Get(api)
	if err != nil {
		http.Error(nil, "Failed to get request from API", http.StatusBadRequest)
	}
	defer resp.Body.Close()
	statusCodeStr := strconv.Itoa(resp.StatusCode)
	return statusCodeStr
}

// PostStatus posts the status of the API's
func PostStatus(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Status", utils.STATUSPATH)
	err := json.NewEncoder(rw).Encode(StructStatus())
	if err != nil {
		http.Error(rw, "Error during JSON encoding (Error: "+err.Error()+")", http.StatusInternalServerError)
	}
}
