package models

type LocationInfo struct {
	Confidence *int         `json:"confidence,omitempty"`
	Loc        UserLocation `json:"loc"`
	Ratio      *int         `json:"ratio,omitempty"`
}
