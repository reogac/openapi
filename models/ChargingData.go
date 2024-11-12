package models

type ChargingData struct {
	MeteringMethod       MeteringMethod `json:"meteringMethod,omitempty"`
	Online               *bool          `json:"online,omitempty"`
	RatingGroup          *int           `json:"ratingGroup,omitempty"`
	ReportingLevel       ReportingLevel `json:"reportingLevel,omitempty"`
	SponsorId            string         `json:"sponsorId,omitempty"`
	AfChargingIdentifier *int           `json:"afChargingIdentifier,omitempty"`
	AfChargId            string         `json:"afChargId,omitempty"`
	ChgId                string         `json:"chgId"`
	Offline              *bool          `json:"offline,omitempty"`
	SdfHandl             *bool          `json:"sdfHandl,omitempty"`
	ServiceId            *int           `json:"serviceId,omitempty"`
	AppSvcProvId         string         `json:"appSvcProvId,omitempty"`
}
