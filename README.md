# Proxy
这是一个用Go语言实现的内网穿透 基于tcp的socket编程，类似nginx但是不是本地的端口反向代理，是远程的反向代理

# 1.0.0

### 功能：

实现了内网穿透功能，把功能添加到Gin的路由中，实现了客户端一键创建-关闭接口

### 说明：

> ## proxy-server：/api/server/router.go:

```go
	server.POST("/init", serverHandler.ServerInit)//建立客户端监听连接端口和展示端口
	server.POST("/check", serverHandler.ServerCheckIsExist)//检测是否端口被占用
	server.GET("/check/show_port/:show_port", serverHandler.ServerCheckByShowPort)//监听内网穿透服务是否存在
	server.GET("/select/show_port/:show_port", serverHandler.Server_SelectByShowPort)//得到对应服务的详细信息
```

> ## proxt-client:/api/client/router.go

```go
	client.POST("/server/init", clientHandler.ServerCheckAndInit)//客户端-一键检测创建服务端监听服务
	client.GET("/server/connect/:show_port", clientHandler.ClientConnectServer)//客户端-一键连接服务端
	client.GET("/server/disconnect/:show_port", clientHandler.ClientDisConnect)//客户端-一键关闭内网穿透服务
```

### 环境：

go+mysql

表：proxy数据库下自动创建server表

### 运行：

```go
//进入对应的文件夹运行
go mod tidy
```

```go
//修改服务端配置文件
//修改客户端配置文件
//修改客户端文件默认字段：/api/client/S_check_init.go 修改ip和port 不然只能在本地设置，之后会改到配置文件里面
```

### 问题：

> ### 客户端请求 /server/init 请求的时间太长--10s左右

> ### 如果不使用有客户端请求服务端，可能会有逻辑错误，例如没有在数据库中创建server记录，想建立监听就会进入socket监听的死循环里面
>
> #### 虽然设置了每个listener最长时间3小时，但是还是会有bug的可能，不会在后端解决，就想放到客户端用客户端中多个net.http请求服务端不断解析操作来避免，最后导致都很复杂，我写的太笨了

> ## 逻辑比较乱
>
> #### 服务端：如何实现内外穿透服务建立中使用很多通道来相互调用，感觉自己写的有问题
>
> #### 客户端：如何实现建立客户端连接和断开连接，是将请求结束然后监听的协程没有关掉，这个是我写的时候遇到的bug，后端写的时候有一直显示端口被占用导致不能进行新的监听。在客户端就利用这一点，然后在别的请求里面发出信号 关闭监听 感觉自己写的有问题，用bug写代码么？不知道，我也才大三
>
> #### 其他问题还未发现，可能还有命名啥的不太规范，我都是一个文件写一个请求的逻辑，然后很多的测试直接就写在逻辑里了，可能很难阅读



###  下个版本预计：

> ##  用Grpc来代替net.http请求server端
>
> ###  客户端添加UI界面，把那个10s的请求分析点，让它快点
>
> ### 修改下原来多的问题，规整下代码

#1.0.1

## 修改

1. 将文件的格式变得规范了，然后在很多的写法上使用了接口的形式，虽然我也不知道为啥那样写，我会写java。我看大佬们那样写就那样写了。
2. 添加的用户的登陆，实现了客户端的接口

	## 解决

> ## 对上个版本的请求10s解决了，解决到1秒好像是我之前写的不咋的，后来放到协程里面，然后多监听了下就快了。咋也不太懂，只能说go牛，啥也不懂的小菜鸡也能优化下



> ###后面想了下，客户端应该是对的。我看到一个帖子里面是这样写的：先发起一个请求拿code，用code请求另外的请求（执行业务）,再另外定时器轮询去监测返回结果。后来我就明白了，哦原来大差不差。这个方式好像微信小程序的openid。然后我这个客户端也差不多是这个思路，之前以为自己写的好笨。

## 问题

> ### 没有一个UI界面，操作还是太笨了
>
> ## 原来的问题没解决，新的问题又一堆----反复在那改

# 1.0.2

## 添加

> ## 感谢开源大佬，有UI界面了。地址是：https://wails.io/zh-Hans/docs/gettingstarted/firstproject 大佬牛！！ 然后就慢慢的先实现出来，然后好用起来！！

最近，也有面试了，祝我好运！！

> ###有演示视频：地址是https://www.bilibili.com/video/BV1Yo4y1A7SP/?spm_id_from=333.999.0.0&vd_source=893db7ccf39aa64845186194a3ab6f1a













































> 
