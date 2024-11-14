package models
type ChargingData struct {
	 AfChargingIdentifier	*int	`json:"afChargingIdentifier,omitempty"`
	 ChgId	string	`json:"chgId"`
	 Online	*bool	`json:"online,omitempty"`
	 SdfHandl	*bool	`json:"sdfHandl,omitempty"`
	 AppSvcProvId	string	`json:"appSvcProvId,omitempty"`
	 ServiceId	*int	`json:"serviceId,omitempty"`
	 SponsorId	string	`json:"sponsorId,omitempty"`
	 AfChargId	string	`json:"afChargId,omitempty"`
	 MeteringMethod	MeteringMethod	`json:"meteringMethod,omitempty"`
	 Offline	*bool	`json:"offline,omitempty"`
	 RatingGroup	*int	`json:"ratingGroup,omitempty"`
	 ReportingLevel	ReportingLevel	`json:"reportingLevel,omitempty"`
}
