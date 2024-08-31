/*
 * @Author: nijineko
 * @Date: 2024-08-31 04:12:35
 * @LastEditTime: 2024-09-01 05:06:55
 * @LastEditors: nijineko
 * @Description: 请求体测试
 * @FilePath: \yuzuhttp\requestBody_test.go
 */
package yuzuhttp

import (
	"bytes"
	"io"
	"mime/multipart"
	"testing"

	"github.com/HyacinthusAcademy/yuzuhttp/request/multipartForm"
)

func TestSetBodyBytes(t *testing.T) {
	var Data map[string]any
	if err := Post("http://" + testServer + "/").SetBodyBytes([]byte("Yuzu")).Do().BodyJSON(&Data); err != nil {
		t.Error(err)
		return
	}

	if Data["body"] != "Yuzu" {
		t.Error("Body error")
		return
	}
}

func TestSetBodyString(t *testing.T) {
	var Data map[string]any
	if err := Post("http://" + testServer + "/").SetBodyString("Yuzu").Do().BodyJSON(&Data); err != nil {
		t.Error(err)
		return
	}

	if Data["body"] != "Yuzu" {
		t.Error("Body error")
		return
	}
}

func TestSetFormData(t *testing.T) {
	// 构建测试表单
	MultipartForm := multipartForm.New()
	MultipartForm.AddField("yuzu", "yuzuhttp")
	MultipartForm.AddFile("test-file", "test.txt", []byte("yuzuhttp"))

	var Data map[string]any
	if err := Post("http://" + testServer + "/").SetBodyFormData(MultipartForm).Do().BodyJSON(&Data); err != nil {
		t.Error(err)
		return
	}

	ResponseBody := Data["body"].(string)

	// 解析返回内容
	ResponseMultipartForm := multipart.NewReader(bytes.NewReader([]byte(ResponseBody)), MultipartForm.GetBoundary())
	Part, err := ResponseMultipartForm.NextPart()
	if err != nil {
		t.Error(err)
		return
	}
	if Part.FormName() != "yuzu" {
		t.Error("Form name error")
		return
	}
	PartValue, err := io.ReadAll(Part)
	if err != nil {
		t.Error(err)
		return
	}
	if string(PartValue) != "yuzuhttp" {
		t.Error("Form value error")
		return
	}

	Part, err = ResponseMultipartForm.NextPart()
	if err != nil {
		t.Error(err)
		return
	}

	if Part.FormName() != "test-file" {
		t.Error("Form name error")
		return
	}
	if Part.FileName() != "test.txt" {
		t.Error("File name error")
		return
	}

	PartValue, err = io.ReadAll(Part)
	if err != nil {
		t.Error(err)
		return
	}
	if string(PartValue) != "yuzuhttp" {
		t.Error("Form value error")
		return
	}
}

func TestSetBodyFormUrlencoded(t *testing.T) {
	var Data map[string]any
	if err := Post("http://" + testServer + "/").SetBodyFormUrlencoded(map[string]string{"Yuzu": "HTTP"}).Do().BodyJSON(&Data); err != nil {
		t.Error(err)
		return
	}

	if Data["body"] != "Yuzu=HTTP" {
		t.Error("Body error")
		return
	}
}

func TestSetBodyJSON(t *testing.T) {
	var Data map[string]any
	if err := Post("http://" + testServer + "/").SetBodyJSON(map[string]string{"Yuzu": "HTTP"}).Do().BodyJSON(&Data); err != nil {
		t.Error(err)
		return
	}

	if Data["body"] != "{\"Yuzu\":\"HTTP\"}" {
		t.Error("Body error")
		return
	}
}

func TestSetBodyXML(t *testing.T) {
	type Reques struct {
		Yuzu string `xml:"Yuzu"`
	}
	RequestData := Reques{
		Yuzu: "HTTP",
	}

	var Data map[string]any
	if err := Post("http://" + testServer + "/").SetBodyXML(RequestData).Do().BodyJSON(&Data); err != nil {
		t.Error(err)
		return
	}

	if Data["body"] != "<Reques><Yuzu>HTTP</Yuzu></Reques>" {
		t.Error("Body error")
		return
	}
}
