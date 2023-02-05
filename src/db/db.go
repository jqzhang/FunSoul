package db

import (
	"GoFun/src/config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var Instance *sql.DB

func init() {
	cfg := config.Config

	mysqlArgs := fmt.Sprintf("%s:%s@tcp(%s)/go_fun?charset=utf8&parseTime=True&loc=Local", cfg.DB.Name, cfg.DB.Password, cfg.DB.Host)
	println(mysqlArgs)
	var err error
	Instance, err = sql.Open("mysql", mysqlArgs)
	if err != nil {
		panic("数据库连接失败: " + err.Error())
	}

}
