package logoloco

import (
	"github.com/alonohana627/logoloco/formatter"
	"github.com/alonohana627/logoloco/level"
	"github.com/alonohana627/logoloco/sink"
	"os"
)

// ConsoleLogger struct contains Sink and Formatter which are needed for it to perform its operations.
type ConsoleLogger struct {
	Sink      *sink.ConsoleSink           // Sink which provides methods for writing logs.
	Formatter *formatter.ConsoleFormatter // Formatter which formats the logs.
}

// NewConsoleLogger is a function which returns a new ConsoleLogger provided with a sink and a formatter.
// It uses the provided sink and formatter to create a ConsoleLogger.
func NewConsoleLogger(sink *sink.ConsoleSink, formatter *formatter.ConsoleFormatter) *ConsoleLogger {
	return &ConsoleLogger{
		Sink:      sink,
		Formatter: formatter,
	}
}

// DefaultConsoleLogger is a function which returns a ConsoleLogger with default parameters.
// It uses a default Sink and formatter which has print level and print date enabled, with level set to 0.
func DefaultConsoleLogger() *ConsoleLogger {
	consoleSink := sink.NewConsoleSink()                                       // Creates a new default ConsoleSink.
	consoleFormatter := formatter.NewConsoleFormatter(formatter.ConsoleConfig{ // Uses a default configuration for ConsoleFormatter.
		Level:      0,
		PrintDate:  true,
		PrintLevel: true,
	})

	return &ConsoleLogger{
		Sink:      consoleSink,
		Formatter: consoleFormatter,
	}
}

// logMessage is a helper function that passes log messages from different log levels to the formatter and then to the sink.
// It takes log level and a log message as parameters.
func (c *ConsoleLogger) logMessage(level level.Level, message string) error {
	messageToSend, err := c.Formatter.Format(level, message) // Formats the message using the given level.
	if err != nil {
		return err
	}

	err = c.Sink.Write(messageToSend) // Send the formatted message to the Sink.
	return err

}

// Debug is a function which sends debug level message to the log.
func (c *ConsoleLogger) Debug(message string) error {
	return c.logMessage(level.DEBUG, message) // Calls logMessage with DEBUG level.
}

// Info is a function which sends information level message to the log.
func (c *ConsoleLogger) Info(message string) error {
	return c.logMessage(level.INFO, message) // Calls logMessage with INFO level.
}

// Warn is a function which sends warning level message to the log.
func (c *ConsoleLogger) Warn(message string) error {
	return c.logMessage(level.WARNING, message) // Calls logMessage with WARNING level.
}

// Error is a function which sends error level message to the log.
func (c *ConsoleLogger) Error(message string) error {
	return c.logMessage(level.ERROR, message) // Calls logMessage with ERROR level.
}

// Fatal is a function which sends fatal error level message to the log and it exits the program with status 1.
func (c *ConsoleLogger) Fatal(message string) error {
	err := c.logMessage(level.INFO, message) // Calls logMessage with INFO level.
	os.Exit(1)
	return err
}
