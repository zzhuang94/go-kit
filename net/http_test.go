package net

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// setupTestServer 创建一个测试 HTTP 服务器 / Create a test HTTP server
func setupTestServer() *httptest.Server {
	mux := http.NewServeMux()

	// GET endpoint - 返回 JSON / Return JSON
	mux.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{
			"method": "GET",
			"path":   r.URL.Path,
			"query":  r.URL.RawQuery,
		}
		json.NewEncoder(w).Encode(response)
	})

	// POST endpoint - 回显请求体 / Echo request body
	mux.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		body, _ := io.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{
			"method": "POST",
			"body":   string(body),
		}
		json.NewEncoder(w).Encode(response)
	})

	// PUT endpoint / PUT endpoint
	mux.HandleFunc("/put", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		body, _ := io.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{
			"method": "PUT",
			"body":   string(body),
		}
		json.NewEncoder(w).Encode(response)
	})

	// DELETE endpoint / DELETE endpoint
	mux.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{
			"method": "DELETE",
			"status": "deleted",
		}
		json.NewEncoder(w).Encode(response)
	})

	// PATCH endpoint / PATCH endpoint
	mux.HandleFunc("/patch", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		body, _ := io.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{
			"method": "PATCH",
			"body":   string(body),
		}
		json.NewEncoder(w).Encode(response)
	})

	// Headers endpoint - 返回请求头 / Return request headers
	mux.HandleFunc("/headers", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{
			"headers": r.Header,
		}
		json.NewEncoder(w).Encode(response)
	})

	// Status endpoint - 返回指定状态码 / Return specified status code
	mux.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		if code == "" {
			code = "200"
		}
		var statusCode int
		fmt.Sscanf(code, "%d", &statusCode)
		w.WriteHeader(statusCode)
		w.Write([]byte(fmt.Sprintf("Status: %d", statusCode)))
	})

	// JSON endpoint - 接收并返回 JSON / Receive and return JSON
	mux.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		var input map[string]interface{}
		json.NewDecoder(r.Body).Decode(&input)
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{
			"received": input,
			"echo":     true,
		}
		json.NewEncoder(w).Encode(response)
	})

	// Upload endpoint - 接收文件上传 / Receive file upload
	mux.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		err := r.ParseMultipartForm(10 << 20) // 10MB
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Failed to parse multipart form"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{
			"status": "uploaded",
			"files":  len(r.MultipartForm.File),
		}
		json.NewEncoder(w).Encode(response)
	})

	// Download endpoint - 返回文件内容 / Return file content
	mux.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=test.txt")
		w.Write([]byte("This is a test file content"))
	})

	// Cookie endpoint - 设置和返回 Cookie / Set and return Cookie
	mux.HandleFunc("/cookie", func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:  "test-cookie",
			Value: "test-value",
			Path:  "/",
		})
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{
			"cookies": len(r.Cookies()),
		}
		json.NewEncoder(w).Encode(response)
	})

	// Timeout endpoint - 延迟响应 / Delayed response
	mux.HandleFunc("/timeout", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		w.Write([]byte("Response after delay"))
	})

	return httptest.NewServer(mux)
}

func TestGet(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	data, err := Get(server.URL+"/get", 5)
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if result["method"] != "GET" {
		t.Errorf("Expected method GET, got %v", result["method"])
	}
}

