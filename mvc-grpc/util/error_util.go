package util

import (
	"encoding/json"
	"net/http"
)

type AppError struct {
	Message    string `jason:"message"`
	StatusCode int    `jason:"status"`
	Code       string `jason:"code"`
}

func SendAppError(appErr *AppError, w http.ResponseWriter) {
	jsonValue, _ := json.Marshal(appErr)
	w.WriteHeader(appErr.StatusCode)
	w.Write(jsonValue)
}
