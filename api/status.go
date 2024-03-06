package api

import (
	"Assignment1/structs"
	"Assignment1/utils"
	"io"
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
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	statusCodeStr := strconv.Itoa(resp.StatusCode)
	return statusCodeStr
}

func HandlerStatus(rw http.ResponseWriter, r *http.Request) {
	utils.PostResponse(rw, StructStatus(), "Status")
}
