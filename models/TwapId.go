package models

type TwapId struct {
	BssId        string `json:"bssId,omitempty"`
	CivicAddress string `json:"civicAddress,omitempty"`
	SsId         string `json:"ssId"`
}
