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

func TestSubstring(t *testing.T) {
	s := "Hello, World"
	if Substring(s, 0, 5) != "Hello" {
		t.Error("Substring failed")
	}
	if Substring(s, -5, -1) != "Worl" {
		t.Error("Substring with negative index failed")
	}
	if Substring(s, 10, 5) != "" {
		t.Error("Substring with invalid range failed")
	}
}

func TestReplaceAll(t *testing.T) {
	s := "hello world hello"
	result := ReplaceAll(s, "hello", "hi")
	if result != "hi world hi" {
		t.Errorf("Expected 'hi world hi', got %q", result)
	}
}

func TestReplaceN(t *testing.T) {
	s := "hello world hello"
	result := ReplaceN(s, "hello", "hi", 1)
	if result != "hi world hello" {
		t.Errorf("Expected 'hi world hello', got %q", result)
	}
}

func TestContains(t *testing.T) {
	if !Contains("hello world", "world") {
		t.Error("Contains failed")
	}
	if Contains("hello world", "xyz") {
		t.Error("Contains failed")
	}
}

func TestContainsIgnoreCase(t *testing.T) {
	if !ContainsIgnoreCase("Hello World", "hello") {
		t.Error("ContainsIgnoreCase failed")
	}
	if !ContainsIgnoreCase("Hello World", "WORLD") {
		t.Error("ContainsIgnoreCase failed")
	}
}

func TestStartsWith(t *testing.T) {
	if !StartsWith("hello world", "hello") {
		t.Error("StartsWith failed")
	}
	if StartsWith("hello world", "world") {
		t.Error("StartsWith failed")
	}
}

func TestEndsWith(t *testing.T) {
	if !EndsWith("hello world", "world") {
		t.Error("EndsWith failed")
	}
	if EndsWith("hello world", "hello") {
		t.Error("EndsWith failed")
	}
}

func TestRepeat(t *testing.T) {
	result := Repeat("ab", 3)
	if result != "ababab" {
		t.Errorf("Expected 'ababab', got %q", result)
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

func TestTruncateWithEllipsis(t *testing.T) {
	result := TruncateWithEllipsis("hello world", 8)
	if result != "hello..." {
		t.Errorf("Expected 'hello...', got %q", result)
	}
	if TruncateWithEllipsis("hi", 5) != "hi" {
		t.Error("TruncateWithEllipsis failed")
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
