package main

import (
	"./wshandler"
	"code.google.com/p/go.net/websocket"
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"strings"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//获取本地的ip，这个方法需要联网，局域网的方法看笔记
	conn, err := net.Dial("udp", "baidu.com:80") //对百度进行拨号
	if err != nil {
		fmt.Println("main.go 18行====不能联网得到本机IP", err)
		fmt.Println(err.Error())
		return
	}
	Ip := strings.Split(conn.LocalAddr().String(), ":")[0]
	conn.Close() //用完后要关闭要不会卡这
	//======获取ip完成========

	//对模版字段进行赋值
	fmt.Println("运行了IndexHandler")
	TemDate := make(map[string]string)
	TemDate["ServerIp"] = Ip

	if r.Method == "GET" { //代表第一次打开页面
		t, _ := template.ParseFiles("template/index.html") //使用这下面的模版
		t.Execute(w, TemDate)
	} else { //如果是其他方法
		//请求的是登录数据，那么执行登录的逻辑判断
		//r.ParseForm()
		fmt.Println("method:", r.Method)
		//fmt.Println("password:", r.Form["password"])
	}
	fmt.Println("结束了IndexHandler")
}

//========模版数据=============
func NotFoundHandler(w http.ResponseWriter, r *http.Request) { //如果路由规则不符合没有注册的如/2333,/22ww等
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/index", http.StatusFound) //地址重定向
	}

	t, err := template.ParseFiles("template/404.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, nil)
}

func main() {
	fmt.Println("程序启动。。")

	http.Handle("/css/", http.FileServer(http.Dir("template")))
	//对于静态的文件css，js等在template文件里面找
	http.Handle("/js/", http.FileServer(http.Dir("template")))
	//这个是加载html文件遇到类似与/js/main.js时执行的路由
	http.Handle("/images/", http.FileServer(http.Dir("template")))
	//如果类似遇到了/images/404images/404.png 就会在前面加上template/images...
	//这里的handle当一个连接过来的时候都会多开一个wshandler
	http.Handle("/ws", websocket.Handler(wshandler.WsHandler)) //响应了ws://127.0.0.1/ws的websocket

	http.HandleFunc("/index", IndexHandler)
	//http.HandleFunc("/login", login)
	http.HandleFunc("/", NotFoundHandler) //当没有找到路径名字时
	fmt.Println("正在监听7878端口")
	err1 := http.ListenAndServe(":7878", nil)
	if err1 != nil {
		log.Fatal("ListenAndServe:", err1)
	} else {
		fmt.Println("正在监听7878端口")
	}
}
