package str

import (
	"strings"
)

// Trim 去除首尾空白字符 / Trim whitespace from both ends of string
func Trim(s string) string {
	return strings.TrimSpace(s)
}

// Substring 截取子串，支持负数索引 / Extract substring, supports negative indices
func Substring(s string, start, end int) string {
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

// ReplaceAll 替换所有匹配的字符串 / Replace all occurrences of string
func ReplaceAll(s, old, new string) string {
	return strings.ReplaceAll(s, old, new)
}

// ReplaceN 替换前N个匹配的字符串 / Replace first N occurrences of string
func ReplaceN(s, old, new string, n int) string {
	return strings.Replace(s, old, new, n)
}

// Contains 检查字符串是否包含子串 / Check if string contains substring
func Contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

// ContainsIgnoreCase 检查字符串是否包含子串（忽略大小写）/ Check if string contains substring (case insensitive)
func ContainsIgnoreCase(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}

// StartsWith 检查字符串是否以指定前缀开始 / Check if string starts with prefix
func StartsWith(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

// EndsWith 检查字符串是否以指定后缀结束 / Check if string ends with suffix
func EndsWith(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}

// Repeat 重复字符串指定次数 / Repeat string specified number of times
func Repeat(s string, count int) string {
	return strings.Repeat(s, count)
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

// TruncateWithEllipsis 截断字符串并添加省略号 / Truncate string and add ellipsis
func TruncateWithEllipsis(s string, length int) string {
	if len(s) <= length {
		return s
	}
	runes := []rune(s)
	if len(runes) <= length {
		return s
	}
	if length <= 3 {
		return strings.Repeat(".", length)
	}
	return string(runes[:length-3]) + "..."
}
