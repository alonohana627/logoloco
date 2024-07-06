package sink

import (
	"fmt"
	"io"
	"os"
	"sync"
)

// ConsoleSink writes messages to the console. It uses a sync.Mutex to ensure
// safe concurrent use, and a function to write to an io.Writer. This can be
// replaced by a mock in tests.

type ConsoleSink struct {
	mu       sync.Mutex                                             // Helps in achieving safe concurrent use.
	fprintln func(w io.Writer, a ...interface{}) (n int, err error) // Function to write to an io.Writer.
}

// NewConsoleSink constructs a new ConsoleSink instance with the default fmt.Fprintln
// function for writing to the console output (os.Stdout).
func NewConsoleSink() *ConsoleSink {
	return &ConsoleSink{
		fprintln: fmt.Fprintln,
	}
}

// Write takes a message string as input and writes it to the console output.
// It returns an error in case it's not able to write to the console output.
func (c *ConsoleSink) Write(message string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, err := c.fprintln(os.Stdout, message)
	return err
}
