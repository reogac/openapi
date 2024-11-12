package models

type WirelineArea struct {
	GlobalLineIds []string `json:"globalLineIds,omitempty"`
	HfcNIds       []string `json:"hfcNIds,omitempty"`
	AreaCodeB     string   `json:"areaCodeB,omitempty"`
	AreaCodeC     string   `json:"areaCodeC,omitempty"`
}
