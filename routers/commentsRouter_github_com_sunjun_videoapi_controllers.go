package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/sunjun/videoapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/sunjun/videoapi/controllers:ObjectController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/sunjun/videoapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/sunjun/videoapi/controllers:ObjectController"],
		beego.ControllerComments{
			"Get",
			`/:objectId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/sunjun/videoapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/sunjun/videoapi/controllers:ObjectController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/sunjun/videoapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/sunjun/videoapi/controllers:ObjectController"],
		beego.ControllerComments{
			"Put",
			`/:objectId`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/sunjun/videoapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/sunjun/videoapi/controllers:ObjectController"],
		beego.ControllerComments{
			"Delete",
			`/:objectId`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/sunjun/videoapi/controllers:RecordController"] = append(beego.GlobalControllerRouter["github.com/sunjun/videoapi/controllers:RecordController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/sunjun/videoapi/controllers:RecordController"] = append(beego.GlobalControllerRouter["github.com/sunjun/videoapi/controllers:RecordController"],
		beego.ControllerComments{
			"CalleeLogin",
			`/caller_login`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/sunjun/videoapi/controllers:RecordController"] = append(beego.GlobalControllerRouter["github.com/sunjun/videoapi/controllers:RecordController"],
		beego.ControllerComments{
			"CallerLogin",
			`/caller_login`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/sunjun/videoapi/controllers:RecordController"] = append(beego.GlobalControllerRouter["github.com/sunjun/videoapi/controllers:RecordController"],
		beego.ControllerComments{
			"ServiceList",
			`/service_list`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/sunjun/videoapi/controllers:RecordController"] = append(beego.GlobalControllerRouter["github.com/sunjun/videoapi/controllers:RecordController"],
		beego.ControllerComments{
			"Logout",
			`/logout`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/sunjun/videoapi/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/sunjun/videoapi/controllers:UserController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/sunjun/videoapi/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/sunjun/videoapi/controllers:UserController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/sunjun/videoapi/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/sunjun/videoapi/controllers:UserController"],
		beego.ControllerComments{
			"Get",
			`/:uid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/sunjun/videoapi/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/sunjun/videoapi/controllers:UserController"],
		beego.ControllerComments{
			"Put",
			`/:uid`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/sunjun/videoapi/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/sunjun/videoapi/controllers:UserController"],
		beego.ControllerComments{
			"Delete",
			`/:uid`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/sunjun/videoapi/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/sunjun/videoapi/controllers:UserController"],
		beego.ControllerComments{
			"Login",
			`/login`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/sunjun/videoapi/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/sunjun/videoapi/controllers:UserController"],
		beego.ControllerComments{
			"Logout",
			`/logout`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/sunjun/videoapi/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/sunjun/videoapi/controllers:UserController"],
		beego.ControllerComments{
			"CalleeLogin",
			`/callee_login`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/sunjun/videoapi/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/sunjun/videoapi/controllers:UserController"],
		beego.ControllerComments{
			"CallerLogin",
			`/caller_login`,
			[]string{"post"},
			nil})

}
