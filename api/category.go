package api

import (
	"encoding/json"
	"go_blog/model"
	"go_blog/utils"
	"io/ioutil"
	_ "io/ioutil"
	"net/http"
)

type ChangeCategoryNameCommand struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.HandleError(400, "请求方式错误", w)
		return
	}

	c := &ChangeCategoryNameCommand{}
	err := utils.MarshalCommand(r, c)
	if err != nil {
		utils.HandleError(500, err.Error(), w)
		return
	}

	if c.ID == 0 {
		utils.HandleError(500, "分类id不能为空", w)
		return
	}

	v, err := model.GetCategoryById(c.ID)
	if err != nil {
		utils.HandleError(500, "获取分类失败", w)
		return
	}
	if v.Id == 0 {
		utils.HandleError(500, "分类不存在", w)
		return
	}

	v.Update(c.Name)

	row, err := model.UpdateCategory(v)
	if err != nil || row <= 0 {
		utils.HandleError(500, "更新分类失败", w)
		return
	}

	utils.HandleSuccess("ok", w)

}
func GetCategories(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.HandleError(400, "请求方式错误", w)
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
		utils.HandleError(400, "请求方式错误", w)
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

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.HandleError(400, "请求方式错误", w)
	}

	category := &model.Category{}
	data, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(data, category)

	if category.Id == 0 {
		// utils.HandleError(500,"")
	}

}
