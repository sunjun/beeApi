package models

import (
	//	"errors"
	"github.com/astaxie/beego/orm"
	//	"strconv"
	"fmt"
	//	"time"
)

type Record struct {
	Id            int
	CallerName    string `orm:form:"caller_name" `
	CallerAddress string `orm:form:"caller_address" `
	CallerGender  string `orm:form:"caller_gender" `
	CallerId      string `orm:form:"caller_id" `
	CalleeId      string `orm:form:"callee_id" `
	RecordType    int    `orm:form:"type" `

	//	StartTime Time
	//	EndTime   Time
	Star int `orm:form:"star" `
}

func init() {
	orm.RegisterModel(new(Record))
}

func AddRecord(r *Record) (int, error) {
	o := orm.NewOrm()
	record := new(Record)
	record.CallerName = r.CallerName
	record.CallerAddress = r.CallerAddress
	record.CallerGender = r.CallerGender
	record.CalleeId = r.CalleeId

	record.RecordType = r.RecordType

	fmt.Printf("%+v\n", record)

	id, err := o.Insert(record)
	if err != nil {
		fmt.Println(err)
	}
	return int(id), err
}
