/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:14 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type GbrQosFlowInformation struct {
	GuaFbrDl                  string                  `json:"guaFbrDl"`
	GuaFbrUl                  string                  `json:"guaFbrUl"`
	NotifControl              NotificationControl     `json:"notifControl,omitempty"`
	MaxPacketLossRateDl       *int                    `json:"maxPacketLossRateDl,omitempty"`
	MaxPacketLossRateUl       *int                    `json:"maxPacketLossRateUl,omitempty"`
	AlternativeQosProfileList []AlternativeQosProfile `json:"alternativeQosProfileList,omitempty"`
	MaxFbrDl                  string                  `json:"maxFbrDl"`
	MaxFbrUl                  string                  `json:"maxFbrUl"`
}