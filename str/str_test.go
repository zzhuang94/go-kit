package str

import (
	"testing"
)

func TestTrim(t *testing.T) {
	if Trim("  hello  ") != "hello" {
		t.Error("Trim failed")
	}
	if Trim("\t\nworld\n\t") != "world" {
		t.Error("Trim failed")
	}
}

func TestSub(t *testing.T) {
	s := "Hello, World"
	if Sub(s, 0, 5) != "Hello" {
		t.Error("Substring failed")
	}
	if Sub(s, -5, -1) != "Worl" {
		t.Error("Substring with negative index failed")
	}
	if Sub(s, 10, 5) != "" {
		t.Error("Substring with invalid range failed")
	}
}

func TestReverse(t *testing.T) {
	if Reverse("hello") != "olleh" {
		t.Errorf("Expected 'olleh', got %q", Reverse("hello"))
	}
	if Reverse("世界") != "界世" {
		t.Errorf("Reverse failed for unicode")
	}
}

func TestPadLeft(t *testing.T) {
	result := PadLeft("hello", 10, "0")
	if result != "00000hello" {
		t.Errorf("Expected '00000hello', got %q", result)
	}
}

func TestPadRight(t *testing.T) {
	result := PadRight("hello", 10, "0")
	if result != "hello00000" {
		t.Errorf("Expected 'hello00000', got %q", result)
	}
}

func TestTruncate(t *testing.T) {
	if Truncate("hello world", 5) != "hello" {
		t.Error("Truncate failed")
	}
	if Truncate("hi", 5) != "hi" {
		t.Error("Truncate failed")
	}
}

func TestCamelCase(t *testing.T) {
	if CamelCase("hello_world") != "helloWorld" {
		t.Errorf("Expected 'helloWorld', got %q", CamelCase("hello_world"))
	}
	if CamelCase("HelloWorld") != "helloWorld" {
		t.Errorf("Expected 'helloWorld', got %q", CamelCase("HelloWorld"))
	}
}

func TestSnakeCase(t *testing.T) {
	if SnakeCase("HelloWorld") != "hello_world" {
		t.Errorf("Expected 'hello_world', got %q", SnakeCase("HelloWorld"))
	}
}

func TestKebabCase(t *testing.T) {
	if KebabCase("HelloWorld") != "hello-world" {
		t.Errorf("Expected 'hello-world', got %q", KebabCase("HelloWorld"))
	}
}

func TestTitleCase(t *testing.T) {
	result := TitleCase("hello world")
	if result != "Hello World" {
		t.Errorf("Expected 'Hello World', got %q", result)
	}
}

func TestPluralize(t *testing.T) {
	if Pluralize("cat") != "cats" {
		t.Error("Pluralize failed")
	}
	if Pluralize("city") != "cities" {
		t.Error("Pluralize failed")
	}
	if Pluralize("box") != "boxes" {
		t.Error("Pluralize failed")
	}
}

func TestIsEmail(t *testing.T) {
	if !IsEmail("test@example.com") {
		t.Error("IsEmail failed")
	}
	if IsEmail("invalid-email") {
		t.Error("IsEmail failed")
	}
}

func TestIsURL(t *testing.T) {
	if !IsURL("https://www.example.com") {
		t.Error("IsURL failed")
	}
	if IsURL("not-a-url") {
		t.Error("IsURL failed")
	}
}

func TestIsIP(t *testing.T) {
	if !IsIP("192.168.1.1") {
		t.Error("IsIP failed")
	}
	if IsIP("256.1.1.1") {
		t.Error("IsIP failed")
	}
	if IsIP("not-an-ip") {
		t.Error("IsIP failed")
	}
}

func TestIsPhone(t *testing.T) {
	if !IsPhone("13800138000") {
		t.Error("IsPhone failed")
	}
	if IsPhone("1234567890") {
		t.Error("IsPhone failed")
	}
}

func TestIsNumeric(t *testing.T) {
	if !IsNumeric("12345") {
		t.Error("IsNumeric failed")
	}
	if IsNumeric("123abc") {
		t.Error("IsNumeric failed")
	}
}

func TestIsAlpha(t *testing.T) {
	if !IsAlpha("Hello") {
		t.Error("IsAlpha failed")
	}
	if IsAlpha("Hello123") {
		t.Error("IsAlpha failed")
	}
}

