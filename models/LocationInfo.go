package models

type LocationInfo struct {
	Ratio      *int         `json:"ratio,omitempty"`
	Confidence *int         `json:"confidence,omitempty"`
	Loc        UserLocation `json:"loc"`
}
