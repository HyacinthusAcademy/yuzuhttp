/*
 * @Author: nijineko
 * @Date: 2024-08-31 04:20:21
 * @LastEditTime: 2024-08-31 04:20:23
 * @LastEditors: nijineko
 * @Description: 响应处理测试
 * @FilePath: \yuzuhttp\response_test.go
 */
package yuzuhttp

import (
	"encoding/json"
	"os"
	"testing"
)

func TestResponseBodyBytes(t *testing.T) {
	var BodyBytes []byte
	if err := Get("http://" + testServer + "/").Do().BodyBytes(&BodyBytes); err != nil {
		t.Error(err)
		return
	}

	var BodyJson map[string]any
	if err := json.Unmarshal(BodyBytes, &BodyJson); err != nil {
		t.Error(err)
		return
	}

	if BodyJson["message"] != "Hello, YuzuHTTP!" {
		t.Error("Bytes response error")
		return
	}
}

func TestResponseBodyString(t *testing.T) {
	var BodyString string
	if err := Get("http://" + testServer + "/").Do().BodyString(&BodyString); err != nil {
		t.Error(err)
		return
	}

	var BodyJson map[string]any
	if err := json.Unmarshal([]byte(BodyString), &BodyJson); err != nil {
		t.Error(err)
		return
	}

	if BodyJson["message"] != "Hello, YuzuHTTP!" {
		t.Error("String response error")
		return
	}
}

func TestResponseBodyJSON(t *testing.T) {
	var BodyJson map[string]any
	if err := Get("http://" + testServer + "/").Do().BodyJSON(&BodyJson); err != nil {
		t.Error(err)
		return
	}

	if BodyJson["message"] != "Hello, YuzuHTTP!" {
		t.Error("JSON response error")
		return
	}
}

func TestResponseBodySave(t *testing.T) {
	if _, err := Get("http://" + testServer + "/").Do().BodySaveFile("test.json"); err != nil {
		t.Error(err)
		return
	}
	defer os.Remove("test.json")

	var BodyJson map[string]any
	TestJsonFile, err := os.ReadFile("test.json")
	if err != nil {
		t.Error(err)
		return
	}

	if err := json.Unmarshal(TestJsonFile, &BodyJson); err != nil {
		t.Error(err)
		return
	}

	if BodyJson["message"] != "Hello, YuzuHTTP!" {
		t.Error("Save response error")
		return
	}
}
