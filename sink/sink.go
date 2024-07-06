package sink

// Sink represents a destination to write messages to. A Sink could represent
// various output streams like file, network etc.
// The Write method in the Sink interface takes a string message as input
// and returns an error if the write operation is unsuccessful.
type Sink interface {
	Write(message string) error
}
