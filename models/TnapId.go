package models

type TnapId struct {
	CivicAddress string `json:"civicAddress,omitempty"`
	SsId         string `json:"ssId,omitempty"`
	BssId        string `json:"bssId,omitempty"`
}
