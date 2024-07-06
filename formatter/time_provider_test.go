package formatter_test

import (
	"github.com/alonohana627/logoloco/formatter"
	"testing"
	"time"
)

type mockTimeProvider struct{}

func (m mockTimeProvider) Now() time.Time {
	t, _ := time.Parse("2006-01-02", "2024-05-04")
	return t
}

func TestDefaultProvider(t *testing.T) {
	defaultTimeProvider := formatter.DefaultTimeProvider{}
	defaultTime := defaultTimeProvider.Now()

	if defaultTime.After(time.Now()) {
		t.Fatalf("Default time provider returned a future date")
	}
}

func TestMockTimeProvider(t *testing.T) {
	mockTimeProvider := mockTimeProvider{}
	mockTime := mockTimeProvider.Now()

	expectedTime, _ := time.Parse("2006-01-02", "2024-05-04")

	if !mockTime.Equal(expectedTime) {
		t.Fatalf("Expected the mock provider to return the mock date, but it did not")
	}
}
