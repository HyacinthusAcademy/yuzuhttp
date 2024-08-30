/*
 * @Author: nijineko
 * @Date: 2024-08-31 04:16:57
 * @LastEditTime: 2024-08-31 04:18:18
 * @LastEditors: nijineko
 * @Description: 请求URL参数测试
 * @FilePath: \yuzuhttp\requestURLValue_test.go
 */
package yuzuhttp

import (
	"testing"
)

func TestAddURLValue(t *testing.T) {
	var Data map[string]any
	if err := Get("http://"+testServer+"/").AddURLValue("client", "yuzuhttp").Do().BodyJSON(&Data); err != nil {
		t.Error(err)
		return
	}

	if Data["url"] != "/?client=yuzuhttp" {
		t.Error("URL error")
		return
	}
}

func TestSetURLValues(t *testing.T) {
	var Data map[string]any
	if err := Get("http://"+testServer+"/").SetURLValues(map[string]string{"client": "yuzuhttp"}).Do().BodyJSON(&Data); err != nil {
		t.Error(err)
		return
	}

	if Data["url"] != "/?client=yuzuhttp" {
		t.Error("URL error")
		return
	}
}