package model

import (
	"fmt"
)

type Article struct {
	Category Category
	Id       string `json:"id"`
	Title    string `json:"title"`
	Cid      string `json:"cid"`
	Content  string `json:"content"`
	Img      string `json:"img"`
	Desc     string `json:"desc"`
}

func SaveArticle(article Article) int64 {
	result, _ := Db.Exec("insert into article(cid,title,content,img,desc) values(?,?,?,?,?)", article.Category.Id, article.Title, article.Content, article.Img, article.Desc)
	id, _ := result.LastInsertId()
	return id

}

func GetArticleById(id int64) Article {
	var article Article
	err := Db.Get(&article, "select * from article")
	if err != nil {
		fmt.Println("select article err:", err)
	}
	return article
}

func GetArticleByCid(cid int64) []Article {
	var article []Article
	err := Db.Select(&article, "select * from article where cid=?", cid)
	if err != nil {
		fmt.Println("select article err:", err)
	}
	return article
}

func DeteleArticle(id int64) int64 {
	res, err := Db.Exec("delete from article where id=?", id)
	if err != nil {
		fmt.Println("delete article err:", err)
	}

	row, _ := res.RowsAffected()
	return row

}

func UpdateArticle(article Article) int64 {
	res, err := Db.Exec("update article set cid=?,title=?,content=?,img=?,desc=? where id=?", article.Category.Id, article.Title, article.Content, article.Img, article.Desc, article.Id)
	if err != nil {
		fmt.Println("update article err:", err)
	}
	row, _ := res.RowsAffected()
	return row
}
