package models

type PsaInformation struct {
	UeIpv6Prefix string        `json:"ueIpv6Prefix,omitempty"`
	PsaUpfId     string        `json:"psaUpfId,omitempty"`
	PsaInd       PsaIndication `json:"psaInd,omitempty"`
	DnaiList     []string      `json:"dnaiList,omitempty"`
}
