# 使用go语言websocket构建基础聊天室

> 基础版本只有一个聊天室，所有客户端直连，每个端连接上来，通过ip:port 识别其身份

## golang websocket 类库
```
golang.org/x/net/websocket
```

## 代码

客户端：https://github.com/junnyxi/golang-learn/blob/master/lr-wsclient.html

服务端：https://github.com/junnyxi/golang-learn/blob/master/lr-wsserver.go

## 其他说明

1. 所有客户端连接通过一个map保存，当客户端刷新、离开时，从map中移除；
2. 服务端收到客户端发送消息，遍历整个map，广播到所有连接中的客户端
3. 广播消息通过goroutine发送
