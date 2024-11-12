package models

type SmcceInfo struct {
	SmcceUeList SmcceUeList `json:"smcceUeList"`
	Dnn         string      `json:"dnn,omitempty"`
	Snssai      *Snssai     `json:"snssai,omitempty"`
}
