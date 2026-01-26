# Network Utilities Package

提供网络操作工具函数。

## 类型定义

### Resp

`Resp` 结构体包含完整的 HTTP 响应信息：

```go
type Resp struct {
    StatusCode int            // HTTP 状态码
    Headers    http.Header    // 响应头
    Body       []byte         // 响应体
    Cookies    []*http.Cookie // Cookie
}
```

## HTTP 请求

### 基础请求方法

- `Get(url string, timeout int) ([]byte, error)` - 发送 GET 请求
- `GetWithHeaders(url string, timeout int, headers map[string]string) ([]byte, error)` - 发送带请求头的 GET 请求
- `Post(url string, data []byte, headers map[string]string, timeout int) ([]byte, error)` - 发送 POST 请求
- `Put(url string, data []byte, headers map[string]string, timeout int) ([]byte, error)` - 发送 PUT 请求
- `Delete(url string, headers map[string]string, timeout int) ([]byte, error)` - 发送 DELETE 请求
- `Patch(url string, data []byte, headers map[string]string, timeout int) ([]byte, error)` - 发送 PATCH 请求

### 通用请求方法

- `Req(method, url string, data []byte, headers map[string]string, timeout int) ([]byte, error)` - 发送通用 HTTP 请求
- `ReqFull(method, url string, data []byte, headers map[string]string, timeout int) (*Resp, error)` - 发送请求并返回完整响应信息（包括状态码、响应头、Cookie 等）

### JSON 请求方法

- `PostJson(url string, data any, headers map[string]string, timeout int) ([]byte, error)` - 发送 JSON POST 请求
- `GetJsonParse(url string, result any, timeout int) error` - 发送 GET 请求并解析 JSON 响应
- `PostJsonParse(url string, data, result any, headers map[string]string, timeout int) error` - 发送 JSON POST 请求并解析 JSON 响应

### 文件操作

- `UploadFile(url, path, name string, headers map[string]string, timeout int) ([]byte, error)` - 上传文件
- `DownloadFile(url, path string, timeout int) error` - 下载文件到本地

## IP 地址处理

- `ClientIp(req *http.Request) string` - 获取客户端 IP 地址
- `LocalIp() string` - 获取本机 IP 地址

## 使用示例

### HTTP 请求

```go
package main

import (
	"fmt"
	"github.com/zzhuang94/go-kit/net"
)

func main() {
	// GET 请求
	response, err := net.Get("https://api.example.com/get", 10)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(response))
	
	// POST 请求
	data := []byte(`{"name": "test"}`)
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	response, err = net.Post("https://api.example.com/post", data, headers, 10)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(response))
	
	// PUT 请求
	response, err = net.Put("https://api.example.com/put", data, headers, 10)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(response))
	
	// DELETE 请求
	response, err = net.Delete("https://api.example.com/delete", headers, 10)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(response))
	
	// JSON POST 请求
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	user := User{Name: "John", Age: 30}
	response, err = net.PostJson("https://api.example.com/users", user, nil, 10)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(response))
	
	// GET 请求并解析 JSON
	var result map[string]any
	err = net.GetJsonParse("https://api.example.com/user/1", &result, 10)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	
	// 上传文件
	response, err = net.UploadFile("https://api.example.com/upload", "./file.txt", "file", nil, 30)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(response))
	
	// 下载文件
	err = net.DownloadFile("https://example.com/file.pdf", "./downloads/file.pdf", 30)
	if err != nil {
		panic(err)
	}
	
	// 获取完整响应信息
	resp, err := net.ReqFull("GET", "https://api.example.com/info", nil, nil, 10)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Headers: %v\n", resp.Headers)
	fmt.Printf("Body: %s\n", string(resp.Body))
	
	// 检查状态码（使用 Resp 结构体中的 StatusCode）
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		fmt.Println("Request succeeded")
	}
}
```

### IP 地址

```go
package main

import (
	"fmt"
	"github.com/zzhuang94/go-kit/net"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// 获取客户端 IP
	clientIP := net.ClientIp(r)
	fmt.Println("Client IP:", clientIP)
}

func main() {
	// 获取本机 IP
	localIP := net.LocalIp()
	fmt.Println("Local IP:", localIP)
	
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
```
