package formatter

import (
	"github.com/alonohana627/logoloco/level"
)

// LogFormatter is an interface specifying a method
// for formatting log messages.

type LogFormatter interface {
	Format(level level.Level, message string) (string, error)
}
