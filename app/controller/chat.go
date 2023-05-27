package controller

import (
	"sync"

	"golang.org/x/net/websocket"
	"gopkg.in/fatih/set.v0"
)

// 本核心在于形成 userid 和 Node的映射关系
type Node struct {
	Conn *websocket.Conn
	//	并行转窜行
	DataQueue chan []byte
	GroupSets set.Interface
}

// userid 和 Node 映射关系表
var clientMap map[int64]*Node = make(map[int64]*Node, 0)

// 读写锁
var rwlocker sync.RWMutex

func AddGroupId(userId, gid int64) {
	//	取得 node
	rwlocker.Lock()
	node, ok := clientMap[userId]
	if ok {
		node.GroupSets.Add(gid)
	}
	rwlocker.Unlock()
}
