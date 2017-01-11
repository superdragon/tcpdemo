这是使用GoLang编写的一个简单TCP长连接服务实例。

使用 `go run tcpdemoser.go` 运行TCP服务。默认端口是**2134**。

使用命令 `lsof -i tcp:2134` 查看服务是否正常运行。

使用命令 `telnet 127.0.0.1 2134` 默认客户端,连接TCP服务。

```

$ telnet 127.0.0.1 2134
Trying 127.0.0.1...
Connected to localhost.
Escape character is '^]'.

```

* 输入 `time`,服务会返回当前时间。
* 输入 `exit`,服务会断开当前客户端连接。
* 输入其他字符串,服务会返回输入的字符串内容。

