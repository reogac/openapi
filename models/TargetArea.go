package models
type TargetArea struct {
	 TaList	[]Tai	`json:"taList,omitempty"`
	 TaiRangeList	[]TaiRange	`json:"taiRangeList,omitempty"`
	 AnyTa	*bool	`json:"anyTa,omitempty"`
}
