package api

import (
	"go_blog/model"
	"go_blog/utils"
	"net/http"
)

type ChangeCategoryNameCommand struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Icon      string `json:"icon"`
	IconColor string `json:"iconColor"`
}

type AddCategoryCommand struct {
	Name      string `json:"name"`
	Icon      string `json:"icon"`
	IconColor string `json:"iconColor"`
}

type GetCategoriesCommand struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Icon      string `json:"icon"`
	IconColor string `json:"iconColor"`
}

type DeleteCategoryCommand struct {
	ID int `json:"id"`
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

	v := model.GetCategoryById(c.ID)

	if v.Id == 0 {
		utils.HandleError(500, "分类不存在", w)
		return
	}

	v.Update(c.Name, c.Icon, c.IconColor)

	row, err := model.UpdateCategory(*v)
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

	var getCategoriesCommand []GetCategoriesCommand
	for _, v := range categories {
		getCategoriesCommand = append(getCategoriesCommand, GetCategoriesCommand{
			ID:        int(v.Id),
			Name:      v.Name,
			Icon:      v.Icon,
			IconColor: v.IconColor,
		})
	}

	if err != nil {
		utils.HandleError(500, "获取分类失败", w)
		return
	}

	utils.HandleSuccess(getCategoriesCommand, w)

}

func AddCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.HandleError(400, "请求方式错误", w)
		return
	}

	c := &AddCategoryCommand{}
	err := utils.MarshalCommand(r, c)
	if err != nil {
		utils.HandleError(500, err.Error(), w)
		return
	}

	if c.Name == "" {
		utils.HandleError(500, "分类名称不能为空", w)
		return
	}

	v, _ := model.GetCategoryByName(c.Name)
	if v.Name != "" {
		utils.HandleError(500, "分类名称重复", w)
		return
	}

	v = model.Category{}
	v.Update(c.Name, c.Icon, c.IconColor)

	id, err := model.SaveCategory(v)
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

	c := &DeleteCategoryCommand{}
	err := utils.MarshalCommand(r, c)
	if err != nil {
		utils.HandleError(500, err.Error(), w)
		return
	}

	v := model.GetCategoryById(c.ID)
	if v.Id == 0 {
		utils.HandleError(500, "分类不存在", w)
		return
	}

	row, _ := model.DeleteCategory(v.Id)
	if row > 0 {
		utils.HandleSuccess("ok", w)
		return
	}

}
