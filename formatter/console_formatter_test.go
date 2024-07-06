package formatter_test

import (
	f "logoloco/formatter"
	"logoloco/level"
	"strings"
	"testing"
	"time"
)

type mockTimeProvider2 struct{}

func (m mockTimeProvider2) Now() time.Time {
	t, _ := time.Parse(time.RFC3339, "2024-05-04T15:04:05Z")
	return t
}

func TestNewConsoleFormatter(t *testing.T) {
	config := f.ConsoleConfig{
		Level:      level.INFO,
		PrintDate:  false,
		PrintLevel: false,
	}

	formatter := f.NewConsoleFormatter(config)

	if formatter == nil {
		t.Fatal("Expected formatter to be not nil")
	}
}

func TestFormat(t *testing.T) {
	config := f.ConsoleConfig{
		Level:      level.INFO,
		PrintDate:  true,
		PrintLevel: true,
	}

	mockTime := mockTimeProvider2{}
	formatter := &f.ConsoleFormatter{
		Config:       config,
		TimeProvider: mockTime,
	}
	message, _ := formatter.Format(level.INFO, "test message")

	if !strings.Contains(message, level.INFO.String()) {
		t.Fatalf("Expected message to contain level, but it did not")
	}

	if !strings.Contains(message, mockTime.Now().Format(time.RFC3339)) {
		t.Fatalf("Expected message to contain date, but it did not")
	}

	//Test case where level is less than config Level
	message, _ = formatter.Format(level.DEBUG, "test message")

	if message != "" {
		t.Fatalf("Expected empty message as log level is below config level")
	}
}