func TestIsAlphaNumeric(t *testing.T) {
	if !IsAlphaNumeric("Hello123") {
		t.Error("IsAlphaNumeric failed")
	}
	if IsAlphaNumeric("Hello-123") {
		t.Error("IsAlphaNumeric failed")
	}
}

func TestIsEmpty(t *testing.T) {
	if !IsEmpty("") {
		t.Error("IsEmpty failed")
	}
	if IsEmpty("hello") {
		t.Error("IsEmpty failed")
	}
}

func TestIsBlank(t *testing.T) {
	if !IsBlank("") {
		t.Error("IsBlank failed")
	}
	if !IsBlank("   ") {
		t.Error("IsBlank failed")
	}
	if IsBlank("hello") {
		t.Error("IsBlank failed")
	}
}

func TestMd5(t *testing.T) {
	result := Md5("hello")
	if len(result) != 32 {
		t.Errorf("Expected MD5 hash length 32, got %d", len(result))
	}

	result2 := Md5(map[string]string{"key": "value"})
	if len(result2) != 32 {
		t.Errorf("Expected MD5 hash length 32, got %d", len(result2))
	}

	// Same input should produce same hash
	result3 := Md5("hello")
	if result != result3 {
		t.Error("MD5 should produce consistent results")
	}
}

func TestUuid(t *testing.T) {
	uuid1 := Uuid()
	uuid2 := Uuid()

	if len(uuid1) == 0 {
		t.Error("Uuid should not be empty")
	}
	if len(uuid2) == 0 {
		t.Error("Uuid should not be empty")
	}

	// UUIDs should be different
	if uuid1 == uuid2 {
		t.Error("Uuid should generate unique values")
	}

	// UUID should be hexadecimal string
	if len(uuid1) < 32 {
		t.Errorf("Uuid length should be at least 32, got %d", len(uuid1))
	}
}

func TestSplitLines(t *testing.T) {
	// Test basic splitting
	input := "line1\nline2\nline3"
	result := SplitLines(input)
	if len(result) != 3 {
		t.Errorf("Expected 3 lines, got %d", len(result))
	}
	if result[0] != "line1" || result[1] != "line2" || result[2] != "line3" {
		t.Error("SplitLines failed to split correctly")
	}

	// Test with empty lines
	input2 := "line1\n\nline2\n  \nline3"
	result2 := SplitLines(input2)
	if len(result2) != 3 {
		t.Errorf("Expected 3 lines (empty lines filtered), got %d", len(result2))
	}

	// Test with duplicate lines
	input3 := "line1\nline2\nline1\nline3"
	result3 := SplitLines(input3)
	if len(result3) != 3 {
		t.Errorf("Expected 3 unique lines, got %d", len(result3))
	}

	// Test with whitespace trimming
	input4 := "  line1  \n  line2  \n  line3  "
	result4 := SplitLines(input4)
	if len(result4) != 3 {
		t.Errorf("Expected 3 lines, got %d", len(result4))
	}
	if result4[0] != "line1" || result4[1] != "line2" || result4[2] != "line3" {
		t.Error("SplitLines failed to trim whitespace")
	}

	// Test empty string
	result5 := SplitLines("")
	if len(result5) != 0 {
		t.Errorf("Expected 0 lines for empty string, got %d", len(result5))
	}
}

func TestParseAndFormatJson(t *testing.T) {
	// Test valid JSON
	validJSON := `{"name":"test","age":30}`
	result, err := ParseAndFormatJson(validJSON)
	if err != nil {
		t.Errorf("ParseAndFormatJson failed with valid JSON: %v", err)
	}
	if result == "" {
		t.Error("ParseAndFormatJson should return formatted JSON")
	}

	// Test invalid JSON
	invalidJSON := `{name:test}`
	result2, err2 := ParseAndFormatJson(invalidJSON)
	if err2 == nil {
		t.Error("ParseAndFormatJson should return error for invalid JSON")
	}
	if result2 != invalidJSON {
		t.Error("ParseAndFormatJson should return original string on error")
	}

	// Test with array
	arrayJSON := `[1,2,3]`
	result3, err3 := ParseAndFormatJson(arrayJSON)
	if err3 != nil {
		t.Errorf("ParseAndFormatJson failed with array JSON: %v", err3)
	}
	if result3 == "" {
		t.Error("ParseAndFormatJson should return formatted JSON for array")
	}
}
