package models

type SACInfo struct {
	PercValueNumPduSess  *int `json:"percValueNumPduSess,omitempty"`
	NumericValNumUes     *int `json:"numericValNumUes,omitempty"`
	NumericValNumPduSess *int `json:"numericValNumPduSess,omitempty"`
	PercValueNumUes      *int `json:"percValueNumUes,omitempty"`
}
