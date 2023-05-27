package controller

import (
	"fastIM/app/args"
	"fastIM/app/model"
	"fastIM/app/service"
	"fastIM/app/util"
	"net/http"
)

var concatService service.ContactService

//添加朋友
func AddFriend(writer http.ResponseWriter, request *http.Request) {
	var arg args.AddNewMember
	util.Bind(request, &arg)
	friend := concatService.SearchFriendByName(arg.DstName)
	if friend.Id == 0 {
		util.RespFail(writer, "您要添加的好友不存在")
	} else {
		//	调用service
		err := concatService.AddFriend(arg.Userid, friend.Id)
		if err != nil {
			util.RespFail(writer, err.Error())
		} else {
			util.RespOk(writer, nil, "好友添加成功")
		}
	}
}

// 加载好友列表
func LoadFriend(writer http.ResponseWriter, request *http.Request) {
	var arg model.Community
	//如果这个用的上，那么可以直接
	util.Bind(request, &arg)
	com, err := concatService.CreateCommunity(arg)
	if err != nil {
		util.RespFail(writer, err.Error())
	} else {
		util.RespOk(writer, com, "")
	}
}

// 创建群
func CreateCommunity(writer http.ResponseWriter, request *http.Request) {
	var arg model.Community
	util.Bind(request, &arg)
	com, err := concatService.CreateCommunity(arg)
	if err != nil {
		util.RespFail(writer, err.Error())
	} else {
		util.RespOk(writer, com, "")
	}
}
