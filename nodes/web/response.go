package web

import (
	"encoding/json"
	"net/http"
)

type BasicResponse struct {
	Success bool `json:"success"`
}

func NewSuccessResponse() BasicResponse {
	return BasicResponse{
		Success: true,
	}
}

type ErrorResponse struct {
	Success bool     `json:"success"`
	Errors  []string `json:"errors"`
}

func NewErrorResponse(errs ...string) ErrorResponse {
	return ErrorResponse{
		Success: false,
		Errors:  errs,
	}
}

func sendJSONResponse(w http.ResponseWriter, status int, res any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	encoder := json.NewEncoder(w)
	encoder.Encode(res)
}

func sendClientErrorResponse(w http.ResponseWriter, status int, errs []error) {
	messages := make([]string, len(errs))
	for index, err := range errs {
		messages[index] = err.Error()
	}

	sendJSONResponse(w, status, NewErrorResponse(messages...))
}

func sendInternalErrorResponse(w http.ResponseWriter) {
	sendJSONResponse(w, http.StatusInternalServerError, NewErrorResponse("an internal error occurred"))
}

type DataResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

func NewSuccessDataResponse(data interface{}) DataResponse {
	return DataResponse{
		Success: true,
		Data:    data,
	}
}
