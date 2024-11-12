package models

type ConsumerNfInformation struct {
	NfSetId string `json:"nfSetId,omitempty"`
	TaiList []Tai  `json:"taiList,omitempty"`
	NfId    string `json:"nfId,omitempty"`
}
