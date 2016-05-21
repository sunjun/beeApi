package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/sunjun/videoapi/models"

	"github.com/astaxie/beego"
)

//response json
type response struct {
	Status string `json:"status"`
	Info   string `json:"info"`
}

// Operations about Users
type UserController struct {
	beego.Controller
}

func u_return_error(u *UserController, err error) {
	res := &response{Status: "fail", Info: "error"}
	u.Data["json"] = res
	u.ServeJSON()
	return
}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	uid := models.AddUser(user)
	u.Data["json"] = map[string]string{"uid": uid}
	u.ServeJSON()
}

// @Title Get
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {
	users := models.GetAllUsers()
	u.Data["json"] = users
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		user, err := models.GetUser(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {
	uid := u.GetString(":uid")
	if uid != "" {
		var user models.User
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		uu, err := models.UpdateUser(uid, &user)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = uu
		}
	}
	u.ServeJSON()
}

// @Title delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	uid := u.GetString(":uid")
	models.DeleteUser(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

// @Title login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [post]
func (u *UserController) Login() {
	username := u.GetString("username")
	password := u.GetString("password")
	if models.Login(username, password) {
		u.Data["json"] = "login success"
	} else {
		u.Data["json"] = "user not exist"
	}
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}

// @Title callee login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /callee_login [post]
func (u *UserController) CalleeLogin() {
	var callee models.CalleeUser
	json.Unmarshal(u.Ctx.Input.RequestBody, &callee)
	ret, err := models.CalleeLogin(callee.Id, callee.Password)

	res := &response{}

	if err != nil || ret != 0 {
		res.Status = "Fail"
		res.Info = err.Error()
	} else {
		res.Status = "Success"
		res.Info = "success"
	}
	u.Data["json"] = res
	u.ServeJSON()
}

// @Title caller login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /caller_login [post]
func (u *UserController) CallerLogin() {
	var caller models.CallerUser
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &caller)

	fmt.Printf("%+v\n", caller)
	if err == nil {
		uid, err := models.AddCallerUser(&caller)
		if err == nil {
			res := &response{Status: "success", Info: strconv.Itoa(uid)}
			u.Data["json"] = res
			fmt.Printf("%d\n", uid)
			u.ServeJSON()
			return
		}
	}
	u_return_error(u, err)
}
