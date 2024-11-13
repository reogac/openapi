package models

type CircumstanceDescription struct {
	Tm      string           `json:"tm,omitempty"`
	LocArea *NetworkAreaInfo `json:"locArea,omitempty"`
	Vol     *int64           `json:"vol,omitempty"`
	Freq    *float64         `json:"freq,omitempty"`
}
