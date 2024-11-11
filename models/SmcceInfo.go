package models

type SmcceInfo struct {
	Dnn         string      `json:"dnn,omitempty"`
	Snssai      *Snssai     `json:"snssai,omitempty"`
	SmcceUeList SmcceUeList `json:"smcceUeList"`
}
