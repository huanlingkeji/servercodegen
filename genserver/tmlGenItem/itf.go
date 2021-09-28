package tmlGenItem

// ITmplGen 指定tmpl标记的生成
type ITmplGen interface {
	GenCode() string
}

// TmplGenMap TmplGenMap
var TmplGenMap = map[string]ITmplGen{}
