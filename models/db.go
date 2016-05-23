package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func InitDB() error {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	//	err := orm.RegisterDataBase("default", "mysql", "root:123456@tcp(104.131.156.105:3306)/admin?charset=utf8&&loc=Asia%2FShanghai")
	err := orm.RegisterDataBase("default", "mysql", "root:123456@/admin?charset=utf8")
	return err
}
