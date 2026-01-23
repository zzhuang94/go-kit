package time

import (
	"time"
)

const (
	DateFormat     = "2006-01-02"
	TimeFormat     = "15:04:05"
	DateTimeFormat = "2006-01-02 15:04:05"
)

// Format 格式化时间为字符串 / Format time to string
func Format(t time.Time, layout string) string {
	return t.Format(layout)
}

// FormatDate 格式化日期（YYYY-MM-DD）/ Format date (YYYY-MM-DD)
func FormatDate(t time.Time) string {
	return t.Format(DateFormat)
}

// FormatTime 格式化时间（HH:mm:ss）/ Format time (HH:mm:ss)
func FormatTime(t time.Time) string {
	return t.Format(TimeFormat)
}

// FormatDateTime 格式化日期时间（YYYY-MM-DD HH:mm:ss）/ Format date and time (YYYY-MM-DD HH:mm:ss)
func FormatDateTime(t time.Time) string {
	return t.Format(DateTimeFormat)
}

// FormatTimestamp 格式化时间戳 / Format timestamp
func FormatTimestamp(t time.Time) int64 {
	return t.Unix()
}