func TestGetWithHeaders(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	headers := map[string]string{
		"X-Custom-Header": "test-value",
		"User-Agent":      "go-kit-test",
	}

	data, err := GetWithHeaders(server.URL+"/headers", 5, headers)
	if err != nil {
		t.Fatalf("GetWithHeaders() error = %v", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	headersMap, ok := result["headers"].(map[string]interface{})
	if !ok {
		t.Fatal("Headers not found in response")
	}

	customHeader := headersMap["X-Custom-Header"].([]interface{})
	if len(customHeader) == 0 || customHeader[0] != "test-value" {
		t.Errorf("Expected X-Custom-Header to be test-value")
	}
}

func TestPost(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	body := []byte(`{"key": "value"}`)
	headers := map[string]string{"Content-Type": "application/json"}

	data, err := Post(server.URL+"/post", body, headers, 5)
	if err != nil {
		t.Fatalf("Post() error = %v", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if result["method"] != "POST" {
		t.Errorf("Expected method POST, got %v", result["method"])
	}

	bodyStr, ok := result["body"].(string)
	if !ok || !strings.Contains(bodyStr, "key") {
		t.Errorf("Expected body to contain request data")
	}
}

func TestPut(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	body := []byte(`{"key": "updated-value"}`)
	headers := map[string]string{"Content-Type": "application/json"}

	data, err := Put(server.URL+"/put", body, headers, 5)
	if err != nil {
		t.Fatalf("Put() error = %v", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if result["method"] != "PUT" {
		t.Errorf("Expected method PUT, got %v", result["method"])
	}
}

func TestDelete(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	data, err := Delete(server.URL+"/delete", nil, 5)
	if err != nil {
		t.Fatalf("Delete() error = %v", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if result["method"] != "DELETE" {
		t.Errorf("Expected method DELETE, got %v", result["method"])
	}
}

func TestPatch(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	body := []byte(`{"key": "patched-value"}`)
	headers := map[string]string{"Content-Type": "application/json"}

	data, err := Patch(server.URL+"/patch", body, headers, 5)
	if err != nil {
		t.Fatalf("Patch() error = %v", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if result["method"] != "PATCH" {
		t.Errorf("Expected method PATCH, got %v", result["method"])
	}
}

func TestPostJson(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	data := map[string]interface{}{
		"name":  "test",
		"value": 123,
	}

	response, err := PostJson(server.URL+"/json", data, nil, 5)
	if err != nil {
		t.Fatalf("PostJson() error = %v", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(response, &result); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if result["echo"] != true {
		t.Errorf("Expected echo to be true")
	}
}

func TestGetJsonParse(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	var result map[string]interface{}
	err := GetJsonParse(server.URL+"/get", &result, 5)
	if err != nil {
		t.Fatalf("GetJsonParse() error = %v", err)
	}

	if result["method"] != "GET" {
		t.Errorf("Expected method GET, got %v", result["method"])
	}
}

func TestPostJsonParse(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	requestData := map[string]interface{}{
		"name": "test",
	}

	var result map[string]interface{}
	err := PostJsonParse(server.URL+"/json", requestData, &result, nil, 5)
	if err != nil {
		t.Fatalf("PostJsonParse() error = %v", err)
	}

	if result["echo"] != true {
		t.Errorf("Expected echo to be true")
	}
}

func TestReqFull(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	resp, err := ReqFull("GET", server.URL+"/get", nil, nil, 5)
	if err != nil {
		t.Fatalf("ReqFull() error = %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	if len(resp.Body) == 0 {
		t.Error("Expected non-empty response body")
	}

	if resp.Headers.Get("Content-Type") == "" {
		t.Error("Expected Content-Type header")
	}
}

func TestReqFullWithCookies(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	resp, err := ReqFull("GET", server.URL+"/cookie", nil, nil, 5)
	if err != nil {
		t.Fatalf("ReqFull() error = %v", err)
	}

	if len(resp.Cookies) == 0 {
		t.Error("Expected cookies in response")
	}

	found := false
	for _, cookie := range resp.Cookies {
		if cookie.Name == "test-cookie" && cookie.Value == "test-value" {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected test-cookie in response")
	}
}

func TestReqErrorHandling(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	// Test 404 error
	_, err := Req("GET", server.URL+"/notfound", nil, nil, 5)
	if err == nil {
		t.Error("Expected error for 404 status")
	}

	// Test 500 error
	_, err = Req("GET", server.URL+"/status?code=500", nil, nil, 5)
	if err == nil {
		t.Error("Expected error for 500 status")
	}
}

func TestUploadFile(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	// Create a temporary test file
	tmpDir := os.TempDir()
	testFile := filepath.Join(tmpDir, "test_upload.txt")
	defer os.Remove(testFile)

	err := os.WriteFile(testFile, []byte("test file content"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	data, err := UploadFile(server.URL+"/upload", testFile, "file", nil, 5)
	if err != nil {
		t.Fatalf("UploadFile() error = %v", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if result["status"] != "uploaded" {
		t.Errorf("Expected status uploaded, got %v", result["status"])
	}
}

func TestDownloadFile(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	// Create a temporary directory for download
	tmpDir := os.TempDir()
	downloadPath := filepath.Join(tmpDir, "downloaded_test.txt")
	defer os.Remove(downloadPath)

	err := DownloadFile(server.URL+"/download", downloadPath, 5)
	if err != nil {
		t.Fatalf("DownloadFile() error = %v", err)
	}

	// Verify file was created
	if _, err := os.Stat(downloadPath); os.IsNotExist(err) {
		t.Error("Downloaded file does not exist")
	}

	// Verify file content
	content, err := os.ReadFile(downloadPath)
	if err != nil {
		t.Fatalf("Failed to read downloaded file: %v", err)
	}

	expectedContent := "This is a test file content"
	if string(content) != expectedContent {
		t.Errorf("Expected content %s, got %s", expectedContent, string(content))
	}
}

func TestTimeout(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	// Test with short timeout
	_, err := Get(server.URL+"/timeout", 1)
	if err == nil {
		t.Error("Expected timeout error")
	}

	// Test with sufficient timeout
	data, err := Get(server.URL+"/timeout", 5)
	if err != nil {
		t.Fatalf("Get() with sufficient timeout error = %v", err)
	}

	if len(data) == 0 {
		t.Error("Expected response data")
	}
}

func TestReqWithCustomHeaders(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	headers := map[string]string{
		"Authorization": "Bearer token123",
		"X-Request-ID":  "req-123",
	}

	data, err := Req("GET", server.URL+"/headers", nil, headers, 5)
	if err != nil {
		t.Fatalf("Req() error = %v", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	headersMap, ok := result["headers"].(map[string]interface{})
	if !ok {
		t.Fatal("Headers not found in response")
	}

	authHeader := headersMap["Authorization"].([]interface{})
	if len(authHeader) == 0 || authHeader[0] != "Bearer token123" {
		t.Error("Expected Authorization header")
	}
}

func TestPostWithDefaultContentType(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	// Post without Content-Type header should default to application/json
	body := []byte(`{"test": "data"}`)
	data, err := Post(server.URL+"/post", body, nil, 5)
	if err != nil {
		t.Fatalf("Post() error = %v", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if result["method"] != "POST" {
		t.Errorf("Expected method POST, got %v", result["method"])
	}
}
