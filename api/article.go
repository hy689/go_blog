package api

import (
	"go_blog/model"
	"go_blog/utils"
	"net/http"
	"strconv"
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
	Cid   int `json:"cid"`
	Limit int `json:"limit"`
	Start int `json:"start"`
}

type GetArticlesResponse struct {
	Category    model.Category `json:"category"`
	Id          int            `json:"id"`
	Title       string         `json:"title"`
	Img         string         `json:"img"`
	Description string         `json:"description" db:"description"`
	Created     int64          `json:"created" db:"created"`
	Updated     int64          `json:"updated" db:"updated"`
}

type GetArticleResponse struct {
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

type SearchArticleCommand struct {
	Title    string `json:"title"`
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
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

	query := r.URL.Query()
	cId, err := strconv.Atoi(query.Get("cId"))
	if err != nil {
		cId = 0
	}

	page, err := strconv.Atoi(query.Get("page"))
	if err != nil {
		page = 1
	}

	pageSize, err := strconv.Atoi(query.Get("pageSize"))
	if err != nil {
		pageSize = 10
	}

	articles, err := model.GetArticles(page, pageSize, cId)
	if err != nil {
		utils.HandleError(500, "数据库错误 GetArticles "+err.Error(), w)
	}

	// var getArticlesResponse []GetArticlesResponse
	getArticlesResponse := make([]GetArticlesResponse, 0)

	for i := 0; i < len(articles); i++ {
		articlesResponse := &GetArticlesResponse{}
		v := model.GetCategoryById(articles[i].Cid)
		articlesResponse.Category = *v
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

func GetArticleById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.HandleError(400, "请求格式错误", w)
		return
	}

	query := r.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)

	if id == 0 {
		utils.HandleError(500, "id不能为空", w)
		return
	}

	article := model.GetArticleById(id)
	if article.Id == 0 {
		utils.HandleError(500, "文章不存在", w)
		return
	}

	category := model.GetCategoryById(article.Cid)
	if category == nil {
		utils.HandleError(500, "分类不存在", w)
		return
	}

	articlesResponse := &GetArticleResponse{}
	articlesResponse.Id = article.Id
	articlesResponse.Title = article.Title
	articlesResponse.Img = article.Img
	articlesResponse.Description = article.Description
	articlesResponse.Created = article.Created
	articlesResponse.Updated = article.Updated
	articlesResponse.Content = article.Content
	articlesResponse.Category = *category

	utils.HandleSuccess(articlesResponse, w)

}

func SearchArticle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.HandleError(400, "请求格式错误", w)
		return
	}

	c := &SearchArticleCommand{}
	utils.MarshalCommand(r, c)

	title := c.Title
	page := c.Page
	pageSize := c.PageSize

	if page == 0 {
		page = 1
	}

	if pageSize == 0 {
		pageSize = 10
	}

	articles, err := model.SearchArticle(title, page, pageSize)
	if err != nil {
		utils.HandleError(500, err.Error(), w)
		return
	}

	articlesResponse := make([]GetArticlesResponse, 0)

	for _, v := range articles {
		category := model.GetCategoryById(v.Cid)

		if category == nil {
			category = &model.Category{}
			return
		}

		articlesResponse = append(articlesResponse, GetArticlesResponse{
			Id:          v.Id,
			Title:       v.Title,
			Img:         v.Img,
			Description: v.Description,
			Created:     v.Created,
			Updated:     v.Updated,
			Category:    *category,
		})
	}

	utils.HandleSuccess(articlesResponse, w)
}
