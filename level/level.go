package level

type Level int

const (
	DEBUG   = Level(0)
	INFO    = Level(1)
	WARNING = Level(2)
	ERROR   = Level(3)
)

func (l Level) String() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	}

	panic("unreachable")
}
