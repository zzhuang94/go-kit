package str

import (
	"net/url"
	"regexp"
	"strings"
)

var (
	emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	ipRegex    = regexp.MustCompile(`^(\d{1,3}\.){3}\d{1,3}$`)
	phoneRegex = regexp.MustCompile(`^1[3-9]\d{9}$`)
)

// IsEmail 验证邮箱格式 / Validate email format
func IsEmail(s string) bool {
	if s == "" {
		return false
	}
	return emailRegex.MatchString(s)
}

// IsURL 验证URL格式 / Validate URL format
func IsURL(s string) bool {
	if s == "" {
		return false
	}
	u, err := url.Parse(s)
	return err == nil && u.Scheme != "" && u.Host != ""
}

// IsIP 验证IP地址格式（IPv4）/ Validate IP address format (IPv4)
func IsIP(s string) bool {
	if s == "" {
		return false
	}
	if !ipRegex.MatchString(s) {
		return false
	}
	parts := strings.Split(s, ".")
	for _, part := range parts {
		if len(part) > 3 {
			return false
		}
		var num int
		for _, r := range part {
			if r < '0' || r > '9' {
				return false
			}
			num = num*10 + int(r-'0')
		}
		if num > 255 {
			return false
		}
	}
	return true
}

// IsPhone 验证手机号格式（中国大陆）/ Validate phone number format (Mainland China)
func IsPhone(s string) bool {
	if s == "" {
		return false
	}
	return phoneRegex.MatchString(s)
}

// IsNumeric 验证是否为数字字符串 / Validate if string is numeric
func IsNumeric(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

// IsAlpha 验证是否为字母字符串 / Validate if string is alphabetic
func IsAlpha(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}

// IsAlphaNumeric 验证是否为字母数字字符串 / Validate if string is alphanumeric
func IsAlphaNumeric(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')) {
			return false
		}
	}
	return true
}

// IsEmpty 检查字符串是否为空 / Check if string is empty
func IsEmpty(s string) bool {
	return len(s) == 0
}

// IsBlank 检查字符串是否为空白（空或只包含空白字符）/ Check if string is blank (empty or only whitespace)
func IsBlank(s string) bool {
	return strings.TrimSpace(s) == ""
}
