package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/sunjun/videoapi/models"
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
	uid, _ := models.AddRecord(record.CallerName, record.CallerId)
	r.Data["json"] = map[string]int{"uid": uid}
	r.ServeJSON()
}
