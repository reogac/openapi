package models

type ChargingData struct {
	AfChargingIdentifier *int           `json:"afChargingIdentifier,omitempty"`
	Online               *bool          `json:"online,omitempty"`
	ServiceId            *int           `json:"serviceId,omitempty"`
	Offline              *bool          `json:"offline,omitempty"`
	SdfHandl             *bool          `json:"sdfHandl,omitempty"`
	RatingGroup          *int           `json:"ratingGroup,omitempty"`
	ReportingLevel       ReportingLevel `json:"reportingLevel,omitempty"`
	SponsorId            string         `json:"sponsorId,omitempty"`
	AppSvcProvId         string         `json:"appSvcProvId,omitempty"`
	ChgId                string         `json:"chgId"`
	MeteringMethod       MeteringMethod `json:"meteringMethod,omitempty"`
	AfChargId            string         `json:"afChargId,omitempty"`
}
