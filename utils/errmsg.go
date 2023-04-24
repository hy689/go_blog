package utils

import (
	"encoding/json"
	"net/http"
)

const (
	SUCCESS = 200
	ERROR   = 500

	// CODE = 100 用户模块错误
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_EXIST      = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
)

var codemsg = map[int]string{
	SUCCESS:                "OK",
	ERROR:                  "FAIL",
	ERROR_USERNAME_USED:    "用户名已存在",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_EXIST:      "TOKEN不存在",
	ERROR_TOKEN_RUNTIME:    "TOKEN已过期",
	ERROR_TOKEN_WRONG:      "TOKEN不正确",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误",
}

func GetErrMag(code int) string {
	return codemsg[code]
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func HandleError(code int, err string, w http.ResponseWriter) {
	var error Error
	error.Code = code
	error.Message = err
	w.WriteHeader(http.StatusOK)
	str, _ := json.Marshal(error)
	w.Write(str)
}

func HandleHttpError(err string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(err))
}
