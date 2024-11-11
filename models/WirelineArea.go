package models

type WirelineArea struct {
	AreaCodeC     string   `json:"areaCodeC,omitempty"`
	GlobalLineIds []string `json:"globalLineIds,omitempty"`
	HfcNIds       []string `json:"hfcNIds,omitempty"`
	AreaCodeB     string   `json:"areaCodeB,omitempty"`
}
