# Clipboard-Share
这是一个跨平台的剪贴板共享工具，支持Windows、Linux、MacOS等操作系统，整个项目使用Go编写，分为服务端和客户端两个部分。
你可以在任意一个平台上（内网/公网）运行一个服务器，然后在多个平台上运行客户端，客户端会默默的同步剪贴板内容到服务器上，服务器会将剪贴板内容同步到所有连接的客户端上。

## 注意
- 这目前只是一个demo，基本功能和运行已经能够保障，但是还会有很多bug和不完善的地方，我会慢慢更新这个项目，直到完善。
- 这是我的第一个go项目，此前我只有简单浏览过go代码，但没有实际上手写过，感谢AI的帮助，帮我解决了很多问题。

## 原理
- 通过Http协议进行通信，客户端不断get请求，同时如果有剪贴板更新，会post请求到服务器
- 其他客户端会通过长轮询的方式获取服务器的剪贴板内容
- 服务器会将剪贴板内容存储在内存中，所有连接的客户端都会获取到最新的剪贴板内容
- server实现：Go的Gin框架
- client实现：Go的clipboard包、resty包

## 使用方法
### 编译服务端
```bash
git clone https://github.com/xucoanthony/clipboard-share.git
cd clipboard-share/server
go build .
```

### 编译客户端
```bash
# 如果你当前在server目录下，请cd到上级目录
cd clipboard-share/client
go build .
```

### 运行服务端
```bash
# 运行服务端（Mac/Linux）
./server
# 运行服务端（Windows）
双击打开server.exe
# 当前服务器没有配置选项，默认运行在8080端口，如果需要修改端口，请在代码中修改，再重新编译
```

### 运行客户端
- 参考运行服务器，需要保持client的二进制文件与config.json在同一目录下，这会在后续的版本中修改路径问题
- 请注意，目前的版本不支持自定义配置文件名和路径，config.json必须与client在同一目录下

## 客户端配置文件
```json
{
  "device_name": "anthony's macbook pro", //设备名称
  "update_frequency_second": 1, //向服务器请求的频率，单位为秒（同时也是客户端读取你剪贴板观察变化的频率）
  "remote_server": {
    "host": "127.0.0.1", //服务器地址，可以是内网/公网ip，可以是域名，可以使用localhost，但请不要加上http://
    "port": 8080 //服务器端口，虽然这里支持配置，但是你需要在代码中修改server的端口，后续版本会升级
  }
}
```

## 未来的计划
### 一般计划
- [ ] 支持自定义配置文件名和路径
- [ ] 支持自定义服务器端口
- [ ] 支持自定义服务器地址
- [ ] 加入https支持（目前的版本是http，你可以使用nginx等反向代理工具来实现https）
### 更宏伟的计划
- [ ] 通过快捷指令或其他方式实现iOS和Android的剪贴板共享
- [ ] 加入类似于AirDrop的非原生实现，支持文件传输
- [ ] English Documentation Support
- [ ] ...