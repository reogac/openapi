package models

type GbrQosFlowInformation struct {
	GuaFbrUl                  string                  `json:"guaFbrUl"`
	NotifControl              string                  `json:"notifControl,omitempty"`
	MaxPacketLossRateDl       *int                    `json:"maxPacketLossRateDl,omitempty"`
	MaxPacketLossRateUl       *int                    `json:"maxPacketLossRateUl,omitempty"`
	AlternativeQosProfileList []AlternativeQosProfile `json:"alternativeQosProfileList,omitempty"`
	MaxFbrDl                  string                  `json:"maxFbrDl"`
	MaxFbrUl                  string                  `json:"maxFbrUl"`
	GuaFbrDl                  string                  `json:"guaFbrDl"`
}