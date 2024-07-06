package formatter

import (
	"fmt"
	"logoloco/level"
	"time"
)

// ConsoleConfig struct for console configuration.
// It includes Level for selecting detail level of output,
// PrintDate and PrintLevel options for including date and log level in output message.

type ConsoleConfig struct {
	Level      level.Level
	PrintDate  bool
	PrintLevel bool
}

// ConsoleFormatter is a simple logging formatter.
// It prints the logs to the console based on the provided configuration and Current Time.
// Config is a struct containing fields related to format and data to log.
// TimeProvider is an interface that abstracts the sequence of time of log messages.

type ConsoleFormatter struct {
	Config       ConsoleConfig
	TimeProvider TimeProvider
}

// NewConsoleFormatter is a constructor for ConsoleFormatter.
// It returns a new ConsoleFormatter with provided configuration and current time.

func NewConsoleFormatter(config ConsoleConfig) *ConsoleFormatter {

	return &ConsoleFormatter{
		Config:       config,
		TimeProvider: DefaultTimeProvider{},
	}
}

// Format function for formatting console messages based on ConsoleConfig.
// It takes as parameters a log level and a message to log.
// Each message is outputted based on the level parameter.
// It returns the formatted string and an error if any occurred during formatting process.

func (c ConsoleFormatter) Format(level level.Level, message string) (string, error) {
	if level < c.Config.Level {
		return "", nil
	}
	var lvlString string
	var dateString string

	if c.Config.PrintLevel {
		lvlString = level.String()
	}

	if c.Config.PrintDate {
		dateString = c.TimeProvider.Now().Format(time.RFC3339)
	}

	// return formatted message
	retMessage := fmt.Sprintf("[%s] [%s] [%s] %s\n", lvlString, dateString, level.String(), message)

	return retMessage, nil
}
