package tmlGenItem

// 指定tmpl标记的生成
type ITmplGen interface {
	GenCode() string
}

var TmplGenMap = map[string]ITmplGen{}