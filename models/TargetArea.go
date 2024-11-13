package models

type TargetArea struct {
	TaiRangeList []TaiRange `json:"taiRangeList,omitempty"`
	AnyTa        *bool      `json:"anyTa,omitempty"`
	TaList       []Tai      `json:"taList,omitempty"`
}
