/*
 * @Author: nijineko
 * @Date: 2024-08-29 20:00:49
 * @LastEditTime: 2024-08-31 04:10:24
 * @LastEditors: nijineko
 * @Description: 请求测试
 * @FilePath: \yuzuhttp\request_test.go
 */
package yuzuhttp

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/HyacinthusAcademy/yuzuhttp/request"
)

const (
	testServer = "localhost:8080"
)

func TestMain(m *testing.M) {
	go startTestServer()

	os.Exit(m.Run())
}

func TestRequest(t *testing.T) {
	Response := request.New("GET", "http://"+testServer+"/").Do()
	if Response.Error != nil {
		t.Error(Response.Error)
		return
	}

	var Body string
	if err := Response.BodyString(&Body); err != nil {
		t.Error(err)
		return
	}
	t.Log(Body)
}

func TestGet(t *testing.T) {
	var Data map[string]any
	if err := Get("http://" + testServer + "/").Do().BodyJSON(&Data); err != nil {
		t.Error(err)
		return
	}

	if Data["method"] != "GET" {
		t.Error("Method error")
		return
	}
}

func TestPost(t *testing.T) {
	var Data map[string]any
	if err := Post("http://" + testServer + "/").Do().BodyJSON(&Data); err != nil {
		t.Error(err)
		return
	}

	if Data["method"] != "POST" {
		t.Error("Method error")
		return
	}
}

func TestPut(t *testing.T) {
	var Data map[string]any
	if err := Put("http://" + testServer + "/").Do().BodyJSON(&Data); err != nil {
		t.Error(err)
		return
	}

	if Data["method"] != "PUT" {
		t.Error("Method error")
		return
	}
}

func TestDelete(t *testing.T) {
	var Data map[string]any
	if err := Delete("http://" + testServer + "/").Do().BodyJSON(&Data); err != nil {
		t.Error(err)
		return
	}

	if Data["method"] != "DELETE" {
		t.Error("Method error")
		return
	}
}

func TestHead(t *testing.T) {
	if err := Head("http://" + testServer + "/").Do().Error; err != nil {
		t.Error(err)
		return
	}
}

func TestConnect(t *testing.T) {
	var Data map[string]any
	if err := Connect("http://" + testServer + "/").Do().BodyJSON(&Data); err != nil {
		t.Error(err)
		return
	}

	if Data["method"] != "CONNECT" {
		t.Error("Method error")
		return
	}
}

func TestOptions(t *testing.T) {
	var Data map[string]any
	if err := Options("http://" + testServer + "/").Do().BodyJSON(&Data); err != nil {
		t.Error(err)
		return
	}

	if Data["method"] != "OPTIONS" {
		t.Error("Method error")
		return
	}
}

func TestTrace(t *testing.T) {
	var Data map[string]any
	if err := Trace("http://" + testServer + "/").Do().BodyJSON(&Data); err != nil {
		t.Error(err)
		return
	}

	if Data["method"] != "TRACE" {
		t.Error("Method error")
		return
	}
}

/**
 * @description: 启动测试服务器
 */
func startTestServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		type Response struct {
			Message    string            `json:"message"`
			Method     string            `json:"method"`
			URL        string            `json:"url"`
			Proto      string            `json:"proto"`
			RemoteAddr string            `json:"remoteAddr"`
			Header     map[string]string `json:"header"`
			Body       string            `json:"body"`
		}

		Body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		ResponseData := Response{
			Message:    "Hello, YuzuHTTP!",
			Method:     r.Method,
			URL:        r.URL.String(),
			Proto:      r.Proto,
			RemoteAddr: r.RemoteAddr,
			Header:     map[string]string{},
			Body:       string(Body),
		}

		for Key, Value := range r.Header {
			ResponseData.Header[Key] = Value[0]
		}

		ResponseJSON, err := json.Marshal(ResponseData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(ResponseJSON)
	})

	http.ListenAndServe(testServer, nil)
}
