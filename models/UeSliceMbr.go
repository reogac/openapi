package models

type UeSliceMbr struct {
	MappedHomeSnssai *Snssai             `json:"mappedHomeSnssai,omitempty"`
	SliceMbr         map[string]SliceMbr `json:"sliceMbr"`
	ServingSnssai    Snssai              `json:"servingSnssai"`
}
