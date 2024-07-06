package sink

import (
	"os"
	"sync"
)

// FileSink is a structure that represents a sink to a file.
// It contains the mutex, filename and a file pointer.
type FileSink struct {
	mu       sync.Mutex
	filename string
	file     *os.File
}

// NewFileSink is a function that creates a new FileSink object.
// The function tries to open the file with the given filename in append mode,
// and returns the created FileSink object if successful.
// In case of any error while opening the file, the error is returned.

func NewFileSink(filename string) (*FileSink, error) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	return &FileSink{
		filename: filename,
		file:     file,
	}, nil
}

// Write is a method on the FileSink type.
// It writes the given message plus a newline to the file of the sink.
// In case of any error while writing, the error is returned.
func (fs *FileSink) Write(message string) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	if _, err := fs.file.WriteString(message + "\n"); err != nil {
		return err
	}
	return nil
}
