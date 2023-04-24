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

	// // 根据请求body创建一个json解析器实例
	// decoder := json.NewDecoder(r.Body)

	// // 用于存放参数key=value数据
	// var params map[string]string

	// // 解析参数 存入map
	// decoder.Decode(&params)

	// fmt.Printf("POST json: username=%s, password=%s, password=%s\n", params["username"], params["password"], params["role"])

	// var user model.User
	// user.Username = params["username"]
	// user.Password = params["password"]
	// role, err := strconv.Atoi(params["role"])
	// if err != nil {
	// 	utils.HandleHttpError("role 参数请求错误", w)
	// 	return
	// }
	// user.Role = role
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
