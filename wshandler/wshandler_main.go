//===============================websockt======================
//代表了websocket的路由控制器
//=======================================================
package wshandler

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"math/rand"
)

//实例化一个聊天室类,注意由于这个类不论有多少连接都只有一个所以要放到最外面
var runningActiveRoom *ActiveRoom = &ActiveRoom{
	OnlineUsers: make(map[int]*OnlineUser)}

func WsHandler(ws *websocket.Conn) { //控制主函数
	fmt.Println("开始了wsHandler")
	//对象初始化化,有4种方法具体可以看笔记
	ran_num := rand.Intn(100) //生成一个随机数作为放到聊天室里面的标识号
	real_user := &OnlineUser{WsConnection: ws,
		InRoom: runningActiveRoom,
		UserId: ran_num}

	runningActiveRoom.OnlineUsers[ran_num] = real_user
	real_user.run()
	fmt.Println("结束了wsHandler")
}

//一个聊天室类
type ActiveRoom struct {
	OnlineUsers map[int]*OnlineUser
	//Broadcast   chan Message
	//CloseSign   chan bool
}

//====在线用户类==================
type OnlineUser struct { //一个在线用户
	WsConnection *websocket.Conn //websocket的连接
	InRoom       *ActiveRoom     //指代客户在哪个聊天室
	UserId       int             //用户id，就是昵称
	//UserInfo   *User
	//Send chan string //Message 类型的channel
}

//一个在线用户的方法
func (this *OnlineUser) run() {
	for {
		reply := ""
		if e := websocket.Message.Receive(this.WsConnection, &reply); e != nil { //接收信息如果有错误，这里应该是一个阻塞的过程
			fmt.Println("cant recive from websocket!!")
			break
		}
		fmt.Println("Received back from client: " + reply) //打印接收到的信息
		msg := "other:" + reply
		fmt.Println("Sending to client: " + msg)
		//给当前用户所在房间里面的每一个在线用户发信息就是广播

		for num, user := range this.InRoom.OnlineUsers {
			fmt.Printf("用户：")
			fmt.Printf("%d", num)
			if err := websocket.Message.Send(user.WsConnection, msg); err != nil {
				fmt.Println("==can,t send")
				break
			}
			fmt.Println("==发送成功")
		}

		/*
			for num, _ := range this.InRoom.OnlineUsers {
				fmt.Printf("用户：")
				fmt.Printf("%d", num)
			}*/
	}
	this.WsConnection.Close()
	delete(this.InRoom.OnlineUsers, this.UserId) //如果退出在在线列表中删除用户
}

//=========在线用户类结束=======================

//=============================================================
//======================处理websocket的=========================
