package logoloco

// Logger interface
type Logger interface {
	Debug(message string) error
	Info(message string) error
	Warn(message string) error
	Error(message string) error
	Fatal(message string)
}
