package gen

import "genserver/gqlmodels2pb/other"

type BannerShowPolicy string

type BannerInput struct {
	ID           *int                      `json:"id"`
	IDArr        []*int                    `json:"id"`
	IDMap        map[string]*int           `json:"id"`
	IDMapArr     map[string][]*int         `json:"id"` // TODO 暂时不支持复杂 结构  map的嵌套  map和slice的嵌套
	Policy       *BannerShowPolicy         `json:"showPolicy"`
	PolicyArr    []*BannerShowPolicy       `json:"showPolicy"`
	PolicyMap    map[int]*BannerShowPolicy `json:"showPolicy"`
	Person       other.A
	Persons      []other.A
	PersonsPtr   []*other.A
	PersonMap    map[int]*other.A
	PersonMapInt map[int]interface{}
}
