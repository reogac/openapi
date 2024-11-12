package models

type TwapId struct {
	SsId         string `json:"ssId"`
	BssId        string `json:"bssId,omitempty"`
	CivicAddress string `json:"civicAddress,omitempty"`
}
