/*
 * @Author: nijineko
 * @Date: 2024-08-29 19:40:54
 * @LastEditTime: 2024-08-30 17:25:10
 * @LastEditors: nijineko
 * @Description: 请求处理
 * @FilePath: \yuzuhttp\request\request.go
 */
package request

import (
	"fmt"
	"io"
	"net/http"

	"github.com/HyacinthusAcademy/yuzuhttp/response"
)

// 请求对象
type Request struct {
	Method   string            // 请求方法
	URL      string            // 请求地址
	URLValue map[string]string // GET参数
	Header   map[string]string // 请求头
	Body     io.ReadCloser     // 请求体

	// 配置
	enableHTTPProxy   bool // 启用HTTP代理环境变量支持
	enableNon200Error bool // 启用非200状态码时返回错误

	Error error // 错误
}

/**
 * @description: 初始化请求
 * @param {string} method 请求方法
 * @param {string} url 请求地址
 * @return {*Request} 请求对象
 */
func New(method, url string) *Request {
	return &Request{
		Method:            method,
		URL:               url,
		enableHTTPProxy:   true,
		enableNon200Error: true,
	}
}

/**
 * @description: 发起请求
 * @return {*response.Response} 响应对象
 */
func (r *Request) Do() *response.Response {
	var ResponseData response.Response

	if r.Error != nil {
		ResponseData.Error = r.Error
		return &ResponseData
	}

	URL := r.URL
	if r.URLValue != nil {
		URL += "?" + buildURLValue(r.URLValue)
	}

	// 创建一个http请求
	Req, err := http.NewRequest(r.Method, URL, r.Body)
	if err != nil {
		ResponseData.Error = err
		return &ResponseData
	}

	// 设置User-Agent
	Req.Header.Set("User-Agent", "yuzuhttp/1.0")

	// 设置请求头
	for Key, Value := range r.Header {
		Req.Header.Set(Key, Value)
	}

	// 创建Client
	Client := http.Client{}

	// 设置HTTP代理
	if r.enableHTTPProxy {
		Client.Transport = &http.Transport{
			Proxy: http.ProxyFromEnvironment,
		}
	}

	// 发起请求
	Resp, err := Client.Do(Req)
	if err != nil {
		ResponseData.Error = err
		return &ResponseData
	}

	ResponseData.Response = Resp

	// 判断非200状态码返回错误
	if r.enableNon200Error {
		if Resp.StatusCode != 200 {
			// 返回错误
			ResponseData.Error = fmt.Errorf("Request failed with status code %d", Resp.StatusCode)
			return &ResponseData
		}
	}

	return &ResponseData
}
