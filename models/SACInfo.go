package models

type SACInfo struct {
	NumericValNumUes     *int `json:"numericValNumUes,omitempty"`
	NumericValNumPduSess *int `json:"numericValNumPduSess,omitempty"`
	PercValueNumUes      *int `json:"percValueNumUes,omitempty"`
	PercValueNumPduSess  *int `json:"percValueNumPduSess,omitempty"`
}
