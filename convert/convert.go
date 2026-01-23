package convert

import (
	"strconv"
)

// ToInt 字符串转整数 / Convert string to integer
func ToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

// ToInt64 字符串转int64 / Convert string to int64
func ToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

// ToFloat64 字符串转float64 / Convert string to float64
func ToFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

// ToBool 字符串转布尔值 / Convert string to boolean
func ToBool(s string) (bool, error) {
	return strconv.ParseBool(s)
}

// IntToString 整数转字符串 / Convert integer to string
func IntToString(i int) string {
	return strconv.Itoa(i)
}

// Int64ToString int64转字符串 / Convert int64 to string
func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

// Float64ToString float64转字符串 / Convert float64 to string
func Float64ToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

// BoolToString 布尔值转字符串 / Convert boolean to string
func BoolToString(b bool) string {
	return strconv.FormatBool(b)
}

// BytesToString 字节数组转字符串 / Convert byte array to string
func BytesToString(b []byte) string {
	return string(b)
}

// StringToBytes 字符串转字节数组 / Convert string to byte array
func StringToBytes(s string) []byte {
	return []byte(s)
}
