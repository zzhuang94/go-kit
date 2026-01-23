# Network Utilities Package

Provides network operation utility functions.

## HTTP Requests

- `Post(url string, data []byte, headers map[string]string, timeoutSecond int) ([]byte, error)` - Send POST request
- `Get(url string, timeoutSecond int) ([]byte, error)` - Send GET request
- `GetWithHeaders(url string, timeoutSecond int, headers map[string]string) ([]byte, error)` - Send GET request with headers

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
	// POST request
	data := []byte(`{"name": "test"}`)
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	response, err := net.Post("https://api.example.com/post", data, headers, 10)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(response))
	
	// GET request
	response, err = net.Get("https://api.example.com/get", 10)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(response))
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
