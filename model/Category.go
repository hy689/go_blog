package model

import (
	"database/sql"
	"fmt"
	"time"
)

type Category struct {
	Id         int          `json:"id" db:"id"`
	Name       string       `json:"name" db:"name"`
	UpateTime  sql.NullTime `json:"updateTime" db:"updateTime"`
	CreateTime sql.NullTime `json:"createTime" db:"createTime"`
}

func GetCategories() ([]Category, error) {
	var Category []Category
	err := Db.Select(&Category, "select * from category")

	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	return Category, nil
}

func SaveCategory(category Category) (int64, error) {

	result, err := Db.Exec("insert into category(createTime,name) values(?,?)", time.Now(), category.Name)
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
