package net

import (
	"net/http/httptest"
	"testing"
)

func TestClientIp(t *testing.T) {
	tests := []struct {
		name           string
		headers        map[string]string
		remoteAddr     string
		expectedResult string
	}{
		{
			name:           "X-real-ip header",
			headers:        map[string]string{"X-real-ip": "192.168.1.100"},
			remoteAddr:     "127.0.0.1:8080",
			expectedResult: "192.168.1.100",
		},
		{
			name:           "x-forwarded-for header",
			headers:        map[string]string{"x-forwarded-for": "10.0.0.1"},
			remoteAddr:     "127.0.0.1:8080",
			expectedResult: "10.0.0.1",
		},
		{
			name:           "Proxy-Client-IP header",
			headers:        map[string]string{"Proxy-Client-IP": "172.16.0.1"},
			remoteAddr:     "127.0.0.1:8080",
			expectedResult: "172.16.0.1",
		},
		{
			name:           "WL-Proxy-Client-IP header",
			headers:        map[string]string{"WL-Proxy-Client-IP": "203.0.113.1"},
			remoteAddr:     "127.0.0.1:8080",
			expectedResult: "203.0.113.1",
		},
		{
			name:           "Unknown header should fallback to RemoteAddr",
			headers:        map[string]string{"X-real-ip": "UNKNOWN"},
			remoteAddr:     "192.168.1.200:8080",
			expectedResult: "192.168.1.200",
		},
		{
			name:           "No headers should use RemoteAddr",
			headers:        nil,
			remoteAddr:     "10.0.0.50:8080",
			expectedResult: "10.0.0.50",
		},
		{
			name:           "IPv6 loopback should convert to 127.0.0.1",
			headers:        nil,
			remoteAddr:     "[::1]:8080",
			expectedResult: "127.0.0.1",
		},
		{
			name:           "Empty X-real-ip should check next header",
			headers:        map[string]string{"X-real-ip": "", "x-forwarded-for": "192.168.1.1"},
			remoteAddr:     "127.0.0.1:8080",
			expectedResult: "192.168.1.1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "http://example.com", nil)
			req.RemoteAddr = tt.remoteAddr

			if tt.headers != nil {
				for k, v := range tt.headers {
					req.Header.Set(k, v)
				}
			}

			result := ClientIp(req)
			if result != tt.expectedResult {
				t.Errorf("ClientIp() = %v, want %v", result, tt.expectedResult)
			}
		})
	}
}

func TestLocalIp(t *testing.T) {
	ip := LocalIp()
	if ip == "" {
		t.Error("LocalIp() returned empty string")
	}
	if ip != "127.0.0.1" {
		t.Logf("LocalIp() = %v (may vary by environment)", ip)
	}
	// LocalIp should return a valid IP address
	// It should either return a non-loopback IP or fallback to 127.0.0.1
	if ip != "127.0.0.1" {
		// If not loopback, verify it's a valid IP format
		// This is a basic check - in real scenarios, you might want more validation
		if len(ip) < 7 || len(ip) > 15 {
			t.Errorf("LocalIp() returned invalid IP format: %v", ip)
		}
	}
}
