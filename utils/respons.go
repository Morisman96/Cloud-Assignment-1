package utils

import (
	"Assignment1/structs"
	"encoding/json"
	"net/http"
)

// PostResponse posts a struct to the api
func PostResponse(rw http.ResponseWriter, response interface{}, key string) {
	rw.Header().Set(key, "application/json")
	switch response.(type) {
	case []structs.BookLangReaders:
		if response == nil || len(response.([]structs.BookLangReaders)) == 0 {
			http.Error(rw, "No Result for Readership Found", http.StatusBadRequest)
		} else {
			EncodeJSON(rw, response, key)
		}
	case []structs.BookLanguage:
		if response == nil || len(response.([]structs.BookLanguage)) == 0 {
			http.Error(rw, "No Result for Bookcount Found", http.StatusBadRequest)
		} else {
			EncodeJSON(rw, response, key)
		}
	case structs.Status:
		if response == nil {
			http.Error(rw, "Failed to get status, returned nil", http.StatusInternalServerError)
		} else {
			EncodeJSON(rw, response, key)
		}
	default:
		http.Error(rw, "No response found: "+key, http.StatusInternalServerError)
	}
}

// GetResponse gets a JSON response from an api
func GetResponse(rw http.ResponseWriter, query string) interface{} {
	var jsonResp interface{}
	resp, err := http.Get(query)
	if err != nil {
		http.Error(rw, "Failed to get request from API: "+query, http.StatusBadRequest)
	}
	if resp.StatusCode != http.StatusOK {
		http.Error(rw, "Failed to get request from  API: "+query, http.StatusServiceUnavailable)
	} else {
		jsonResp = HandleGetResponse(query, resp)
	}
	return jsonResp
}

// DecodeJSON decodes the JSON from the API
func DecodeJSON(resp *http.Response, jsonStruct interface{}) {
	if resp.StatusCode != http.StatusOK || resp.Body == nil {
		http.Error(nil, "Failed to get request from API: ", http.StatusServiceUnavailable)
	} else {
		err := json.NewDecoder(resp.Body).Decode(jsonStruct)
		if err != nil {
			http.Error(nil, "Failed to decode JSON", http.StatusBadRequest)
		}
	}
}

func EncodeJSON(rw http.ResponseWriter, jsonStruct interface{}, key string) {
	err := json.NewEncoder(rw).Encode(jsonStruct)
	if err != nil {
		http.Error(rw, "Error during JSON encoding: "+key, http.StatusInternalServerError)
	}
}

// HandleGetResponse handles the response JSON from the API and returns a struct.
func HandleGetResponse(api string, resp *http.Response) interface{} {
	var switchAPI, _ = ExtractHostAndPort(api)
	switch switchAPI {
	case GUTENDEXCASE:
		var gutendexResult structs.GutendexResult
		DecodeJSON(resp, &gutendexResult)
		return gutendexResult

	case LANGUAGE2COUNTRIESCASE:
		var language2countriesList []structs.Language2Countries
		DecodeJSON(resp, &language2countriesList)
		return language2countriesList

	case RESTCOUNTRIESCASE:
		var population []structs.RestCountry
		DecodeJSON(resp, &population)
		return population

	default:
		http.Error(nil, "API not found", http.StatusBadRequest)
		return nil
	}

}
