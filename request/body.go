/*
 * @Author: nijineko
 * @Date: 2024-08-29 19:53:53
 * @LastEditTime: 2024-08-30 17:13:32
 * @LastEditors: nijineko
 * @Description: 请求体处理
 * @FilePath: \yuzuhttp\request\body.go
 */
package request

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"

	"github.com/HyacinthusAcademy/yuzuhttp/request/multipartForm"
)

/**
 * @description: 设置请求体
 * @param {io.ReadCloser} Body 请求体
 * @return {*Request} 请求对象
 */
func (r *Request) SetBody(Body io.ReadCloser) *Request {
	r.Body = Body

	return r
}

/**
 * @description: 设置Bytes请求体
 * @param {[]byte} Body 请求体
 * @return {*Request} 请求对象
 */
func (r *Request) SetBodyBytes(Body []byte) *Request {
	r.Body = io.NopCloser(bytes.NewReader(Body))

	return r
}

/**
 * @description: 设置String请求体
 * @param {string} Body 请求体
 * @return {*Request} 请求对象
 */
func (r *Request) SetBodyString(Body string) *Request {
	r.SetBody(io.NopCloser(strings.NewReader(Body)))

	return r
}

/**
 * @description: 设置FormData格式请求体
 * @param {map[string]string} FormData 表单数据
 * @return {*Request} 请求对象
 */
func (r *Request) SetBodyFormData(FormData *multipartForm.Form) *Request {
	// 构建请求体
	Body, MultipartHeader, err := FormData.BuildBody()
	if err != nil {
		r.Error = err
		return r
	}

	// 设置请求头
	r.AddHeader("Content-Type", MultipartHeader)

	r.SetBody(io.NopCloser(Body))

	return r
}

/**
 * @description: 设置FormUrlencoded格式请求体
 * @param {map[string]string} FormData 表单数据
 * @return {*Request} 请求对象
 */
func (r *Request) SetBodyFormUrlencoded(FormData map[string]string) *Request {
	// 设置请求头
	r.AddHeader("Content-Type", "application/x-www-form-urlencoded")

	// 构建请求体
	FormDataStr := buildURLValue(FormData)
	r.SetBodyString(FormDataStr)

	return r
}

/**
 * @description: 设置Json格式请求体
 * @param {any} JsonData Json数据
 * @return {*Request} 请求对象
 */
func (r *Request) SetBodyJson(JsonData any) *Request {
	JsonBytes, err := json.Marshal(JsonData)
	if err != nil {
		r.Error = err
		return r
	}

	// 设置请求头
	r.AddHeader("Content-Type", "application/json")

	r.SetBodyBytes(JsonBytes)

	return r
}
