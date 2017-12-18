package datautils

import "strings"

func IsWhite(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func IsNotWhite(s string) bool {
	return !IsWhite(s)
}
