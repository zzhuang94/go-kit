package str

import (
	"strings"
	"unicode"
)

// CamelCase 转换为驼峰命名 / Convert to camelCase
func CamelCase(s string) string {
	if s == "" {
		return s
	}
	words := splitWords(s)
	if len(words) == 0 {
		return s
	}
	result := strings.ToLower(words[0])
	for i := 1; i < len(words); i++ {
		if words[i] != "" {
			result += strings.ToUpper(words[i][:1]) + strings.ToLower(words[i][1:])
		}
	}
	return result
}

// SnakeCase 转换为蛇形命名 / Convert to snake_case
func SnakeCase(s string) string {
	return strings.ToLower(joinWords(s, "_"))
}

// KebabCase 转换为短横线命名 / Convert to kebab-case
func KebabCase(s string) string {
	return strings.ToLower(joinWords(s, "-"))
}

// TitleCase 转换为标题格式（每个单词首字母大写）/ Convert to Title Case (first letter of each word capitalized)
func TitleCase(s string) string {
	words := splitWords(s)
	result := make([]string, 0, len(words))
	for _, word := range words {
		if word != "" {
			result = append(result, strings.ToUpper(word[:1])+strings.ToLower(word[1:]))
		}
	}
	return strings.Join(result, " ")
}

// Pluralize 简单的复数化（基础规则）/ Simple pluralization (basic rules)
func Pluralize(s string) string {
	if s == "" {
		return s
	}
	lower := strings.ToLower(s)
	if strings.HasSuffix(lower, "y") && len(lower) > 1 {
		prev := lower[len(lower)-2]
		if !isVowel(prev) {
			return s[:len(s)-1] + "ies"
		}
	}
	if strings.HasSuffix(lower, "s") || strings.HasSuffix(lower, "x") ||
		strings.HasSuffix(lower, "z") || strings.HasSuffix(lower, "ch") ||
		strings.HasSuffix(lower, "sh") {
		return s + "es"
	}
	return s + "s"
}

func splitWords(s string) []string {
	if s == "" {
		return nil
	}
	var words []string
	var current strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) {
			if current.Len() > 0 {
				words = append(words, current.String())
				current.Reset()
			}
			current.WriteRune(r)
		} else if unicode.IsLower(r) || unicode.IsDigit(r) {
			current.WriteRune(r)
		} else {
			if current.Len() > 0 {
				words = append(words, current.String())
				current.Reset()
			}
		}
		if i == len(s)-1 && current.Len() > 0 {
			words = append(words, current.String())
		}
	}
	if len(words) == 0 {
		return []string{s}
	}
	return words
}

func joinWords(s string, sep string) string {
	words := splitWords(s)
	return strings.Join(words, sep)
}

func isVowel(r byte) bool {
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' ||
		r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U'
}
