package lib

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

func TestLogCfg_ParseLevel(t *testing.T) {
	cfg := &LogCfg{}
	
	// Test default level (empty string)
	level := cfg.ParseLevel()
	if level != logrus.InfoLevel {
		t.Errorf("Expected InfoLevel for empty level, got %v", level)
	}
	
	// Test valid levels
	cfg.Level = "debug"
	level = cfg.ParseLevel()
	if level != logrus.DebugLevel {
		t.Errorf("Expected DebugLevel, got %v", level)
	}
	
	cfg.Level = "info"
	level = cfg.ParseLevel()
	if level != logrus.InfoLevel {
		t.Errorf("Expected InfoLevel, got %v", level)
	}
	
	cfg.Level = "warn"
	level = cfg.ParseLevel()
	if level != logrus.WarnLevel {
		t.Errorf("Expected WarnLevel, got %v", level)
	}
	
	cfg.Level = "error"
	level = cfg.ParseLevel()
	if level != logrus.ErrorLevel {
		t.Errorf("Expected ErrorLevel, got %v", level)
	}
	
	// Test invalid level
	cfg.Level = "invalid"
	level = cfg.ParseLevel()
	if level != logrus.InfoLevel {
		t.Errorf("Expected InfoLevel for invalid level, got %v", level)
	}
}

func TestLogCfg_BuildLogger(t *testing.T) {
	cfg := &LogCfg{
		Level:    "info",
		KeepDays: 1,
	}
	
	logger := cfg.BuildLogger()
	if logger == nil {
		t.Error("BuildLogger should return a logger")
	}
	if logger.Level != logrus.InfoLevel {
		t.Errorf("Expected InfoLevel, got %v", logger.Level)
	}
}

func TestLogCfg_BuildLoggerWithPath(t *testing.T) {
	// Create temporary directory for test logs
	tmpDir := os.TempDir()
	logPath := filepath.Join(tmpDir, "test_log")
	
	cfg := &LogCfg{
		Path:     logPath,
		Level:    "debug",
		KeepDays: 1,
	}
	
	logger := cfg.BuildLogger()
	if logger == nil {
		t.Error("BuildLogger should return a logger")
	}
	
	// Test logging
	logger.Info("Test log message")
	
	// Cleanup
	os.Remove(logPath)
}

func TestLogCfg_BuildLogWriter(t *testing.T) {
	tmpDir := os.TempDir()
	logPath := filepath.Join(tmpDir, "test_writer_log")
	
	cfg := &LogCfg{
		Path:     logPath,
		KeepDays: 1,
	}
	
	writer := cfg.BuildLogWriter()
	if writer == nil {
		t.Error("BuildLogWriter should return a writer")
	}
	
	// Test default KeepDays
	cfg2 := &LogCfg{
		Path: logPath,
	}
	writer2 := cfg2.BuildLogWriter()
	if writer2 == nil {
		t.Error("BuildLogWriter should return a writer with default KeepDays")
	}
	
	// Cleanup
	os.Remove(logPath)
}

func TestGetFormatter(t *testing.T) {
	formatter := GetFormatter()
	if formatter == nil {
		t.Error("GetFormatter should return a formatter")
	}
	
	// Test formatter with entry
	logger := logrus.New()
	logger.SetFormatter(formatter)
	logger.SetReportCaller(true)
	
	var buf bytes.Buffer
	logger.SetOutput(&buf)
	logger.Info("Test message")
	
	output := buf.String()
	if !strings.Contains(output, "INFO") {
		t.Error("Formatter output should contain log level")
	}
	if !strings.Contains(output, "Test message") {
		t.Error("Formatter output should contain log message")
	}
}

func TestLogCfg_InitLogrus(t *testing.T) {
	cfg := &LogCfg{
		Level:    "warn",
		KeepDays: 1,
	}
	
	cfg.InitLogrus()
	
	// Verify logrus is configured
	if logrus.GetLevel() != logrus.WarnLevel {
		t.Errorf("Expected WarnLevel, got %v", logrus.GetLevel())
	}
}

func TestLogCfg_InitLogrusWithPath(t *testing.T) {
	tmpDir := os.TempDir()
	logPath := filepath.Join(tmpDir, "test_init_log")
	
	cfg := &LogCfg{
		Path:     logPath,
		Level:    "error",
		KeepDays: 1,
	}
	
	cfg.InitLogrus()
	
	if logrus.GetLevel() != logrus.ErrorLevel {
		t.Errorf("Expected ErrorLevel, got %v", logrus.GetLevel())
	}
	
	// Test logging
	logrus.Info("This should not appear")
	logrus.Error("This should appear")
	
	// Cleanup
	os.Remove(logPath)
}

func TestFormatter_Format(t *testing.T) {
	formatter := &formatter{}
	
	logger := logrus.New()
	logger.SetFormatter(formatter)
	logger.SetReportCaller(true)
	
	var buf bytes.Buffer
	logger.SetOutput(&buf)
	
	logger.WithFields(logrus.Fields{
		"key": "value",
	}).Info("Test message with fields")
	
	output := buf.String()
	if !strings.Contains(output, "INFO") {
		t.Error("Formatter output should contain log level")
	}
	if !strings.Contains(output, "Test message with fields") {
		t.Error("Formatter output should contain log message")
	}
}

func TestLogCfg_KeepDays(t *testing.T) {
	cfg := &LogCfg{
		Path:     "/tmp/test",
		KeepDays: 7 * 24 * time.Hour,
	}
	
	writer := cfg.BuildLogWriter()
	if writer == nil {
		t.Error("BuildLogWriter should return a writer")
	}
}
