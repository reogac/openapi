package models

type UeSliceMbr struct {
	ServingSnssai    Snssai              `json:"servingSnssai"`
	MappedHomeSnssai *Snssai             `json:"mappedHomeSnssai,omitempty"`
	SliceMbr         map[string]SliceMbr `json:"sliceMbr"`
}
