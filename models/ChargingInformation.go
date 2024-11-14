/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:43 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ChargingInformation struct {
	SecondaryChfAddress    string `json:"secondaryChfAddress,omitempty"`
	PrimaryChfSetId        string `json:"primaryChfSetId,omitempty"`
	PrimaryChfInstanceId   string `json:"primaryChfInstanceId,omitempty"`
	SecondaryChfSetId      string `json:"secondaryChfSetId,omitempty"`
	SecondaryChfInstanceId string `json:"secondaryChfInstanceId,omitempty"`
	PrimaryChfAddress      string `json:"primaryChfAddress"`
}
