package time

import (
	"testing"
	"time"
)

func TestFormatDate(t *testing.T) {
	tm := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)
	result := FormatDate(tm)
	if result != "2024-01-15" {
		t.Errorf("Expected '2024-01-15', got %q", result)
	}
}

func TestFormatTime(t *testing.T) {
	tm := time.Date(2024, 1, 15, 10, 30, 45, 0, time.UTC)
	result := FormatTime(tm)
	if result != "10:30:45" {
		t.Errorf("Expected '10:30:45', got %q", result)
	}
}

func TestFormatDateTime(t *testing.T) {
	tm := time.Date(2024, 1, 15, 10, 30, 45, 0, time.UTC)
	result := FormatDateTime(tm)
	if result != "2024-01-15 10:30:45" {
		t.Errorf("Expected '2024-01-15 10:30:45', got %q", result)
	}
}

func TestParseDate(t *testing.T) {
	tm, err := ParseDate("2024-01-15")
	if err != nil {
		t.Fatalf("ParseDate failed: %v", err)
	}
	if tm.Year() != 2024 || tm.Month() != 1 || tm.Day() != 15 {
		t.Error("ParseDate failed")
	}
}

func TestParseDateTime(t *testing.T) {
	tm, err := ParseDateTime("2024-01-15 10:30:45")
	if err != nil {
		t.Fatalf("ParseDateTime failed: %v", err)
	}
	if tm.Hour() != 10 || tm.Minute() != 30 {
		t.Error("ParseDateTime failed")
	}
}

func TestAddDays(t *testing.T) {
	tm := time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)
	result := AddDays(tm, 5)
	if result.Day() != 20 {
		t.Errorf("Expected day 20, got %d", result.Day())
	}
}

func TestDiffDays(t *testing.T) {
	t1 := time.Date(2024, 1, 20, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)
	diff := DiffDays(t1, t2)
	if diff != 5 {
		t.Errorf("Expected 5 days, got %d", diff)
	}
}

func TestIsToday(t *testing.T) {
	now := time.Now()
	if !IsToday(now) {
		t.Error("IsToday failed")
	}
	tomorrow := AddDays(now, 1)
	if IsToday(tomorrow) {
		t.Error("IsToday failed")
	}
}

func TestStartOfDay(t *testing.T) {
	tm := time.Date(2024, 1, 15, 10, 30, 45, 0, time.UTC)
	result := StartOfDay(tm)
	if result.Hour() != 0 || result.Minute() != 0 || result.Second() != 0 {
		t.Error("StartOfDay failed")
	}
}

func TestStartOfWeek(t *testing.T) {
	tm := time.Date(2024, 1, 15, 10, 30, 45, 0, time.UTC)
	result := StartOfWeek(tm)
	if result.Weekday() != time.Monday {
		t.Error("StartOfWeek failed")
	}
}

func TestStartOfMonth(t *testing.T) {
	tm := time.Date(2024, 1, 15, 10, 30, 45, 0, time.UTC)
	result := StartOfMonth(tm)
	if result.Day() != 1 {
		t.Error("StartOfMonth failed")
	}
}
