package json

import (
	"encoding/json"
	"fmt"
	"os"
)

// Marshal JSON序列化 / JSON serialization
func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// Unmarshal JSON反序列化 / JSON deserialization
func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// MarshalIndent JSON序列化（格式化）/ JSON serialization (formatted)
func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

// PrettyPrint 美化打印JSON / Pretty print JSON
func PrettyPrint(v interface{}) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}

// Get 获取JSON字段值（简单实现）/ Get JSON field value (simple implementation)
func Get(data []byte, key string) (interface{}, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, err
	}
	return m[key], nil
}

// Set 设置JSON字段值（简单实现）/ Set JSON field value (simple implementation)
func Set(data []byte, key string, value interface{}) ([]byte, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, err
	}
	m[key] = value
	return json.Marshal(m)
}

// Merge 合并JSON对象 / Merge JSON objects
func Merge(data1, data2 []byte) ([]byte, error) {
	var m1, m2 map[string]interface{}
	if err := json.Unmarshal(data1, &m1); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data2, &m2); err != nil {
		return nil, err
	}
	for k, v := range m2 {
		m1[k] = v
	}
	return json.Marshal(m1)
}

// ReadFile 从文件读取JSON / Read JSON from file
func ReadFile(path string, v interface{}) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

// WriteFile 写入JSON到文件 / Write JSON to file
func WriteFile(path string, v interface{}) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
