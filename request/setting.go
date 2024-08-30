/*
 * @Author: nijineko
 * @Date: 2024-08-29 20:18:41
 * @LastEditTime: 2024-08-29 20:20:53
 * @LastEditors: nijineko
 * @Description: 参数设置
 * @FilePath: \yuzuhttp\request\setting.go
 */
package request

/**
 * @description: 设置是否启用HTTP代理环境变量支持
 * @param {bool} enable
 * @return {*Request}
 */
func (r *Request) SetEnableHTTPProxy(Enable bool) *Request {
	r.enableHTTPProxy = Enable

	return r
}

/**
 * @description: 设置是否启用非200状态码时返回错误
 * @param {bool} enable
 * @return {*Request}
 */
func (r *Request) SetEnableNon200Error(Enable bool) *Request {
	r.enableNon200Error = Enable

	return r
}
