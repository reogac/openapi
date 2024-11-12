package models

type ChargingData struct {
	ChgId                string         `json:"chgId"`
	MeteringMethod       MeteringMethod `json:"meteringMethod,omitempty"`
	ServiceId            *int           `json:"serviceId,omitempty"`
	AfChargingIdentifier *int           `json:"afChargingIdentifier,omitempty"`
	SponsorId            string         `json:"sponsorId,omitempty"`
	AppSvcProvId         string         `json:"appSvcProvId,omitempty"`
	AfChargId            string         `json:"afChargId,omitempty"`
	Offline              *bool          `json:"offline,omitempty"`
	Online               *bool          `json:"online,omitempty"`
	SdfHandl             *bool          `json:"sdfHandl,omitempty"`
	RatingGroup          *int           `json:"ratingGroup,omitempty"`
	ReportingLevel       ReportingLevel `json:"reportingLevel,omitempty"`
}
