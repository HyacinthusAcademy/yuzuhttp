/*
 * @Author: nijineko
 * @Date: 2024-08-29 19:36:36
 * @LastEditTime: 2024-08-29 19:36:41
 * @LastEditors: nijineko
 * @Description: 初始化请求
 * @FilePath: \yuzuhttp\init.go
 */
package yuzuhttp

import "github.com/HyacinthusAcademy/yuzuhttp/request"

/**
 * @description: 创建GET请求
 * @param {string} url 请求地址
 * @return {*request.Request} 请求对象
 */
func Get(url string) *request.Request {
	return request.New("GET", url)
}

/**
 * @description: 创建HEAD请求
 * @param {string} url 请求地址
 * @return {*request.Request} 请求对象
 */
func Head(url string) *request.Request {
	return request.New("HEAD", url)
}

/**
 * @description: 创建POST请求
 * @param {string} url 请求地址
 * @return {*request.Request} 请求对象
 */
func Post(url string) *request.Request {
	return request.New("POST", url)
}

/**
 * @description: 创建PUT请求
 * @param {string} url 请求地址
 * @return {*request.Request} 请求对象
 */
func Put(url string) *request.Request {
	return request.New("PUT", url)
}

/**
 * @description: 创建DELETE请求
 * @param {string} url 请求地址
 * @return {*request.Request} 请求对象
 */
func Delete(url string) *request.Request {
	return request.New("DELETE", url)
}

/**
 * @description: 创建CONNECT请求
 * @param {string} url 请求地址
 * @return {*request.Request} 请求对象
 */
func Connect(url string) *request.Request {
	return request.New("CONNECT", url)
}

/**
 * @description: 创建OPTIONS请求
 * @param {string} url 请求地址
 * @return {*request.Request} 请求对象
 */
func Options(url string) *request.Request {
	return request.New("OPTIONS", url)
}

/**
 * @description: 创建TRACE请求
 * @param {string} url 请求地址
 * @return {*request.Request} 请求对象
 */
func Trace(url string) *request.Request {
	return request.New("TRACE", url)
}

/**
 * @description: 创建PATCH请求
 * @param {string} url 请求地址
 * @return {*request.Request} 请求对象
 */
func Patch(url string) *request.Request {
	return request.New("PATCH", url)
}
