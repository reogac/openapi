/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:11:57 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ChargingData struct {
	AfChargingIdentifier *int           `json:"afChargingIdentifier,omitempty"`
	ChgId                string         `json:"chgId"`
	MeteringMethod       MeteringMethod `json:"meteringMethod,omitempty"`
	Offline              *bool          `json:"offline,omitempty"`
	Online               *bool          `json:"online,omitempty"`
	SdfHandl             *bool          `json:"sdfHandl,omitempty"`
	RatingGroup          *int           `json:"ratingGroup,omitempty"`
	AppSvcProvId         string         `json:"appSvcProvId,omitempty"`
	AfChargId            string         `json:"afChargId,omitempty"`
	ReportingLevel       ReportingLevel `json:"reportingLevel,omitempty"`
	ServiceId            *int           `json:"serviceId,omitempty"`
	SponsorId            string         `json:"sponsorId,omitempty"`
}
