// Package clock provides the Clock interface for getting the current time
package clock

import "time"

// Clock provides an interface to the current time, useful for testing
type Clock interface {
	// Now returns the current time
	Now() time.Time
}

// New returns an implementation of Clock that uses the system time
func New() Clock {
	return SysClock{}
}

// SysClock uses the system clock to return the time
type SysClock struct{}

// Now implements Clock.Now
func (cl SysClock) Now() time.Time {
	return time.Now()
}