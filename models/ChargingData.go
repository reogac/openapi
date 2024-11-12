package models

type ChargingData struct {
	ReportingLevel       ReportingLevel `json:"reportingLevel,omitempty"`
	ServiceId            *int           `json:"serviceId,omitempty"`
	SponsorId            string         `json:"sponsorId,omitempty"`
	ChgId                string         `json:"chgId"`
	MeteringMethod       MeteringMethod `json:"meteringMethod,omitempty"`
	Offline              *bool          `json:"offline,omitempty"`
	Online               *bool          `json:"online,omitempty"`
	RatingGroup          *int           `json:"ratingGroup,omitempty"`
	AfChargId            string         `json:"afChargId,omitempty"`
	SdfHandl             *bool          `json:"sdfHandl,omitempty"`
	AppSvcProvId         string         `json:"appSvcProvId,omitempty"`
	AfChargingIdentifier *int           `json:"afChargingIdentifier,omitempty"`
}
