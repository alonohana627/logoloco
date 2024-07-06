package sink

import (
	"bytes"
	"errors"
	"io"
	"os"
	"sync"
	"testing"
)

// TestConsoleSink_Write tests the case where the Write method of the ConsoleSink instance
// is supposed to write a message to console without any errors.

func TestConsoleSink_Write(t *testing.T) {
	consoleSink := NewConsoleSink()

	expected := "test message\n"

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Test error where `Write` operation fails.
	err := consoleSink.Write("test message")
	if err != nil {
		t.Errorf("Unexpected error in write operation: %v", err)
	}

	err = w.Close()
	if err != nil {
		t.Errorf("Unexpected error in write operation: %v", err)
		return
	}
	os.Stdout = old

	var buf bytes.Buffer

	_, err = io.Copy(&buf, r)
	if err != nil {
		t.Errorf("Unexpected error in write operation: %v", err)
		return
	}

	if buf.String() != expected {
		t.Errorf("Unexpected output Write(%q) == %q, want %q", "test message", buf.String(), expected)
	}
}

// TestConsoleSink_WriteError tests the case where the write operation to the console is supposed to fail
func TestConsoleSink_WriteError(t *testing.T) {
	sink := &ConsoleSink{
		mu: sync.Mutex{},
		fprintln: func(w io.Writer, a ...interface{}) (n int, err error) {
			return 0, errors.New("simulated write error") // Simulate an error in writing to the io.Writer
		},
	}

	// Test error where `Write` operation successfully fails.
	err := sink.Write("test message")
	if err == nil || err.Error() != "simulated write error" {
		t.Error("Expected Write to fail with 'simulated write error', but it didn't")
	}
}
