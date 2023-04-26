package api

import (
	"encoding/json"
	"fmt"
	"go_blog/model"
	"go_blog/utils"
	"io/ioutil"
	"net/http"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.HandleHttpError("only post", w)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.HandleHttpError(err.Error(), w)
		return
	}

	user := &model.User{}
	err = json.Unmarshal(data, user)
	if err != nil {
		utils.HandleHttpError(err.Error(), w)
		return
	}

	id := model.SaveUser(*user)
	type Response struct {
		Id int64
	}
	res1 := Response{
		Id: id,
	}
	aaa, _ := json.Marshal(res1)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(aaa)

}

type page struct {
	limit int `json:"limit"`
	start int `json:"start"`
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.HandleHttpError("only post", w)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.HandleHttpError(err.Error(), w)
		return
	}

	page := &page{}
	err = json.Unmarshal(data, page)

	fmt.Println(page, "datatatatat")

}
