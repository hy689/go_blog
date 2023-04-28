package api

import (
	"go_blog/model"
	"go_blog/utils"
	"net/http"
)

type ArticleRequest struct {
	Title       string `json:"title"`
	Content     string `json:"content"`
	Img         string `json:"img"`
	Description string `json:"description"`
	Cid         int    `json:"cid"`
}

func AddArticle(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		utils.HandleError(400, "请求方式错误", w)
		return
	}

	c := &ArticleRequest{}
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

	type AddArticleResponse struct {
		Id int `json:"id"`
	}

	addArticleResponse := AddArticleResponse{
		Id: int(lastId),
	}

	utils.HandleSuccess(addArticleResponse, w)
}
