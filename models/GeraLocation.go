/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:23:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type GeraLocation struct {
	Cgi                      *CellGlobalId   `json:"cgi,omitempty"`
	MscNumber                string          `json:"mscNumber,omitempty"`
	UeLocationTimestamp      string          `json:"ueLocationTimestamp,omitempty"`
	GeographicalInformation  string          `json:"geographicalInformation,omitempty"`
	AgeOfLocationInformation *int            `json:"ageOfLocationInformation,omitempty"`
	GeodeticInformation      string          `json:"geodeticInformation,omitempty"`
	LocationNumber           string          `json:"locationNumber,omitempty"`
	Sai                      *ServiceAreaId  `json:"sai,omitempty"`
	Lai                      *LocationAreaId `json:"lai,omitempty"`
	Rai                      *RoutingAreaId  `json:"rai,omitempty"`
	VlrNumber                string          `json:"vlrNumber,omitempty"`
}
