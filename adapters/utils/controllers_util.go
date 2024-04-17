package adapter_util

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func GetControllerBodyData(httpRequest *http.Request, requiredParams []string) (*map[string]interface{}, error) {
	jsonBody := make(map[string]interface{})
	err := json.NewDecoder(httpRequest.Body).Decode(&jsonBody)
	if err != nil {
		return nil, errors.New("error to parse request body")
	}

	errRequiredParams := validRequiredBodyParams(requiredParams, jsonBody)
	if errRequiredParams != nil {
		return nil, errRequiredParams
	}

	return &jsonBody, nil
}

func validRequiredBodyParams(requiredParams []string, params map[string]interface{}) error {
	listOfParamsNotInformed := []string{}
	for _, param := range requiredParams {
		data := params[param]

		if data == nil {
			listOfParamsNotInformed = append(listOfParamsNotInformed, param)
		}
	}
	if len(listOfParamsNotInformed) > 0 {
		errorText := fmt.Sprintf("required params not informed: %v", listOfParamsNotInformed)
		return errors.New(errorText)
	}
	return nil
}
