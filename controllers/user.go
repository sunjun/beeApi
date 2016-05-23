package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/sunjun/videoapi/models"

	"github.com/astaxie/beego"
)

//response json
type response struct {
	Status string        `json:"status"`
	Info   string        `json:"info"`
	Url    string        `json:"url"`
	List   []*clientLine `json:"list"`
}

type l_response struct {
	Status string        `json:"status"`
	Info   string        `json:"info"`
	List   []*clientLine `json:"list"`
}

type callerCalleeId struct {
	CallerId string
	CalleeId int
}

// Operations about Users
type UserController struct {
	beego.Controller
}

type serverQueue struct {
	queueType int
	name      string
	free      int
	busy      int
	lines     []*serverLine
}

type serverLine struct {
	isBusy   bool
	lineURL  string
	callerID string
	calleeID int
}

type clientQueue struct {
	queueType int
	name      string
	free      int
	wait      int
	lines     []*clientLine
}

type clientLine struct {
	IsWait     bool   `json:"is_wait"`
	LineURL    string `json:"line_url"`
	CallerID   string `json:"caller_id"`
	CalleeID   int    `json:"callee_id"`
	CreateTime int64  `json:"create_time"`
}

var serverQueues []*serverQueue
var clientQueues []*clientQueue

func InitServerQueue() {
	serverQueues = make([]*serverQueue, 100)
	services, _, err := models.GetAllServices()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, s := range services {
		sQueue := &serverQueue{queueType: s.Id, name: s.Name}
		serverQueues = append(serverQueues, sQueue)
	}

	for v := range serverQueues {
		if serverQueues[v] != nil {
			fmt.Println(serverQueues[v])
		}
	}
}

func InitClientQueue() {
	clientQueues = make([]*clientQueue, 100)
	services, _, err := models.GetAllServices()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, s := range services {
		cQueue := &clientQueue{queueType: s.Id, name: s.Name}
		clientQueues = append(clientQueues, cQueue)
	}

	for v := range clientQueues {
		if clientQueues[v] != nil {
			fmt.Println(clientQueues[v])
		}
	}
}

func calleeCreateLine(id int, serviceType int) {
	for v := range serverQueues {
		if sQueue := serverQueues[v]; sQueue != nil {
			if sQueue.queueType == serviceType {
				sQueue.free++
				sLine := &serverLine{isBusy: false, lineURL: "/chat.html", calleeID: id}
				sQueue.lines = append(sQueue.lines, sLine)
				fmt.Printf("%+v\n", sQueue.lines)
				break
			}
		}
	}
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

	if err != nil || ret <= 0 {
		res.Status = "Fail"
		res.Info = err.Error()
	} else {
		res.Status = "Success"
		res.Info = "success"
		calleeCreateLine(callee.Id, ret)
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

// @Title caller create line
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /caller_create_line [post]
func (u *UserController) CallerCreateLine() {
	var callerLine models.CallerLine
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &callerLine)
	var waitNumber int

	fmt.Printf("%+v\n", callerLine)
	if err == nil {
		for v := range clientQueues {
			if cQueue := clientQueues[v]; cQueue != nil {
				if cQueue.queueType == callerLine.LineId {
					waitNumber = cQueue.wait
					cQueue.wait++
					cLine := &clientLine{IsWait: true, CallerID: callerLine.IdNumber, CreateTime: time.Now().UnixNano()}
					cQueue.lines = append(cQueue.lines, cLine)
					fmt.Printf("%+v\n", cLine)
					fmt.Printf("%+v\n", cQueue.lines)
					break
				}
			}
		}
		var info string
		if waitNumber > 0 {
			info = fmt.Sprintf("前面有%d用户在等待", waitNumber)
		} else {
			info = "请等待客服人员接通"
		}
		res := &response{Status: "success", Info: info}
		u.Data["json"] = res
		u.ServeJSON()
		return
	}
	u_return_error(u, err)
}

func getClientLine(callerLine *models.CallerLine, cQueue *clientQueue) (cLine *clientLine) {
	cLines := cQueue.lines
	for i := range cLines {
		if cLine := cLines[i]; cLine != nil {
			if strings.Compare(cLine.CallerID, callerLine.IdNumber) == 0 {
				return cLine
			}
		}
	}

	return nil
}

// @Title caller get line status
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /caller_get_line_status [post]
func (u *UserController) CallerGetLineStatus() {
	var callerLine models.CallerLine
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &callerLine)
	var cLine *clientLine
	var waitNumber int
	res := &response{}
	if err == nil {
		for v := range clientQueues {
			if cQueue := clientQueues[v]; cQueue != nil {
				if cQueue.queueType == callerLine.LineId {
					waitNumber = cQueue.wait
					cLine = getClientLine(&callerLine, cQueue)
					break
				}
			}
		}
		if cLine != nil {
			var info string
			if cLine.IsWait {
				if waitNumber > 0 {
					info = fmt.Sprintf("前面有%d用户在等待", waitNumber)
				} else {
					info = "请等待客服人员接通"
				}
			} else {
				info = "正在接通，请稍后"
				res.Url = cLine.LineURL
			}
			res.Status = "success"
			res.Info = info
			u.Data["json"] = res
			u.ServeJSON()
			return
		}
	}
	u_return_error(u, err)
}

