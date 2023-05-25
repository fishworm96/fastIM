package main

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var (
	//完成握手操作
	upgrade = websocket.Upgrader{
		//允许跨域(一般来讲,websocket都是独立部署的)
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	var (
		conn *websocket.Conn
		err error
		data []byte
	)
	//服务端对客户端的http请求(升级为websocket协议)进行应答，应答之后，协议升级为websocket，http建立连接时的tcp三次握手将保持。
	if conn, err = upgrade.Upgrade(w, r, nil); err != nil {
		return
	}

	// 启动一个协程，每隔1s向客户端发送一次心跳消息
	go func() {
		var (
			err error
		)
		for {
			if err = conn.WriteMessage(websocket.TextMessage, []byte("heartbeat")); err != nil {
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()

	// 得到websocket的场链接之后，就可以对客户端传递的数据进行操作了
	for {
		// 通过websocket长连接读到的数据可以是text文本数据，也可以是二进制binary
		if _, data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		if err = conn.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}
	ERR:
	// 出错后，关闭socket连接
	conn.Close()
}

func main() {
	http.HandleFunc("/ws", wsHandler)
	http.ListenAndServe("0.0.0.0:7777", nil)
}