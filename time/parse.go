package time

import (
	"time"
)

// Parse 解析时间字符串 / Parse time string
func Parse(layout, value string) (time.Time, error) {
	return time.Parse(layout, value)
}

// ParseDate 解析日期字符串（YYYY-MM-DD）/ Parse date string (YYYY-MM-DD)
func ParseDate(value string) (time.Time, error) {
	return time.Parse(DateFormat, value)
}

// ParseTime 解析时间字符串（HH:mm:ss）/ Parse time string (HH:mm:ss)
func ParseTime(value string) (time.Time, error) {
	return time.Parse(TimeFormat, value)
}

// ParseDateTime 解析日期时间字符串（YYYY-MM-DD HH:mm:ss）/ Parse date and time string (YYYY-MM-DD HH:mm:ss)
func ParseDateTime(value string) (time.Time, error) {
	return time.Parse(DateTimeFormat, value)
}
