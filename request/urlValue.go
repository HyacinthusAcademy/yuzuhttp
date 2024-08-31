/*
 * @Author: nijineko
 * @Date: 2024-08-29 19:54:36
 * @LastEditTime: 2024-08-30 09:04:28
 * @LastEditors: nijineko
 * @Description: URL参数处理
 * @FilePath: \yuzuhttp\request\urlValue.go
 */
package request

import "net/url"

/**
 * @description: 初始化
 */
func (r *Request) initURLValue() {
	if r.URLValue == nil {
		r.URLValue = make(map[string]string)
	}
}

/**
 * @description: 构建URL参数，将参数组成URL编码的字符串，例如 (a=1&b=2)
 * @param {map[string]string} urlValues URL参数
 * @return {string} URL参数
 */
func buildURLValue(urlValues map[string]string) string {
	var Params string
	for key, Value := range urlValues {
		// 转译参数
		key = url.QueryEscape(key)
		Value = url.QueryEscape(Value)

		Params += key + "=" + Value + "&"
	}

	// 去掉最后一个&
	Params = Params[:len(Params)-1]

	return Params
}

/**
 * @description: 添加URL参数
 * @param {string} Key Key
 * @param {string} Value Value
 * @return {*Request} 请求对象
 */
func (r *Request) AddURLValue(Key, Value string) *Request {
	r.initURLValue()

	r.URLValue[Key] = Value

	return r
}

/**
 * @description: 设置URL参数
 * @param {map[string]string} Params GET参数
 * @return {*Request} 请求对象
 */
func (r *Request) SetURLValues(Params map[string]string) *Request {
	r.initURLValue()

	for k, v := range Params {
		r.URLValue[k] = v
	}

	return r
}

/**
 * @description: 获取URL参数
 * @param {string} Key Key
 * @return {string} Value
 */
func (r *Request) GetURLValue(Key string) string {
	if r.URLValue == nil {
		return ""
	}

	return r.URLValue[Key]
}

/**
 * @description: 移除URL参数
 * @param {string} Key Key
 * @return {*Request} 请求对象
 */
func (r *Request) RemoveURLValue(Key string) *Request {
	if r.URLValue != nil {
		delete(r.URLValue, Key)
	}

	return r
}
