# 设置GET参数
通过`AddURLValue()`方法可以即可快速添加GET参数
```go
if err := yuzuhttp.Get("https://example.com/").AddURLValue("client", "yuzuhttp").AddURLValue("yuzu", "ismywife").Do().Error; err != nil {
    panic(err)
}
```
也可以使用别名方法`AddQuery()`
```go
if err := yuzuhttp.Get("https://example.com/").AddURLValue("client", "yuzuhttp").AddQuery("yuzu", "ismywife").Do().Error; err != nil {
    panic(err)
}
```

也可以通过`SetURLValues()`方法使用Map批量设置
```go
// 创建一个Map
var URLValues = map[string]string{
    "client": "yuzuhttp",
    "yuzu": "ismywife",
}

if err := yuzuhttp.Get("https://example.com/").SetURLValues(URLValues).Do().Error; err != nil {
    panic(err)
}
```
也可以使用别名方法`SetQuerys()`
```go
// 创建一个Map
var URLValues = map[string]string{
    "client": "yuzuhttp",
    "yuzu": "ismywife",
}

if err := yuzuhttp.Get("https://example.com/").SetQuerys(URLValues).Do().Error; err != nil {
    panic(err)
}
```