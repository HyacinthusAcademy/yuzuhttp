/*
 * @Author: nijineko
 * @Date: 2024-08-29 19:48:25
 * @LastEditTime: 2024-08-29 20:02:31
 * @LastEditors: nijineko
 * @Description: 请求头处理
 * @FilePath: \yuzuhttp\request\header.go
 */
package request

/**
 * @description: 初始化请求头
 */
func (r *Request) initHeader() {
	if r.Header == nil {
		r.Header = make(map[string]string)
	}
}

/**
 * @description: 添加请求头
 * @param {string} Key Key
 * @param {string} Value Value
 */
func (r *Request) AddHeader(Key, Value string) *Request {
	r.initHeader()

	r.Header[Key] = Value

	return r
}

/**
 * @description: 设置请求头
 * @param {map[string]string} Headers 请求头
 */
func (r *Request) SetHeaders(Headers map[string]string) *Request {
	r.initHeader()

	for k, v := range Headers {
		r.Header[k] = v
	}

	return r
}

/**
 * @description: 获取请求头
 * @param {string} Key Key
 * @return {string} Value
 */
func (r *Request) GetHeader(Key string) string {
	if r.Header == nil {
		return ""
	}

	return r.Header[Key]
}

/**
 * @description: 移除请求头
 * @param {string} key Key
 * @return {*Request} 请求对象
 */
func (r *Request) RemoveHeader(Key string) *Request {
	if r.Header != nil {
		delete(r.Header, Key)
	}

	return r
}
