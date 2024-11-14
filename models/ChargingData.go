package models
type ChargingData struct {
	 ServiceId	*int	`json:"serviceId,omitempty"`
	 AppSvcProvId	string	`json:"appSvcProvId,omitempty"`
	 AfChargingIdentifier	*int	`json:"afChargingIdentifier,omitempty"`
	 ChgId	string	`json:"chgId"`
	 MeteringMethod	MeteringMethod	`json:"meteringMethod,omitempty"`
	 Offline	*bool	`json:"offline,omitempty"`
	 RatingGroup	*int	`json:"ratingGroup,omitempty"`
	 ReportingLevel	ReportingLevel	`json:"reportingLevel,omitempty"`
	 Online	*bool	`json:"online,omitempty"`
	 SdfHandl	*bool	`json:"sdfHandl,omitempty"`
	 SponsorId	string	`json:"sponsorId,omitempty"`
	 AfChargId	string	`json:"afChargId,omitempty"`
}
