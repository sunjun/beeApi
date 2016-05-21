package controllers

import (
	"github.com/astaxie/beego"
	"github.com/sunjun/videoapi/models"
)

type s_response struct {
	Status string            `json:"status"`
	Info   string            `json:"info"`
	Count  int               `json:"count"`
	Lists  []*models.Service `json:"lists"`
}

type ServiceController struct {
	beego.Controller
}

// @Title createRecord
// @Description create record
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /list [get]
func (r *ServiceController) List() {
	list, num, err := models.GetAllServices()
	res := &s_response{}
	if err != nil {
		res.Status = "fail"
		res.Info = err.Error()
	} else {
		res.Status = "Success"
		res.Info = "success"
		res.Count = num
		res.Lists = list
	}
	r.Data["json"] = res
	r.ServeJSON()
}
