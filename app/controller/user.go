package controller

import (
	"fastIM/app/service"
	"fastIM/app/util"
	"net/http"
)

var UserService service.UserService

// UserLogin 用户登录
func UserLogin(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()

	mobile := request.PostForm.Get("mobile")
	plainpwd := request.PostForm.Get("passwd")

	//校验参数
	if len(mobile) == 0 || len(plainpwd) == 0 {
		util.RespFail(writer, "用户名或密码不正确")
	}

	loginUser, err := UserService.Login(mobile, plainpwd)
	if err != nil {
		util.RespFail(writer, err.Error())
	} else {
		util.RespOk(writer, loginUser, "")
	}
}
