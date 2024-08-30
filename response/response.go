/*
 * @Author: nijineko
 * @Date: 2024-08-29 20:06:22
 * @LastEditTime: 2024-08-29 20:06:24
 * @LastEditors: nijineko
 * @Description: 响应处理
 * @FilePath: \yuzuhttp\response\response.go
 */
package response

import "net/http"

// 响应对象，基于http.Response封装
type Response struct {
	*http.Response       // 响应
	Error          error // 错误
}
