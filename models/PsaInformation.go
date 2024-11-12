package models

type PsaInformation struct {
	DnaiList     []string      `json:"dnaiList,omitempty"`
	UeIpv6Prefix string        `json:"ueIpv6Prefix,omitempty"`
	PsaUpfId     string        `json:"psaUpfId,omitempty"`
	PsaInd       PsaIndication `json:"psaInd,omitempty"`
}
