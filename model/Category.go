package model

import (
	"fmt"
	"time"
)

type Category struct {
	Id         int    `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	UpateTime  int64  `json:"updateTime" db:"updateTime"`
	CreateTime int64  `json:"createTime" db:"createTime"`
}

func GetCategories() ([]Category, error) {

	var Category []Category = []Category{}
	err := Db.Select(&Category, "select * from category")

	fmt.Println(err)
	if err != nil {
		return nil, err
	}

	return Category, nil
}

func GetCategoryById(id int) (Category, error) {
	var category Category
	err := Db.Get(&category, "select * from category where id=?", id)
	if err != nil {
		return category, err
	}
	return category, nil
}

func GetCategoryByName(name string) (Category, error) {
	var category Category
	err := Db.Get(&category, "select * from category where name=?", name)
	if err != nil {
		return category, err
	}
	return category, nil
}

func SaveCategory(category Category) (int64, error) {

	result, err := Db.Exec("insert into category(createTime,updateTime,name) values(?,?)", time.Now().Unix()*1000, time.Now().Unix()*1000, category.Name)
	if err != nil {
		fmt.Println("insert category err:", err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("读取id失败", err)
		return 0, err
	}

	return id, nil
}
