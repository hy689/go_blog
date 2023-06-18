package model

import (
	"fmt"
	"time"
)

type Category struct {
	Id         int    `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	Icon       string `json:"icon" db:"icon"`
	IconColor  string `json:"iconColor" db:"iconColor"`
	UpdateTime int64  `json:"updateTime" db:"updateTime"`
	CreateTime int64  `json:"createTime" db:"createTime"`
}

func (c *Category) Update(name string, icon string, iconColor string) {
	c.Name = name
	c.Icon = icon
	c.IconColor = iconColor
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

func GetCategoryById(id int) *Category {
	var category Category

	Db.Get(&category, "select * from category where id=?", id)

	return &category
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

	result, err := Db.Exec("insert into category(createTime,updateTime,name,icon,iconColor) values(?,?,?,?,?)", time.Now().Unix(), time.Now().Unix(), category.Name, category.Icon, category.IconColor)
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

func UpdateCategory(category Category) (int64, error) {
	res, err := Db.Exec("update category set name=?,updateTime=?,icon=?,iconColor=? where id=?", category.Name, time.Now().Unix(), category.Icon, category.IconColor, category.Id)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return 0, err
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed, ", err)
		return 0, err
	}
	fmt.Println("update succ:", row)
	return row, nil
}

func DeleteCategory(id int) (int64, error) {
	res, err := Db.Exec("delete from category where id=?", id)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return 0, err
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed, ", err)
		return 0, err
	}
	fmt.Println("delete succ:", row)
	return row, nil
}
