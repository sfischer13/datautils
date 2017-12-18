package datautils

import "time"

// CountTrue count the number of true values in a slice.
func CountTrue(s []bool) int {
	n := 0
	for _, b := range s {
		if b {
			n++
		}
	}
	return n
}

// ParseRFCDate parses a date string in RFC3339 format.
// On failure, the current time will be returned.
func ParseRFCDate(timeString string) time.Time {
	date, err := time.Parse(time.RFC3339, timeString)
	if err != nil {
		return time.Now()
	}
	return date
}
