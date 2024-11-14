package models
type UeSliceMbr struct {
	 SliceMbr	map[string]SliceMbr	`json:"sliceMbr"`
	 ServingSnssai	Snssai	`json:"servingSnssai"`
	 MappedHomeSnssai	*Snssai	`json:"mappedHomeSnssai,omitempty"`
}
