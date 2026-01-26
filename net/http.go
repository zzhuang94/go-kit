package net

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Resp 包含完整的 HTTP 响应信息 / Contains complete HTTP Resp information
type Resp struct {
	StatusCode int
	Headers    http.Header
	Body       []byte
	Cookies    []*http.Cookie
}

// doRequest 执行 HTTP 请求并返回响应 / Execute HTTP request and return response
func doRequest(method, url string, data []byte, headers map[string]string, timeout int) (*http.Response, []byte, error) {
	var body io.Reader
	if data != nil {
		body = bytes.NewBuffer(data)
	}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create request: %v", err)
	}

	// 设置请求头 / Set headers
	if headers != nil {
		for k, v := range headers {
			if strings.ToLower(k) == "host" {
				req.Host = v
			} else {
				req.Header.Set(k, v)
			}
		}
	} else if data != nil && req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	client := &http.Client{Timeout: time.Duration(timeout) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("%v [Timeout: %ds]", err, timeout)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read response: %v", err)
	}

	return resp, bodyBytes, nil
}

// Req 发送通用 HTTP 请求 / Send generic HTTP request
func Req(method, url string, data []byte, headers map[string]string, timeout int) ([]byte, error) {
	resp, bodyBytes, err := doRequest(method, url, data, headers, timeout)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return bodyBytes, fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(bodyBytes))
	}

	return bodyBytes, nil
}

// ReqFull 发送请求并返回完整响应信息 / Send request and return full response information
func ReqFull(method, url string, data []byte, headers map[string]string, timeout int) (*Resp, error) {
	resp, bodyBytes, err := doRequest(method, url, data, headers, timeout)
	if err != nil {
		return nil, err
	}

	return &Resp{
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
		Body:       bodyBytes,
		Cookies:    resp.Cookies(),
	}, nil
}

// Get 发送 GET 请求 / Send GET Req
func Get(url string, timeout int) ([]byte, error) {
	return Req("GET", url, nil, nil, timeout)
}

// GetWithHeaders 发送带请求头的 GET 请求 / Send GET Req with headers
func GetWithHeaders(url string, timeout int, headers map[string]string) ([]byte, error) {
	return Req("GET", url, nil, headers, timeout)
}

// Post 发送 POST 请求 / Send POST Req
func Post(url string, data []byte, headers map[string]string, timeout int) ([]byte, error) {
	if headers == nil {
		headers = map[string]string{"Content-Type": "application/json"}
	}
	return Req("POST", url, data, headers, timeout)
}

// Put 发送 PUT 请求 / Send PUT Req
func Put(url string, data []byte, headers map[string]string, timeout int) ([]byte, error) {
	return Req("PUT", url, data, headers, timeout)
}

// Delete 发送 DELETE 请求 / Send DELETE Req
func Delete(url string, headers map[string]string, timeout int) ([]byte, error) {
	return Req("DELETE", url, nil, headers, timeout)
}

// Patch 发送 PATCH 请求 / Send PATCH Req
func Patch(url string, data []byte, headers map[string]string, timeout int) ([]byte, error) {
	return Req("PATCH", url, data, headers, timeout)
}

// PostJson 发送 JSON POST 请求 / Send JSON POST Req
func PostJson(url string, data any, headers map[string]string, timeout int) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSON: %v", err)
	}

	if headers == nil {
		headers = make(map[string]string)
	}
	headers["Content-Type"] = "application/json"

	return Post(url, jsonData, headers, timeout)
}

// GetJsonParse 发送 GET 请求并解析 JSON 响应 / Send GET Req and parse JSON Resp
func GetJsonParse(url string, result any, timeout int) error {
	data, err := Get(url, timeout)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, result)
}

// PostJsonParse 发送 JSON POST 请求并解析 JSON 响应 / Send JSON POST Req and parse JSON Resp
func PostJsonParse(url string, data, result any, headers map[string]string, timeout int) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}

	if headers == nil {
		headers = make(map[string]string)
	}
	headers["Content-Type"] = "application/json"

	Resp, err := Post(url, jsonData, headers, timeout)
	if err != nil {
		return err
	}

	return json.Unmarshal(Resp, result)
}

// UploadFile 上传文件 / Upload file
func UploadFile(url, path, name string, headers map[string]string, timeout int) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile(name, filepath.Base(path))
	if err != nil {
		return nil, fmt.Errorf("failed to create form file: %v", err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return nil, fmt.Errorf("failed to copy file: %v", err)
	}

	err = writer.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to close writer: %v", err)
	}

	if headers == nil {
		headers = make(map[string]string)
	}
	headers["Content-Type"] = writer.FormDataContentType()

	return Req("POST", url, body.Bytes(), headers, timeout)
}

// DownloadFile 下载文件到本地 / Download file to local
func DownloadFile(url, path string, timeout int) error {
	data, err := Req("GET", url, nil, nil, timeout)
	if err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	// 创建目录 / Create directory
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}

	// 创建文件并写入 / Create file and write
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}

	return nil
}
