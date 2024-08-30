/*
 * @Author: nijineko
 * @Date: 2024-08-31 04:13:43
 * @LastEditTime: 2024-08-31 04:15:55
 * @LastEditors: nijineko
 * @Description: 请求头测试
 * @FilePath: \yuzuhttp\requestHeader_test.go
 */
package yuzuhttp

import "testing"

func TestAddHeader(t *testing.T) {
	var Data map[string]any
	if err := Get("http://"+testServer+"/").AddHeader("Client", "yuzuhttp").Do().BodyJSON(&Data); err != nil {
		t.Error(err)
		return
	}

	if Data["header"].(map[string]any)["Client"] != "yuzuhttp" {
		t.Error("Header error")
		return
	}
}

func TestSetHeaders(t *testing.T) {
	var Data map[string]any
	if err := Get("http://" + testServer + "/").SetHeaders(map[string]string{"Client": "yuzuhttp"}).Do().BodyJSON(&Data); err != nil {
		t.Error(err)
		return
	}

	if Data["header"].(map[string]any)["Client"] != "yuzuhttp" {
		t.Error("Header error")
		return
	}
}
