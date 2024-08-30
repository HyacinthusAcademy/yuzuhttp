/*
 * @Author: nijineko
 * @Date: 2024-08-30 16:37:03
 * @LastEditTime: 2024-08-30 17:09:25
 * @LastEditors: nijineko
 * @Description: 表单处理封装
 * @FilePath: \yuzuhttp\request\multipartForm\multipartForm.go
 */
package multipartForm

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"math/big"
	"mime/multipart"
)

// multipart表单结构数据
type formData struct {
	value    []byte  // 数据
	filename *string // 文件名
}

// Multipart表单结构
type Form struct {
	datas    map[string]formData
	boundary string
}

/**
 * @description: 构建请求体
 * @return {*bytes.Buffer} 请求体
 */
func (m *Form) BuildBody() (*bytes.Buffer, string, error) {
	// 构建请求体
	Body := &bytes.Buffer{}
	MultipartBody := multipart.NewWriter(Body)

	// 设置分隔符
	err := MultipartBody.SetBoundary(m.boundary)
	if err != nil {
		return nil, "", err
	}

	for Key, Value := range m.datas {
		if Value.filename == nil {
			// 添加字段
			Part, err := MultipartBody.CreateFormField(Key)
			if err != nil {
				return nil, "", err
			}

			_, err = Part.Write(Value.value)
			if err != nil {
				return nil, "", err
			}
		} else {
			// 添加文件
			Part, err := MultipartBody.CreateFormFile(Key, *Value.filename)
			if err != nil {
				return nil, "", err
			}

			_, err = Part.Write(Value.value)
			if err != nil {
				return nil, "", err
			}
		}
	}
	MultipartBody.Close()

	return Body, fmt.Sprintf("multipart/form-data; boundary=%s", m.boundary), nil
}

/**
 * @description: 创建Multipart表单
 * @return {*multipartForm} Multipart表单
 */
func New() *Form {
	return &Form{
		datas:    make(map[string]formData),
		boundary: createRandomBoundary(),
	}
}

/**
 * @description: 表单添加字段
 * @param {string} Key Key
 * @param {string} Value Value
 * @return {*multipartForm} Multipart表单
 */
func (m *Form) AddField(Key, Value string) *Form {
	m.datas[Key] = formData{
		value: []byte(Value),
	}

	return m
}

/**
 * @description: 表单添加文件
 * @param {string} Key
 * @param {[]byte} Value
 * @param {string} Filename
 * @return {*}
 */
func (m *Form) AddFile(Key string, Filename string, Value []byte) *Form {
	m.datas[Key] = formData{
		value:    Value,
		filename: &Filename,
	}

	return m
}

/**
 * @description: 获取表单数据
 * @param {string} Key Key
 * @return {[]byte} Value Value
 */
func (m *Form) Get(Key string) []byte {
	if Data, ok := m.datas[Key]; ok {
		return Data.value
	}

	return nil
}

/**
 * @description: 移除表单数据
 * @param {string} Key Key
 */
func (m *Form) Remove(Key string) {
	delete(m.datas, Key)
}

/**
 * @description: 创建随机分隔符
 * @return {string} 分隔符
 */
func createRandomBoundary() string {
	RandomStr, err := generateRandomString(14)
	if err != nil {
		panic(err)
	}

	return "------yuzuhttpFormBoundary" + RandomStr
}

/**
 * @description: 创建随机字符串
 * @param {int} length 字符串长度
 * @return {string} 随机字符串
 * @return {error} 错误
 */
func generateRandomString(length int) (string, error) {
	const Dictionary = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	Result := make([]byte, length)
	for i := range Result {
		randNum, err := rand.Int(rand.Reader, big.NewInt(int64(len(Dictionary))))
		if err != nil {
			return "", err
		}
		Result[i] = Dictionary[randNum.Int64()]
	}

	return string(Result), nil
}
