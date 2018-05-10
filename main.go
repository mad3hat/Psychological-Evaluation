package main

import (
	"time"
	_ "hello/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/astaxie/beego/session/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:111111@tcp(127.0.0.1:3306)/xinli?charset=utf8")
	orm.DefaultTimeLoc = time.UTC
}

func main() {
	beego.Run()
}

