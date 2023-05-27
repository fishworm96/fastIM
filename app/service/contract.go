package service

import (
	"errors"
	"fastIM/app/model"
	"time"
)

type ContactService struct{}

// 添加朋友
func (c *ContactService) AddFriend(userid int64, dstid int64) error {
	if dstid == userid {
		return errors.New("不能添加自己为好友")
	}

	//	判断是否已经添加了好友
	friend := model.Contact{}
	model.DbEngine.Where("ownerid = ?", userid).And("dstobj = ?", dstid).And("cate = ?", model.ConcatCateUser).Get(&friend)

	//	如果好友已经存在，则不添加
	if friend.Id > 0 {
		return errors.New("该好友已经添加过了")
	}

	//	开启事务
	session := model.DbEngine.NewSession()
	session.Begin()

	//	插入两条好友关系数据
	_, s1 := session.InsertOne(model.Contact{
		Ownerid:  userid,
		Dstobj:   dstid,
		Cate:     model.ConcatCateUser,
		Createat: time.Now(),
	})
	_, s2 := session.InsertOne(model.Contact{
		Ownerid:  dstid,
		Dstobj:   userid,
		Cate:     model.ConcatCateUser,
		Createat: time.Now(),
	})
	if s1 == nil && s2 == nil {
		session.Commit()
		return nil
	} else {
		session.Rollback()
		if s1 != nil {
			return s1
		}
		return s2
	}
}

// 搜索社群
func (c *ContactService) SearchCommunity(userId int64) []model.Community {
	conconts := make([]model.Contact, 0)
	comIds := make([]int64, 0)

	model.DbEngine.Where("ownerid = ? and cate = ?", userId, model.ConcatCateComunity).Find(&conconts)

	for _, v := range conconts {
		comIds = append(comIds, v.Dstobj)
	}
	coms := make([]model.Community, 0)
	if len(comIds) == 0 {
		return coms
	}
	model.DbEngine.In("id", comIds).Find(&coms)
	return coms
}

// 根据姓名搜索用户
func (c *ContactService) SearchFriendByName(mobile string) model.User {
	user := model.User{}
	model.DbEngine.Where("mobile = ?", mobile).Get(&user)
	return user
}