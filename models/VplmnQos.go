package models

type VplmnQos struct {
	Arp         *Arp   `json:"arp,omitempty"`
	SessionAmbr *Ambr  `json:"sessionAmbr,omitempty"`
	MaxFbrDl    string `json:"maxFbrDl,omitempty"`
	MaxFbrUl    string `json:"maxFbrUl,omitempty"`
	GuaFbrDl    string `json:"guaFbrDl,omitempty"`
	GuaFbrUl    string `json:"guaFbrUl,omitempty"`
	FiveQi      *int   `json:"5qi,omitempty"`
}
