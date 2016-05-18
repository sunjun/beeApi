package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/sunjun/videoapi/models"
)

const (
	LOG_IN = iota
	LOG_OUT
	START_SERVICE
	STOP_SERVICE
)

type Response struct {
	Status string `json:"status"`
	Info   string `json:"info"`
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
	res := &Response{"fail", "error"}
	r.Data["json"] = res
	r.ServeJSON()
	return
}

// @Title login
// @Description Logs user into the system
// @Param	body		body 	models.Record	true		"body for user login content"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [post]
func (r *RecordController) Login() {

	var record models.Record
	err := json.Unmarshal(r.Ctx.Input.RequestBody, &record)

	fmt.Printf("%+v\n", record)
	if err == nil {
		record.RecordType = LOG_IN
		uid, err := models.AddRecord(&record)
		if err == nil {
			res := &Response{"success", "1"}
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
// @router /logout [get]
func (u *RecordController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}