func getClientWaitLines(cQueue *clientQueue) []*clientLine {
	aLines := cQueue.lines
	var waitLines []*clientLine
	for v := range aLines {
		if cLine := aLines[v]; cLine != nil && cLine.IsWait {
			waitLines = append(waitLines, cLine)
		}
	}

	return waitLines
}

func getCallerQueue(serviceType int) *clientQueue {
	for v := range clientQueues {
		if cQueue := clientQueues[v]; cQueue != nil {
			if cQueue.queueType == serviceType {
				return cQueue
			}
		}
	}
	return nil
}

func getCalleeQueue(serviceType int) *serverQueue {
	for v := range serverQueues {
		if sQueue := serverQueues[v]; sQueue != nil {
			if sQueue.queueType == serviceType {
				return sQueue
			}
		}
	}
	return nil
}

func getCallerLine(callerId string, serviceType int) *clientLine {
	for v := range clientQueues {
		if cQueue := clientQueues[v]; cQueue != nil {
			if cQueue.queueType == serviceType {
				cLines := cQueue.lines
				for i := range cLines {
					if cLine := cLines[i]; cLine != nil {
						if strings.Compare(callerId, cLine.CallerID) == 0 {
							return cLine
						}
					}
				}
			}
		}
	}
	return nil
}

func getCalleeLine(calleeId int, serviceType int) *serverLine {
	fmt.Printf("%d %d\n", calleeId, serviceType)
	for v := range serverQueues {
		if sQueue := serverQueues[v]; sQueue != nil {
			fmt.Printf("%+v \n", sQueue)
			if sQueue.queueType == serviceType {
				fmt.Printf("%+v \n", sQueue)
				sLines := sQueue.lines
				for i := range sLines {
					if sLine := sLines[i]; sLine != nil {
						fmt.Printf("%+v %d\n", sLine, calleeId)
						if sLine.calleeID == calleeId {
							return sLine
						}
					}
				}
			}
		}
	}
	return nil
}

func getCalleeServiceType(calleeId int) int {
	return 1
}

// @Title callee get user call list
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /callee_get_user_call_list [post]
func (u *UserController) CalleeGetUserCallList() {
	var calleeUser models.CalleeUser
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &calleeUser)
	fmt.Printf("%+v\n", calleeUser)
	res := &l_response{}
	var waitLines []*clientLine
	if err == nil {
		serviceType := getCalleeServiceType(calleeUser.Id)
		for v := range clientQueues {
			if cQueue := clientQueues[v]; cQueue != nil {
				if cQueue.queueType == serviceType {
					waitLines = getClientWaitLines(cQueue)
					break
				}
			}
		}
		res.Status = "success"
		res.Info = "success"
		fmt.Printf("%+v\n", waitLines)
		res.List = waitLines

		for v := range res.List {
			if c := res.List[v]; c != nil {
				fmt.Printf("%+v\n", c)
			}
		}
		u.Data["json"] = res
		u.ServeJSON()
		return
	}
	u_return_error(u, err)
}

// @Title callee connect caller
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /callee_connect_caller [post]
func (u *UserController) CalleeConnectCaller() {
	var id callerCalleeId
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &id)
	fmt.Printf("%+v\n", id)
	res := &response{}
	if err == nil {
		serviceType := getCalleeServiceType(id.CalleeId)
		sLine := getCalleeLine(id.CalleeId, serviceType)
		cLine := getCallerLine(id.CallerId, serviceType)
		sLine.callerID = id.CallerId
		cLine.CalleeID = id.CalleeId
		sLine.lineURL = "/chat1.html"
		cLine.LineURL = "/chat1.html"

		cLine.IsWait = false
		sLine.isBusy = true

		cQueue := getCallerQueue(serviceType)
		sQueue := getCalleeQueue(serviceType)

		fmt.Printf("%+v\n", cLine)
		fmt.Printf("%+v\n", sLine)
		cQueue.wait--
		sQueue.busy++
		res.Status = "success"
		res.Info = "success"
		res.Url = sLine.lineURL
		u.Data["json"] = res
		u.ServeJSON()
		return
	}
	u_return_error(u, err)
}
