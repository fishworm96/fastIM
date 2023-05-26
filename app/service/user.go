package service

import (
	"errors"
	"fastIM/app/model"
	"fastIM/app/util"
)

type UserService struct{}

//用户登录
func (s *UserService) Login(mobile, plainpwd string) (user model.User, err error) {
	//数据库操作
	loginUser := model.User{}
	model.DbEngine.Where("mobile = ?", mobile).Get(&loginUser)
	if loginUser.Id == 0 {
		return loginUser, errors.New("用户不存在")
	}
	//判断密码是否正确
	if !util.ValidatePasswd(plainpwd, loginUser.Salt, loginUser.Passwd) {
		return loginUser, errors.New("密码不正确")
	}
	//刷新用户登录的token值
	token := util.GenRandomStr(32)
	loginUser.Token = token
	model.DbEngine.ID(loginUser.Id).Cols("token").Update(&loginUser)

	//返回新用户信息
	return loginUser, nil
}
