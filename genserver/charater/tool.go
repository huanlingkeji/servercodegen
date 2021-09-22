package charater

import "strings"

// LowerFirstChar LowerFirstChar
func LowerFirstChar(s string) string {
	if s == "ID" {
		return "id"
	}
	return strings.ToLower(string(s[0])) + s[1:]
}

// UpperAllChar UpperAllChar
func UpperAllChar(s string) string {
	return strings.ToUpper(s)
}

// UpperFirstChar UpperFirstChar
func UpperFirstChar(s string) string {
	if s == "ip" {
		return "IP"
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}
