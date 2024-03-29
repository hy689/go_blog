package utils

import (
	"encoding/json"
	"net/http"
)

type Success struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
}

// type
func HandleSuccess(data interface{}, w http.ResponseWriter) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	success := &Success{
		Code:    200,
		Data:    data,
		Success: true,
	}
	str, _ := json.Marshal(success)
	w.Write(str)

}
