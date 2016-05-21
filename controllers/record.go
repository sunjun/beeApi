package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/sunjun/videoapi/models"
)

const (
	CALLER_LOG_IN = iota
	CALLER_LOG_OUT
	CALLEE_LOG_IN
	CALLEE_LOG_OUT
	START_SERVICE
	STOP_SERVICE
)

type List struct {
	Service1 string `json:"service1"`
	Service2 string `json:"service2"`
	Service3 string `json:"service3"`
	Service4 string `json:"service4"`
	Service5 string `json:"service5"`
	Service6 string `json:"service6"`
	Service7 string `json:"service7"`
	Service8 string `json:"service8"`
	Service9 string `json:"service9"`
}
type List1 struct {
	ServiceName string `json:"service_name"`
	ServiceLink string `json:"service_link"`
}

type Response struct {
	Status string  `json:"status"`
	Info   string  `json:"info"`
	Lists  []List1 `json:"lists"`
}

type RecordController struct {
	beego.Controller
}

// @Title createRecord
// @Description create record
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (r *RecordController) Post() {
	var record models.Record
	json.Unmarshal(r.Ctx.Input.RequestBody, &record)
	uid, _ := models.AddRecord(&record)
	r.Data["json"] = map[string]int{"uid": uid}
	r.ServeJSON()
}

func return_error(r *RecordController, err error) {
	res := &Response{Status: "fail", Info: "error"}
	r.Data["json"] = res
	r.ServeJSON()
	return
}

// @Title CalleeLogin
// @Description Logs user into the system
// @Param	body		body 	models.Record	true		"body for user login content"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /caller_login [post]
func (r *RecordController) CalleeLogin() {

	var record models.Record
	err := json.Unmarshal(r.Ctx.Input.RequestBody, &record)

	fmt.Printf("%+v\n", record)
	if err == nil {
		record.RecordType = CALLEE_LOG_IN
		uid, err := models.AddRecord(&record)
		if err == nil {
			res := &Response{Status: "success", Info: strconv.Itoa(uid)}
			r.Data["json"] = res
			fmt.Printf("%d\n", uid)
			r.ServeJSON()
			return
		}
	}
	return_error(r, err)
}

// @Title CallerLogin
// @Description Logs user into the system
// @Param	body		body 	models.Record	true		"body for user login content"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /caller_login [post]
func (r *RecordController) CallerLogin() {

	var record models.Record
	err := json.Unmarshal(r.Ctx.Input.RequestBody, &record)

	fmt.Printf("%+v\n", record)
	if err == nil {
		record.RecordType = CALLER_LOG_IN
		uid, err := models.AddRecord(&record)
		if err == nil {
			res := &Response{Status: "success", Info: strconv.Itoa(uid)}
			r.Data["json"] = res
			fmt.Printf("%d\n", uid)
			r.ServeJSON()
			return
		}
	}
	return_error(r, err)
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /service_list [get]
func (r *RecordController) ServiceList() {
	//	res := &List{"教育", "户籍", "税务", "教育", "户籍", "税务", "教育", "户籍", "税务"}
	res := &Response{"success", "1", []List1{
		{"税务", "/chat1"},
		{"教育", "/chat2"},
		{"税务2", "/chat3"},
		{"税务3", "/chat4"},
	},
	}
	r.Data["json"] = res
	r.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *RecordController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}
