package pointerutil

import (
	"os"
	"time"
)

// StringPtr returns a pointer to a string value
func StringPtr(s string) *string {
	return &s
}

// BoolPtr returns a pointer to a boolean value
func BoolPtr(b bool) *bool {
	return &b
}

// TimeDurationPtr returns a pointer to a time duration value
func TimeDurationPtr(duration string) *time.Duration {
	d, _ := time.ParseDuration(duration)

	return &d
}

// FileModePtr returns a pointer to the given os.FileMode
func FileModePtr(o os.FileMode) *os.FileMode {
	return &o
}
