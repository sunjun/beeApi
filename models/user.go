package models

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
	"strings"
)

var (
	UserList map[string]*User
)

type CallerUser struct {
	Id       int
	IdNumber string    `orm:form:"id_number"`
	Name     string    `orm:form:"name"`
	Address  string    `orm:form:"addres"`
	Gender   string    `orm:form:"gender"`
	AddTime  time.Time `orm:form:"add_time"`
}

type CallerLine struct {
	IdNumber string
	LineId   int
}

type CalleeUser struct {
	Id       int
	Password string
	Name     string
	Gender   string
	Service  int
}

type User struct {
	Id       string
	Username string
	Password string
	Profile  Profile
}

type Profile struct {
	Gender  string
	Age     int
	Address string
	Email   string
}

func init() {
	UserList = make(map[string]*User)
	u := User{"user_11111", "astaxie", "11111", Profile{"male", 20, "Singapore", "astaxie@gmail.com"}}
	UserList["user_11111"] = &u
	orm.RegisterModel(new(CallerUser))
	orm.RegisterModel(new(CalleeUser))
}

func AddUser(u User) string {
	u.Id = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	UserList[u.Id] = &u
	return u.Id
}

func GetUser(uid string) (u *User, err error) {
	if u, ok := UserList[uid]; ok {
		return u, nil
	}
	return nil, errors.New("User not exists")
}

func GetAllUsers() map[string]*User {
	return UserList
}

func UpdateUser(uid string, uu *User) (a *User, err error) {
	if u, ok := UserList[uid]; ok {
		if uu.Username != "" {
			u.Username = uu.Username
		}
		if uu.Password != "" {
			u.Password = uu.Password
		}
		if uu.Profile.Age != 0 {
			u.Profile.Age = uu.Profile.Age
		}
		if uu.Profile.Address != "" {
			u.Profile.Address = uu.Profile.Address
		}
		if uu.Profile.Gender != "" {
			u.Profile.Gender = uu.Profile.Gender
		}
		if uu.Profile.Email != "" {
			u.Profile.Email = uu.Profile.Email
		}
		return u, nil
	}
	return nil, errors.New("User Not Exist")
}

func Login(username, password string) bool {
	for _, u := range UserList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

func DeleteUser(uid string) {
	delete(UserList, uid)
}

func CalleeServiceType(calleeId int) int {
	o := orm.NewOrm()

	var user CalleeUser
	err := o.QueryTable("callee_user").Filter("id", calleeId).One(&user)

	if err == orm.ErrNoRows {
		// No result
		return 1
	}

	return int(user.Service)
}

func CalleeLogin(id int, password string) (int, error) {
	o := orm.NewOrm()

	var user CalleeUser
	err := o.QueryTable("callee_user").Filter("id", id).One(&user)

	fmt.Printf("%d\n", id)
	if err == orm.ErrNoRows {
		// No result
		return -1, errors.New("用户不存在")
	}

	if strings.Compare(password, user.Password) != 0 {
		return -1, errors.New("用户名密码不正确")
	}

	return int(user.Service), nil
}

func AddCallerUser(c *CallerUser) (int, error) {
	o := orm.NewOrm()
	caller := new(CallerUser)
	caller.IdNumber = c.IdNumber
	caller.Name = c.Name
	caller.Address = c.Address
	caller.Gender = c.Gender
	caller.AddTime = time.Now()

	fmt.Printf("%+v\n", caller)

	id, err := o.Insert(caller)
	if err != nil {
		fmt.Println(err)
	}
	return int(id), err
}
