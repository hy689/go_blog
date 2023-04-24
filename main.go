package main

import (
	"go_blog/model"
	"go_blog/routes"
)

func main() {
	model.InitDb()
	routes.InitRouter()

	// result, err := model.Db.Exec("insert into users(username,password,role) values(?,?,?)", "admin", "123456", 0)
	// if err != nil {
	// 	fmt.Println("insert err:", err)
	// }

	// id, err := result.LastInsertId()
	// if err != nil {
	// 	fmt.Println("insert err:", err)
	// }

	// fmt.Println("insert succ:", id)
}
