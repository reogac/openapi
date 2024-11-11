package models

type ChargingInformation struct {
	SecondaryChfSetId      string `json:"secondaryChfSetId,omitempty"`
	SecondaryChfInstanceId string `json:"secondaryChfInstanceId,omitempty"`
	PrimaryChfAddress      string `json:"primaryChfAddress"`
	SecondaryChfAddress    string `json:"secondaryChfAddress"`
	PrimaryChfSetId        string `json:"primaryChfSetId,omitempty"`
	PrimaryChfInstanceId   string `json:"primaryChfInstanceId,omitempty"`
}
