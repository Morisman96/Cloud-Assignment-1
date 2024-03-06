package utils

import (
	"Assignment1/structs"
	"encoding/json"
	"fmt"
	"net/http"
)

// PostResponse posts a struct to the api
func PostResponse(rw http.ResponseWriter, response interface{}, key string) {
	rw.Header().Set(key, "application/json")
	err := json.NewEncoder(rw).Encode(response)
	if err != nil {
		http.Error(rw, "Error during JSON encoding"+key, http.StatusInternalServerError)
	}
}

// GetResponse gets a response from an api
func GetResponse(rw http.ResponseWriter, query string) interface{} {
	resp, err := http.Get(query)
	if err != nil {
		http.Error(rw, "Failed to get request from"+query, http.StatusBadRequest)
	}
	var jsonResp = HandleGetResponse(query, resp)

	return jsonResp
}

func DecodeJSON(resp *http.Response, jsonStruct interface{}) {
	err := json.NewDecoder(resp.Body).Decode(jsonStruct)
	if err != nil {
		http.Error(nil, "Failed to decode JSON", http.StatusBadRequest)
	}
}

func HandleGetResponse(api string, resp *http.Response) interface{} {
	var switchAPI, _ = ExtractHostAndPort(api)
	switch switchAPI {
	case GUTENDEXCASE:
		var gutendexResult structs.GutendexResult
		DecodeJSON(resp, &gutendexResult)
		return gutendexResult

	case Language2CountriesAPI:
		var language2countriesList []*structs.Language2Countries
		DecodeJSON(resp, &language2countriesList)
		return language2countriesList

	case RESTCountriesAPI:
		var population structs.RestCountry
		DecodeJSON(resp, &population)
		return population

	default:
		return fmt.Errorf("Failed to get request from API: %s", api)
	}
}
