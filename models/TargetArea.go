package models
type TargetArea struct {
	 AnyTa	*bool	`json:"anyTa,omitempty"`
	 TaList	[]Tai	`json:"taList,omitempty"`
	 TaiRangeList	[]TaiRange	`json:"taiRangeList,omitempty"`
}
