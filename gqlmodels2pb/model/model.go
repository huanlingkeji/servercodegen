package model

type MyObject struct {
	Name     string
	FieldArr []*MyField
	FieldMap map[string]*MyField
}

type MyField struct {
	Name       string
	Type       string
	IsBaseType bool
	IsArr      bool
	IsSelector bool
	IsStar     bool
	IsMap      bool
	MKeyType   string
	MValueType string
}
