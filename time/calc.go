package time

import (
	"time"
)

// AddDays 添加天数 / Add days
func AddDays(t time.Time, days int) time.Time {
	return t.AddDate(0, 0, days)
}

// AddHours 添加小时 / Add hours
func AddHours(t time.Time, hours int) time.Time {
	return t.Add(time.Duration(hours) * time.Hour)
}

// AddMinutes 添加分钟 / Add minutes
func AddMinutes(t time.Time, minutes int) time.Time {
	return t.Add(time.Duration(minutes) * time.Minute)
}

// DiffDays 计算天数差 / Calculate days difference
func DiffDays(t1, t2 time.Time) int {
	return int(t1.Sub(t2).Hours() / 24)
}

// DiffHours 计算小时差 / Calculate hours difference
func DiffHours(t1, t2 time.Time) int64 {
	return int64(t1.Sub(t2).Hours())
}

// DiffMinutes 计算分钟差 / Calculate minutes difference
func DiffMinutes(t1, t2 time.Time) int64 {
	return int64(t1.Sub(t2).Minutes())
}

// IsToday 判断是否为今天 / Check if time is today
func IsToday(t time.Time) bool {
	now := time.Now()
	return t.Year() == now.Year() && t.Month() == now.Month() && t.Day() == now.Day()
}

// IsYesterday 判断是否为昨天 / Check if time is yesterday
func IsYesterday(t time.Time) bool {
	return IsToday(AddDays(t, 1))
}

// IsTomorrow 判断是否为明天 / Check if time is tomorrow
func IsTomorrow(t time.Time) bool {
	return IsToday(AddDays(t, -1))
}

// StartOfDay 获取一天的开始时间 / Get start of day
func StartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// EndOfDay 获取一天的结束时间 / Get end of day
func EndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location())
}

// StartOfWeek 获取一周的开始时间（周一）/ Get start of week (Monday)
func StartOfWeek(t time.Time) time.Time {
	weekday := int(t.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	return StartOfDay(AddDays(t, -weekday+1))
}

// StartOfMonth 获取一月的开始时间 / Get start of month
func StartOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

// StartOfYear 获取一年的开始时间 / Get start of year
func StartOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
}
