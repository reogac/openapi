package models

type ChargingData struct {
	ReportingLevel       ReportingLevel `json:"reportingLevel,omitempty"`
	SponsorId            string         `json:"sponsorId,omitempty"`
	AppSvcProvId         string         `json:"appSvcProvId,omitempty"`
	AfChargingIdentifier *int           `json:"afChargingIdentifier,omitempty"`
	MeteringMethod       MeteringMethod `json:"meteringMethod,omitempty"`
	Online               *bool          `json:"online,omitempty"`
	SdfHandl             *bool          `json:"sdfHandl,omitempty"`
	RatingGroup          *int           `json:"ratingGroup,omitempty"`
	ServiceId            *int           `json:"serviceId,omitempty"`
	AfChargId            string         `json:"afChargId,omitempty"`
	ChgId                string         `json:"chgId"`
	Offline              *bool          `json:"offline,omitempty"`
}
