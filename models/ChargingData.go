package models

type ChargingData struct {
	Online               *bool          `json:"online,omitempty"`
	SdfHandl             *bool          `json:"sdfHandl,omitempty"`
	ServiceId            *int           `json:"serviceId,omitempty"`
	SponsorId            string         `json:"sponsorId,omitempty"`
	MeteringMethod       MeteringMethod `json:"meteringMethod,omitempty"`
	Offline              *bool          `json:"offline,omitempty"`
	ReportingLevel       ReportingLevel `json:"reportingLevel,omitempty"`
	AppSvcProvId         string         `json:"appSvcProvId,omitempty"`
	AfChargingIdentifier *int           `json:"afChargingIdentifier,omitempty"`
	AfChargId            string         `json:"afChargId,omitempty"`
	ChgId                string         `json:"chgId"`
	RatingGroup          *int           `json:"ratingGroup,omitempty"`
}
