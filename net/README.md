# Network Utilities Package

提供网络操作工具函数。

## HTTP 请求

- `Post(url string, data []byte, headers map[string]string, timeoutSecond int) ([]byte, error)` - 发送 POST 请求
- `Get(url string, timeoutSecond int) ([]byte, error)` - 发送 GET 请求
- `GetWithHeaders(url string, timeoutSecond int, headers map[string]string) ([]byte, error)` - 发送带请求头的 GET 请求

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
	// POST 请求
	data := []byte(`{"name": "test"}`)
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	response, err := net.Post("https://api.example.com/post", data, headers, 10)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(response))
	
	// GET 请求
	response, err = net.Get("https://api.example.com/get", 10)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(response))
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
