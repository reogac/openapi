/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:41 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ChargingData struct {
	ReportingLevel       ReportingLevel `json:"reportingLevel,omitempty"`
	AppSvcProvId         string         `json:"appSvcProvId,omitempty"`
	AfChargingIdentifier *int           `json:"afChargingIdentifier,omitempty"`
	MeteringMethod       MeteringMethod `json:"meteringMethod,omitempty"`
	Offline              *bool          `json:"offline,omitempty"`
	SdfHandl             *bool          `json:"sdfHandl,omitempty"`
	RatingGroup          *int           `json:"ratingGroup,omitempty"`
	ServiceId            *int           `json:"serviceId,omitempty"`
	SponsorId            string         `json:"sponsorId,omitempty"`
	AfChargId            string         `json:"afChargId,omitempty"`
	ChgId                string         `json:"chgId"`
	Online               *bool          `json:"online,omitempty"`
}
