package api

import (
	"go_blog/model"
	"go_blog/utils"
	"net/http"
)

type AddArticleCommand struct {
	Title       string `json:"title"`
	Content     string `json:"content"`
	Img         string `json:"img"`
	Description string `json:"description"`
	Cid         int    `json:"cid"`
}

type AddArticleResponse struct {
	Id int `json:"id"`
}

type GetArticlesCommand struct {
	Cid int `json:"cid"`
}

type GetArticlesResponse struct {
	Category    model.Category `json:"category"`
	Id          int            `json:"id"`
	Title       string         `json:"title"`
	Content     string         `json:"content"`
	Img         string         `json:"img"`
	Description string         `json:"description" db:"description"`
	Created     int64          `json:"created" db:"created"`
	Updated     int64          `json:"updated" db:"updated"`
}

type DeleteArticleCommand struct {
	Id int `json:"id"`
}

type DeleteArticleResponse struct {
}

type UpdateArticleCommand struct {
	Title       string `json:"title"`
	Content     string `json:"content"`
	Img         string `json:"img"`
	Description string `json:"description"`
	Cid         int    `json:"cid"`
	Id          int    `json:"id"`
}

func AddArticle(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		utils.HandleError(400, "请求方式错误", w)
		return
	}

	c := &AddArticleCommand{}
	err := utils.MarshalCommand(r, c)

	if err != nil {
		utils.HandleError(500, err.Error(), w)
		return
	}

	category := model.GetCategoryById(c.Cid)

	if category == nil {
		utils.HandleError(500, "分类不存在", w)
		return
	}

	if c.Content == "" {
		utils.HandleError(500, "内容不能为空", w)
		return
	}

	if c.Img == "" {
		utils.HandleError(500, "图片地址不能为空", w)
		return
	}

	if c.Description == "" {
		utils.HandleError(500, "文章描述不能为空", w)
		return
	}

	if c.Title == "" {
		utils.HandleError(500, "文章标题不能为空", w)
		return
	}

	article := &model.Article{}
	article.Update(*category, c.Title, c.Content, c.Img, c.Description)

	lastId, err := model.SaveArticle(*article)
	if err != nil {
		utils.HandleError(500, err.Error(), w)
		return
	}
	if lastId == 0 {
		utils.HandleError(500, "添加文章失败", w)
	}

	addArticleResponse := AddArticleResponse{
		Id: int(lastId),
	}

	utils.HandleSuccess(addArticleResponse, w)
}

func GetArticleList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.HandleError(400, "请求格式错误", w)
		return
	}

	articles, err := model.GetArticles()
	if err != nil {
		utils.HandleError(500, "数据库错误 GetArticles "+err.Error(), w)
	}

	var getArticlesResponse []GetArticlesResponse
	// getArticlesResponse := make([]GetArticlesResponse, 0)

	for i := 0; i < len(articles); i++ {
		articlesResponse := &GetArticlesResponse{}
		v := model.GetCategoryById(articles[i].Cid)
		articlesResponse.Category = *v
		articlesResponse.Content = articles[i].Content
		articlesResponse.Id = articles[i].Id
		articlesResponse.Title = articles[i].Title
		articlesResponse.Img = articles[i].Img
		articlesResponse.Description = articles[i].Description
		articlesResponse.Created = articles[i].Created
		articlesResponse.Updated = articles[i].Updated
		getArticlesResponse = append(getArticlesResponse, *articlesResponse)
	}

	utils.HandleSuccess(getArticlesResponse, w)

}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.HandleError(400, "请求格式错误", w)
		return
	}

	c := &DeleteArticleCommand{}
	utils.MarshalCommand(r, c)

	v := model.GetArticleById(int64(c.Id))

	if v.Id == 0 {
		utils.HandleError(500, "文章不存在", w)
		return
	}

	res := model.DeleteArticle(int64(v.Id))
	if res == 0 {
		utils.HandleError(500, "删除失败", w)
		return
	}

	utils.HandleSuccess("ok", w)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.HandleError(400, "请求格式错误", w)
		return
	}

	c := &UpdateArticleCommand{}
	utils.MarshalCommand(r, c)

	if c.Id == 0 {
		utils.HandleError(500, "id不能为空", w)
		return
	}

	article := model.GetArticleById(int64(c.Id))
	if article.Id == 0 {
		utils.HandleError(500, "文章不存在", w)
		return
	}

	category := model.GetCategoryById(c.Cid)
	if category == nil {
		utils.HandleError(500, "分类不存在", w)
		return
	}

	article.Update(*category, c.Title, c.Content, c.Img, c.Description)

	_, err := model.UpdateArticle(article)
	if err != nil {
		utils.HandleError(500, err.Error(), w)
		return
	}

	utils.HandleSuccess("ok", w)

}
