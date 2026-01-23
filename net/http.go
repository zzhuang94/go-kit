package net

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func Post(url string, data []byte, headers map[string]string, timeoutSecond int) ([]byte, error) {
	if headers == nil {
		headers = map[string]string{"Content-Type": "application/json"}
	}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
	for k, v := range headers {
		if strings.ToLower(k) == "host" {
			req.Host = v
		} else {
			req.Header.Set(k, v)
		}
	}

	client := &http.Client{Timeout: time.Duration(timeoutSecond) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%v [Timeout: %ds]", err, timeoutSecond)
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func Get(url string, timeoutSecond int) ([]byte, error) {
	client := &http.Client{Timeout: time.Duration(timeoutSecond) * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("%v [Timeout: %ds]", err, timeoutSecond)
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func GetWithHeaders(url string, timeoutSecond int, headers map[string]string) ([]byte, error) {
	req, _ := http.NewRequest("GET", url, nil)
	for k, v := range headers {
		if strings.ToLower(k) == "host" {
			req.Host = v
		} else {
			req.Header.Set(k, v)
		}
	}
	client := &http.Client{Timeout: time.Duration(timeoutSecond) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%v [Timeout: %ds]", err, timeoutSecond)
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
