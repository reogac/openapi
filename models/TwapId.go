package models

type TwapId struct {
	CivicAddress string `json:"civicAddress,omitempty"`
	SsId         string `json:"ssId"`
	BssId        string `json:"bssId,omitempty"`
}
