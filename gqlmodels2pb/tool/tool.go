package tool

import "strings"

// 获取字符串的pb格式的字段名
func Ns(s string) string {
	lth := len(s)
	i := lth - 1
	t := false
	for ; i > 1; i-- {
		c := s[i]
		if !t && c >= 'A' && c <= 'Z' {
			t = true
		}
		if t && 'a' <= c && c <= 'z' {
			s = s[:i+1] + "_" + s[i+1:]
			t = false
		}
	}
	ind := strings.Index(s, "ID")
	lth = len(s)
	if ind > 0 && ind+2 < lth && s[ind+2] >= 'A' && s[ind+2] <= 'Z' {
		s = s[:ind+2] + "_" + s[ind+2:]
	}
	return strings.ToLower(s)
}
