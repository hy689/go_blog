package model

import (
	"fmt"
	"go_blog/utils"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     int    `json:"role"`
}

func SaveUser(user User) int64 {

	password := utils.Encryption(user.Password)

	result, err := Db.Exec("insert into users(created_at,username,password,role) values(sysdate(),?,?,?)", user.Username, password, user.Role)
	if err != nil {
		fmt.Println("insert err:", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("insert err:", err)
	}

	fmt.Println("insert succ:", id)
	return id
}

func QueryUsers(id int64) []User {
	var user []User

	err := Db.Get(&user, "select username,role,created_at from users")
	if err != nil {
		fmt.Println("select err:", err)
	}

	return user
}
