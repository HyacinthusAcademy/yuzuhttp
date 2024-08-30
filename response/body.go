/*
 * @Author: nijineko
 * @Date: 2024-08-31 02:10:40
 * @LastEditTime: 2024-08-31 02:38:50
 * @LastEditors: nijineko
 * @Description: 返回体处理
 * @FilePath: \yuzuhttp\response\body.go
 */
package response

import (
	"encoding/json"
	"io"
	"os"
	"path"
)

/**
 * @description: 以Bytes获取返回
 * @note: 执行后将关闭Body
 * @return {[]byte} 返回体字节
 */
func (r *Response) BodyBytes() []byte {
	BodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return nil
	}
	defer r.Body.Close()

	return BodyBytes
}

/**
 * @description: 以String获取返回
 * @note: 执行后将关闭Body
 * @return {string} 返回体字符串
 */
func (r *Response) BodyString() string {
	return string(r.BodyBytes())
}

/**
 * @description: 反序列化返回内容为JSON
 * @note: 执行后将关闭Body
 * @param {any} Value 反序列化值 (需要传入指针)
 * @return {error} 错误
 */
func (r *Response) BodyJSON(Value any) error {
	err := json.Unmarshal(r.BodyBytes(), Value)
	if err != nil {
		return err
	}

	return nil
}

/**
 * @description: 保存返回内容到文件
 * @param {string} FilePath 文件路径
 * @return {int64} 写入字节数
 * @return {error} 错误
 */
func (r *Response) BodySaveFile(FilePath string) (int64, error) {
	err := os.MkdirAll(path.Dir(FilePath), 0644)
	if err != nil {
		return 0, err
	}

	File, err := os.OpenFile(FilePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return 0, err
	}
	defer File.Close()

	// 写入文件
	Size, err := io.Copy(File, r.Body)
	if err != nil {
		return 0, err
	}

	return Size, nil
}
