# 设置Cookie
通过`AddCookie()`方法即可快速添加Cookie
```go
if err := yuzuhttp.Get("https://example.com/").AddCookie("client", "yuzuhttp").AddCookie("Yuzu", "ismywife").Do().Error; err != nil {
    panic(err)
}
```

也可以通过`SetCookies()`方法使用Map批量设置
```go
// 创建一个Map
var Cookies = map[string]string{
    "client": "yuzuhttp",
    "Yuzu": "ismywife",
}

if err := yuzuhttp.Get("https://example.com/").SetCookies(Cookies).Do().Error; err != nil {
    panic(err)
}
```