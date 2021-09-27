package charater

import "strings"

func LowerFirstChar(s string) string {
	if s == "ID" {
		return "id"
	}
	return strings.ToLower(string(s[0])) + s[1:]
}

func UpperFirstChar(s string) string {
	if s == "ip" {
		return "IP"
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}
