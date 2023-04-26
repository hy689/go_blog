package api

import (
	"encoding/json"
	"go_blog/model"
	"go_blog/utils"
	"io/ioutil"
	_ "io/ioutil"
	"net/http"
)

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("请求方式不对"))
		return
	}

	category := &model.Category{}
	data, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(data, category)

	if category.Id == 0 {
		utils.HandleError(500, "分类id不能为空", w)
		return
	}

	row, err := model.UpdateCategory(*category)
	if err != nil || row <= 0 {
		utils.HandleError(500, "更新分类失败", w)
		return
	}

	utils.HandleSuccess("ok", w)

}
func GetCategories(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("请求方式不对"))
		return
	}

	categories, err := model.GetCategories()
	if err != nil {
		utils.HandleError(500, "获取分类失败", w)
		return
	}

	utils.HandleSuccess(categories, w)

}

func AddCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("请求方式不对"))
		return
	}

	data, _ := ioutil.ReadAll(r.Body)

	var category = &model.Category{}
	json.Unmarshal(data, category)

	if category.Name == "" {
		utils.HandleError(500, "分类名称不能为空", w)
		return
	}

	category1, _ := model.GetCategoryByName(category.Name)
	if category1.Name != "" {
		utils.HandleError(500, "分类名称重复", w)
		return
	}

	id, err := model.SaveCategory(*category)
	if err != nil {
		utils.HandleError(500, "添加分类失败", w)
		return
	}

	type CategoryResponse struct {
		Id int `json:"id"`
	}
	categoryResponse := &CategoryResponse{
		Id: int(id),
	}

	utils.HandleSuccess(categoryResponse, w)
}
