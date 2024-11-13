package models

type ChargingData struct {
	ChgId                string         `json:"chgId"`
	MeteringMethod       MeteringMethod `json:"meteringMethod,omitempty"`
	RatingGroup          *int           `json:"ratingGroup,omitempty"`
	AppSvcProvId         string         `json:"appSvcProvId,omitempty"`
	Offline              *bool          `json:"offline,omitempty"`
	Online               *bool          `json:"online,omitempty"`
	SdfHandl             *bool          `json:"sdfHandl,omitempty"`
	ReportingLevel       ReportingLevel `json:"reportingLevel,omitempty"`
	ServiceId            *int           `json:"serviceId,omitempty"`
	SponsorId            string         `json:"sponsorId,omitempty"`
	AfChargingIdentifier *int           `json:"afChargingIdentifier,omitempty"`
	AfChargId            string         `json:"afChargId,omitempty"`
}
