package datautils

import "strings"

// IsWhite TODO
func IsWhite(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

// IsNotWhite TODO
func IsNotWhite(s string) bool {
	return !IsWhite(s)
}
