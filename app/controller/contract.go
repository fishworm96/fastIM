package controller

import (
	"log"
	"net/http"

	"fastIM/app/args"
	"fastIM/app/model"
	"fastIM/app/service"
	"fastIM/app/util"
)

var concatService service.ContactService

// 添加朋友
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

// 用户加群
func JoinCommunity(writer http.ResponseWriter, request *http.Request) {
	var arg args.AddNewMember
	util.Bind(request, &arg)
	//	查看群是否存在
	com := concatService.SearchCommunityByName(arg.DstName)
	if com.Id == 0 {
		util.RespFail(writer, "您要加入的群不存在")
	} else {
		log.Printf("community id:%d", com.Id)
		err := concatService.JoinCommunity(arg.Userid, com.Id)
		// 刷新用户的群组信息
		AddGroupId(arg.Userid, com.Id)
		if err != nil {
			util.RespFail(writer, err.Error())
		} else {
			util.RespOk(writer, nil, "")
		}
	}
}
