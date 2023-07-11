package model

import (
	"fmt"
	"time"
)

type Article struct {
	Category    Category `json:"category"`
	Id          int      `json:"id"`
	Title       string   `json:"title"`
	Content     string   `json:"content"`
	Img         string   `json:"img"`
	Description string   `json:"description" db:"description"`
	Cid         int      `json:"cid"`
	Created     int64    `json:"created" db:"created"`
	Updated     int64    `json:"updated" db:"updated"`
}

func (article *Article) Update(category Category, title string, content string, img string, description string) {
	article.Category = category
	article.Title = title
	article.Content = content
	article.Img = img
	article.Description = description
}

func (article *Article) SetCategory(category Category) {
	article.Category = category
}

func SaveArticle(article Article) (int64, error) {
	result, err := Db.Exec("insert into article(cid,title,content,img,description,created,updated) values(?,?,?,?,?,?,?)", article.Category.Id, article.Title, article.Content, article.Img, article.Description, time.Now().Unix(), time.Now().Unix())
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	return id, err

}

func GetArticles(page int, pageSize int, cId int) ([]Article, error) {
	var article []Article
	var err error

	if cId != 0 {
		err = Db.Select(&article, "select * from article where cid=? limit ? offset ?", cId, pageSize, (page-1)*pageSize)
	} else {
		err = Db.Select(&article, "select * from article limit ? offset ?", pageSize, (page-1)*pageSize)
	}

	if err != nil {
		fmt.Println("select article err:", err)
		return nil, err
	}

	return article, nil
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

func DeleteArticle(id int64) int64 {
	res, err := Db.Exec("delete from article where id=?", id)
	if err != nil {
		fmt.Println("delete article err:", err)
	}

	row, _ := res.RowsAffected()
	return row

}

func UpdateArticle(article Article) (int64, error) {
	res, err := Db.Exec("update article set cid=?,title=?,content=?,img=?,description=?,updated=? where id=?", article.Category.Id, article.Title, article.Content, article.Img, article.Description, time.Now().Unix(), article.Id)
	if err != nil {
		fmt.Println("update article err:", err)
		return 0, err
	}
	row, _ := res.RowsAffected()
	return row, nil
}

func SearchArticle(title string, page int, pageSize int) ([]Article, error) {
	var articles []Article
	err := Db.Select(&articles, "select * from article where title like ? limit ? offset ?", "%"+title+"%", pageSize, (page-1)*pageSize)
	if err != nil {
		fmt.Println("select article err:", err)
		return nil, err
	}
	return articles, nil
}
