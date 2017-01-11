package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

//处理TCP连接
func handleClient(conn net.Conn) {
	//设置两分钟超时。
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute))
	//设置用户输入的命令
	user_cmd := make([]byte, 128)
	defer conn.Close()

	for {
		read_len, err := conn.Read(user_cmd)

		if err != nil {
			fmt.Println(err)
			break
		}
		if read_len == 0 {
			break
		}

		cmd_str := ""
		for i := 0; i < len(user_cmd); i++ {
			if user_cmd[i] == 0 {
				break
			} else {
				cmd_str += string(user_cmd[i])
			}
		}

		//替换字符,例如输入time回车换行后,TCP服务收到的字符串是'time/r/n'
		cmd_str = strings.Replace(cmd_str, "\r\n", "", -1)

		fmt.Printf("read_len = %d user_cmd = %s\n", read_len, user_cmd)

		if cmd_str == "time" { //显示当前时间命令
			time_now := time.Now().String()
			conn.Write([]byte(time_now))
		} else if cmd_str == "exit" { //退出命令
			conn.Close()
		} else if cmd_str != "" {
			conn.Write([]byte(cmd_str))
		}

		// 清空该次读取的内容
		user_cmd = make([]byte, 128)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

//使用【telnet 127.0.0.1 2134】连接TCP服务器
func main() {
	listener, err := net.Listen("tcp", ":2134")
	checkError(err)

	//循环接受TCP的连接请求,通过协程处理连接成功的TCP消息
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		//每接受一个连接,就开启一个协程处理,所以Go很容易实现高并发的TCP服务器
		go handleClient(conn) //创建一个goroutinue处理
	}
}
