package str

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"strings"

	uuid "github.com/nu7hatch/gouuid"
)

// Trim 去除首尾空白字符 / Trim whitespace from both ends of string
func Trim(s string) string {
	return strings.TrimSpace(s)
}

func Md5(params any) string {
	data, _ := json.Marshal(params)
	mbs := md5.Sum(data)
	return fmt.Sprintf("%x", mbs)
}
func Uuid() string {
	u4, err := uuid.NewV4()
	if err != nil {
		panic(err.Error())
	}
	return fmt.Sprintf("%x", u4[0:])
}

// Sub 截取子串，支持负数索引 / Extract substring, supports negative indices
func Sub(s string, start, end int) string {
	runes := []rune(s)
	length := len(runes)
	if start < 0 {
		start = length + start
	}
	if end < 0 {
		end = length + end
	}
	if start < 0 {
		start = 0
	}
	if end > length {
		end = length
	}
	if start >= end {
		return ""
	}
	return string(runes[start:end])
}

// Reverse 反转字符串 / Reverse string
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// PadLeft 左填充字符串到指定长度 / Pad string to specified length on the left
func PadLeft(s string, length int, pad string) string {
	if len(s) >= length {
		return s
	}
	padLen := length - len(s)
	return strings.Repeat(pad, padLen) + s
}

// PadRight 右填充字符串到指定长度 / Pad string to specified length on the right
func PadRight(s string, length int, pad string) string {
	if len(s) >= length {
		return s
	}
	padLen := length - len(s)
	return s + strings.Repeat(pad, padLen)
}

// Truncate 截断字符串到指定长度 / Truncate string to specified length
func Truncate(s string, length int) string {
	if len(s) <= length {
		return s
	}
	runes := []rune(s)
	if len(runes) <= length {
		return s
	}
	return string(runes[:length])
}

func SplitLines(str string) []string {
	ans := make([]string, 0)
	m := make(map[string]bool)
	for line := range strings.SplitSeq(str, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if _, ok := m[line]; ok {
			continue
		}
		m[line] = true
		ans = append(ans, line)
	}
	return ans
}
