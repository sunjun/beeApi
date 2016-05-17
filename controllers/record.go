package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/sunjun/videoapi/models"
)

const (
	LOG_IN = iota
	LOG_OUT
	START_SERVICE
	STOP_SERVICE
)

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
	r.Data["status"] = "fail"
	//	r.Data["info"] = err.Error()
	r.Data["info"] = "error"
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
	if err == nil {
		record.RecordType = LOG_IN
		id, err := models.AddRecord(&record)
		if err == nil {
			r.Data["status"] = "success"
			r.Data["info"] = id
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
