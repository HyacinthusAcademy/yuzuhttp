# Yuzuhttp
更好的Golang http请求库，通过链式封装实现更加简洁，方便的http请求

## 快速开始
1. 安装
```shell
go get -u github.com/HyacinthusAcademy/yuzuhttp
```
2. 使用
```go
// 发送GET请求
var BodyStr string
if err := yuzuhttp.Get("https://example.com/").Do().BodyString(&BodyStr); err != nil {
    panic(err)
    return
}
```

## 实现的功能
- [x] 所有[HTTP请求方法](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Methods)封装
- [x] 支持链式设置GET参数
- [x] 支持链式设置请求头
- [x] 支持链式设置Cookie
- [x] 支持链式设置Bytes格式请求体
- [x] 支持链式设置String格式请求体
- [x] 支持链式设置FormData格式请求体
- [x] 支持链式设置FormUrlencoded格式请求体
- [x] 支持链式设置JSON格式请求体
- [x] 支持链式设置XML格式请求体
- [x] 支持使用HTTP代理环境变量
- [x] 封装非200状态码错误处理
- [x] 支持链式解析JSON格式响应体
- [x] 支持链式解析XML格式响应体
- [x] 支持链式保存响应体到文件
- [x] 完善的单元测试

## 使用方法
更多功能使用请阅读[项目文档](/docs/README.md)

## 许可证
本项目基于`Apache License 2.0`协议开源