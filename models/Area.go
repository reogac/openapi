package models

type Area struct {
	AreaCode string   `json:"areaCode,omitempty"`
	Tacs     []string `json:"tacs,omitempty"`
}
