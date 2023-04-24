package model

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func InitDb() {
	conn, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1)/ginblog")
	if err != nil {
		fmt.Println("数据库连接失败", err)
	}
	Db = conn
}
