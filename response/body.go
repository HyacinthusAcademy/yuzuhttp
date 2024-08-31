/*
 * @Author: nijineko
 * @Date: 2024-08-31 02:10:40
 * @LastEditTime: 2024-09-01 04:44:47
 * @LastEditors: nijineko
 * @Description: 返回体处理
 * @FilePath: \yuzuhttp\response\body.go
 */
package response

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io"
	"os"
	"path"
)

/**
 * @description: 以Bytes获取返回
 * @note: 执行后将关闭Body
 * @param {*[]byte} Bytes 反序列化值
 * @return {[]byte} 返回体字节
 */
func (r *Response) BodyBytes(Value *[]byte) error {
	if r.Error != nil {
		return r.Error
	}

	var Buffer bytes.Buffer
	if _, err := io.Copy(&Buffer, r.Body); err != nil {
		return err
	}
	defer r.Body.Close()

	*Value = Buffer.Bytes()

	return nil
}

/**
 * @description: 以String获取返回
 * @note: 执行后将关闭Body
 * @param {*string} Value 反序列化值
 * @return {string} 返回体字符串
 */
func (r *Response) BodyString(Value *string) error {
	if r.Error != nil {
		return r.Error
	}

	var Body []byte
	if err := r.BodyBytes(&Body); err != nil {
		return err
	}

	*Value = string(Body)

	return nil
}

/**
 * @description: 反序列化返回内容为JSON
 * @note: 执行后将关闭Body
 * @param {any} Value 反序列化值 (需要传入指针)
 * @return {error} 错误
 */
func (r *Response) BodyJSON(Value any) error {
	if r.Error != nil {
		return r.Error
	}

	var Body []byte
	if err := r.BodyBytes(&Body); err != nil {
		return err
	}

	if err := json.Unmarshal(Body, Value); err != nil {
		return err
	}

	return nil
}

/**
 * @description: 反序列化返回内容为XML
 * @note: 执行后将关闭Body
 * @param {any} Value 反序列化值 (需要传入指针)
 * @return {error} 错误
 */
func (r *Response) BodyXML(Value any) error {
	if r.Error != nil {
		return r.Error
	}

	err := xml.NewDecoder(r.Body).Decode(Value)
	if err != nil {
		return err
	}

	return nil
}

/**
 * @description: 保存返回内容到文件
 * @note: 执行后将关闭Body
 * @param {string} FilePath 文件路径
 * @return {int64} 写入字节数
 * @return {error} 错误
 */
func (r *Response) BodySaveFile(FilePath string) (int64, error) {
	if r.Error != nil {
		return 0, r.Error
	}

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
		// 删除文件
		err := File.Close()
		if err == nil {
			err := os.Remove(FilePath)
			if err != nil {
				return 0, err
			}
		}

		return 0, err
	}
	defer r.Body.Close()

	return Size, nil
}
