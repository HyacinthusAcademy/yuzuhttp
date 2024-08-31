/*
 * @Author: nijineko
 * @Date: 2024-09-01 04:30:52
 * @LastEditTime: 2024-09-01 04:32:09
 * @LastEditors: nijineko
 * @Description: Cookie测试
 * @FilePath: \yuzuhttp\requestCookie_test.go
 */
package yuzuhttp

import "testing"

func TestAddCookie(t *testing.T) {
	var Data map[string]any
	if err := Get("http://"+testServer+"/").AddCookie("client", "yuzuhttp").Do().BodyJSON(&Data); err != nil {
		t.Error(err)
		return
	}

	if Data["header"].(map[string]any)["Cookie"] != "client=yuzuhttp" {
		t.Error("Cookie error")
		return
	}
}

func TestSetCookies(t *testing.T) {
	var Data map[string]any
	if err := Get("http://" + testServer + "/").SetCookies(map[string]string{"client": "yuzuhttp"}).Do().BodyJSON(&Data); err != nil {
		t.Error(err)
		return
	}

	if Data["header"].(map[string]any)["Cookie"] != "client=yuzuhttp" {
		t.Error("Cookie error")
		return
	}
}
