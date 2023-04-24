package model

import (
	"fmt"
	"time"
)

type Category struct {
	Name string ` json:"name"`
}

func GetCategories() ([]Category, error) {
	var categories []Category
	err := Db.Select(&categories, "select * from category")
	if err != nil {
		return nil, err
	}
	return categories, nil

}

func SaveCategory(category Category) (int64, error) {

	result, err := Db.Exec("insert into category(created_at,name) values(?,?)", time.Now(), category.Name)
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
