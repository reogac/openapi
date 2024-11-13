package models

type PsaInformation struct {
	PsaInd       PsaIndication `json:"psaInd,omitempty"`
	DnaiList     []string      `json:"dnaiList,omitempty"`
	UeIpv6Prefix string        `json:"ueIpv6Prefix,omitempty"`
	PsaUpfId     string        `json:"psaUpfId,omitempty"`
}
