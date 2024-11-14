package models
type ChargingInformation struct {
	 SecondaryChfAddress	string	`json:"secondaryChfAddress,omitempty"`
	 PrimaryChfSetId	string	`json:"primaryChfSetId,omitempty"`
	 PrimaryChfInstanceId	string	`json:"primaryChfInstanceId,omitempty"`
	 SecondaryChfSetId	string	`json:"secondaryChfSetId,omitempty"`
	 SecondaryChfInstanceId	string	`json:"secondaryChfInstanceId,omitempty"`
	 PrimaryChfAddress	string	`json:"primaryChfAddress"`
}
