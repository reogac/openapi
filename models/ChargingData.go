package models

type ChargingData struct {
	AfChargId            string         `json:"afChargId,omitempty"`
	SdfHandl             *bool          `json:"sdfHandl,omitempty"`
	RatingGroup          *int           `json:"ratingGroup,omitempty"`
	ReportingLevel       ReportingLevel `json:"reportingLevel,omitempty"`
	ServiceId            *int           `json:"serviceId,omitempty"`
	AfChargingIdentifier *int           `json:"afChargingIdentifier,omitempty"`
	AppSvcProvId         string         `json:"appSvcProvId,omitempty"`
	ChgId                string         `json:"chgId"`
	MeteringMethod       MeteringMethod `json:"meteringMethod,omitempty"`
	Offline              *bool          `json:"offline,omitempty"`
	Online               *bool          `json:"online,omitempty"`
	SponsorId            string         `json:"sponsorId,omitempty"`
}
