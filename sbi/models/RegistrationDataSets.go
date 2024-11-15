/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:11:55 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RegistrationDataSets struct {
	SmfRegistration   *SmfRegistrationInfo          `json:"smfRegistration,omitempty"`
	Smsf3Gpp          *SmsfRegistration             `json:"smsf3Gpp,omitempty"`
	SmsfNon3Gpp       *SmsfRegistration             `json:"smsfNon3Gpp,omitempty"`
	IpSmGw            *IpSmGwRegistration           `json:"ipSmGw,omitempty"`
	NwdafRegistration *NwdafRegistrationInfo        `json:"nwdafRegistration,omitempty"`
	Amf3Gpp           *Amf3GppAccessRegistration    `json:"amf3Gpp,omitempty"`
	AmfNon3Gpp        *AmfNon3GppAccessRegistration `json:"amfNon3Gpp,omitempty"`
}
