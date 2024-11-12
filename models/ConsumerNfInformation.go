package models

type ConsumerNfInformation struct {
	NfId    string `json:"nfId,omitempty"`
	NfSetId string `json:"nfSetId,omitempty"`
	TaiList []Tai  `json:"taiList,omitempty"`
}
