package main

import (
	`time`
)

// Should time use value or pointer semantics? If you need to modify a time
// value should you mutate the value or create a new one?
type Time struct {
	sec  int64
	nsec int32
	loc  *Location
}

// Factory functions dictate the semantics that will be used
func Now() Time {
	// implementation
}

// Add is using a value receiver and returning a value of type Time. Value semantics
func (t Time) Add(d time.Duration) Time {
	// implementation
}
