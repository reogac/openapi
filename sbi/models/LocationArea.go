/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type LocationArea struct {
	GeographicAreas []GeographicArea `json:"geographicAreas,omitempty"`
	CivicAddresses  []CivicAddress   `json:"civicAddresses,omitempty"`
	NwAreaInfo      *NetworkAreaInfo `json:"nwAreaInfo,omitempty"`
	UmtTime         *UmtTime         `json:"umtTime,omitempty"`
}
