package formatter

import "time"

type TimeProvider interface {
	Now() time.Time
}

type DefaultTimeProvider struct{}

func (p DefaultTimeProvider) Now() time.Time {
	return time.Now()
}
