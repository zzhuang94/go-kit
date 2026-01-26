# Network Utilities Package

Provides network operation utility functions.

## Type Definitions

### Resp

The `Resp` struct contains complete HTTP response information:

```go
type Resp struct {
    StatusCode int            // HTTP status code
    Headers    http.Header    // Response headers
    Body       []byte         // Response body
    Cookies    []*http.Cookie // Cookies
}
```

## HTTP Requests

### Basic Request Methods

- `Get(url string, timeout int) ([]byte, error)` - Send GET request
- `GetWithHeaders(url string, timeout int, headers map[string]string) ([]byte, error)` - Send GET request with headers
- `Post(url string, data []byte, headers map[string]string, timeout int) ([]byte, error)` - Send POST request
- `Put(url string, data []byte, headers map[string]string, timeout int) ([]byte, error)` - Send PUT request
- `Delete(url string, headers map[string]string, timeout int) ([]byte, error)` - Send DELETE request
- `Patch(url string, data []byte, headers map[string]string, timeout int) ([]byte, error)` - Send PATCH request

### Generic Request Methods

- `Req(method, url string, data []byte, headers map[string]string, timeout int) ([]byte, error)` - Send generic HTTP request
- `ReqFull(method, url string, data []byte, headers map[string]string, timeout int) (*Resp, error)` - Send request and return full response information (including status code, headers, cookies, etc.)

### JSON Request Methods

- `PostJson(url string, data any, headers map[string]string, timeout int) ([]byte, error)` - Send JSON POST request
- `GetJsonParse(url string, result any, timeout int) error` - Send GET request and parse JSON response
- `PostJsonParse(url string, data, result any, headers map[string]string, timeout int) error` - Send JSON POST request and parse JSON response

### File Operations

- `UploadFile(url, path, name string, headers map[string]string, timeout int) ([]byte, error)` - Upload file
- `DownloadFile(url, path string, timeout int) error` - Download file to local

## IP Address Handling

- `ClientIp(req *http.Request) string` - Get client IP address
- `LocalIp() string` - Get local IP address

## Usage Examples

### HTTP Requests

```go
package main

import (
	"fmt"
	"github.com/zzhuang94/go-kit/net"
)

func main() {
	// GET request
	response, err := net.Get("https://api.example.com/get", 10)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(response))
	
	// POST request
	data := []byte(`{"name": "test"}`)
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	response, err = net.Post("https://api.example.com/post", data, headers, 10)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(response))
	
	// PUT request
	response, err = net.Put("https://api.example.com/put", data, headers, 10)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(response))
	
	// DELETE request
	response, err = net.Delete("https://api.example.com/delete", headers, 10)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(response))
	
	// JSON POST request
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
	
	// GET request and parse JSON
	var result map[string]any
	err = net.GetJsonParse("https://api.example.com/user/1", &result, 10)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	
	// Upload file
	response, err = net.UploadFile("https://api.example.com/upload", "./file.txt", "file", nil, 30)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(response))
	
	// Download file
	err = net.DownloadFile("https://example.com/file.pdf", "./downloads/file.pdf", 30)
	if err != nil {
		panic(err)
	}
	
	// Get full response information
	resp, err := net.ReqFull("GET", "https://api.example.com/info", nil, nil, 10)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Headers: %v\n", resp.Headers)
	fmt.Printf("Body: %s\n", string(resp.Body))
	
	// Check status code (using StatusCode from Resp struct)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		fmt.Println("Request succeeded")
	}
}
```

### IP Address

```go
package main

import (
	"fmt"
	"github.com/zzhuang94/go-kit/net"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Get client IP
	clientIP := net.ClientIp(r)
	fmt.Println("Client IP:", clientIP)
}

func main() {
	// Get local IP
	localIP := net.LocalIp()
	fmt.Println("Local IP:", localIP)
	
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
```
