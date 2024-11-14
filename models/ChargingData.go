/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:22:57 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ChargingData struct {
	ReportingLevel       ReportingLevel `json:"reportingLevel,omitempty"`
	ServiceId            *int           `json:"serviceId,omitempty"`
	AfChargingIdentifier *int           `json:"afChargingIdentifier,omitempty"`
	AfChargId            string         `json:"afChargId,omitempty"`
	MeteringMethod       MeteringMethod `json:"meteringMethod,omitempty"`
	Offline              *bool          `json:"offline,omitempty"`
	SdfHandl             *bool          `json:"sdfHandl,omitempty"`
	SponsorId            string         `json:"sponsorId,omitempty"`
	AppSvcProvId         string         `json:"appSvcProvId,omitempty"`
	ChgId                string         `json:"chgId"`
	Online               *bool          `json:"online,omitempty"`
	RatingGroup          *int           `json:"ratingGroup,omitempty"`
}
