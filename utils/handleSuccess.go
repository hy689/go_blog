package utils

import (
	"encoding/json"
	"net/http"
)

type Success struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

// type
func HandleSuccess(data any, w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	success := &Success{
		Code: 200,
		Data: data,
	}
	str, _ := json.Marshal(success)
	w.Write(str)

}
