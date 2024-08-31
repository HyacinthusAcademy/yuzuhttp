/*
 * @Author: nijineko
 * @Date: 2024-09-01 04:22:54
 * @LastEditTime: 2024-09-01 04:22:59
 * @LastEditors: nijineko
 * @Description: cookie处理
 * @FilePath: \yuzuhttp\request\cookie.go
 */
package request

/**
 * @description: 初始化Cookie
 */
func (r *Request) initCookie() {
	if r.Cookie == nil {
		r.Cookie = make(map[string]string)
	}
}

/**
 * @description: 添加Cookie
 * @param {string} Key Key
 * @param {string} Value Value
 * @return {*Request} 请求对象
 */
func (r *Request) AddCookie(Key, Value string) *Request {
	r.initCookie()

	r.Cookie[Key] = Value

	return r
}

/**
 * @description: 设置Cookie
 * @param {map[string]string} Cookies Cookie
 * @return {*Request} 请求对象
 */
func (r *Request) SetCookies(Cookies map[string]string) *Request {
	r.initCookie()

	for k, v := range Cookies {
		r.Cookie[k] = v
	}

	return r
}

/**
 * @description: 获取Cookie
 * @param {string} Key Key
 * @return {string} Value
 */
func (r *Request) GetCookie(Key string) string {
	if r.Cookie == nil {
		return ""
	}

	return r.Cookie[Key]
}

/**
 * @description: 移除Cookie
 * @param {string} Key Key
 * @return {*Request} 请求对象
 */
func (r *Request) RemoveCookie(Key string) *Request {
	if r.Cookie != nil {
		delete(r.Cookie, Key)
	}

	return r
}
