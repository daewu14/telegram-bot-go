package utils

import (
	"encoding/json"
	"errors"
	"net/http"
)

// HandleResponseAndParseToModel is Method to parse from response http client to model
func HandleResponseAndParseToModel(response *http.Response, data any) error {
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return errors.New("Failed to fetch the data")
	}
	errUm := json.NewDecoder(response.Body).Decode(&data)
	if errUm != nil {
		return errUm
	}

	return nil
}
