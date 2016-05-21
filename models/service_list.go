package models

import "github.com/astaxie/beego/orm"

type Service struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Link string `json:"link"`
}

func init() {
	orm.RegisterModel(new(Service))
}

func GetAllServices() ([]*Service, int, error) {
	var services []*Service
	o := orm.NewOrm()
	num, err := o.QueryTable("service").All(&services)

	return services, int(num), err
}
