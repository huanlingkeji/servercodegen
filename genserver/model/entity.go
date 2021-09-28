package model

// MyEntity 实体
type MyEntity struct {
	ModelName string     // 实体名字
	Fields    []*MyField // 实体设计的字段
	ModelZH   string     // 备注
}
