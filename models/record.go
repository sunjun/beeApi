package models

import (
	//	"errors"
	"github.com/astaxie/beego/orm"
	//	"strconv"
	//	"time"
	"fmt"
)

type Record struct {
	Id         int
	CallerName string `orm:form:"caller_name" `
	CallerId   string `orm:form:"caller_id" `
	CalleeId   string `orm:form:"callee_id" `
	//	StartTime  Time
	//	EndTime    Time
	Star int
}

func init() {
	orm.RegisterModel(new(Record))
}

func AddRecord(callerName string, callerID string) (int, error) {
	fmt.Printf("callerName111", "callerID222")
	fmt.Printf(callerName, callerID)
	o := orm.NewOrm()
	record := new(Record)
	record.CallerName = callerName
	record.CallerId = callerID
	record.CalleeId = callerID

	fmt.Printf("%+v\n", record)

	id, err := o.Insert(record)
	if err != nil {
		fmt.Println(err)
	}
	return int(id), err
}
