/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:40 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmRequestedValueRep struct {
	AccessTypes       []string                `json:"accessTypes,omitempty"`
	RatTypes          []string                `json:"ratTypes,omitempty"`
	AllowedSnssais    []Snssai                `json:"allowedSnssais,omitempty"`
	N3gAllowedSnssais []Snssai                `json:"n3gAllowedSnssais,omitempty"`
	UserLoc           *UserLocation           `json:"userLoc,omitempty"`
	PraStatuses       map[string]PresenceInfo `json:"praStatuses,omitempty"`
}
