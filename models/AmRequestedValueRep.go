/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:22:56 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmRequestedValueRep struct {
	RatTypes          []string                `json:"ratTypes,omitempty"`
	AllowedSnssais    []Snssai                `json:"allowedSnssais,omitempty"`
	N3gAllowedSnssais []Snssai                `json:"n3gAllowedSnssais,omitempty"`
	UserLoc           *UserLocation           `json:"userLoc,omitempty"`
	PraStatuses       map[string]PresenceInfo `json:"praStatuses,omitempty"`
	AccessTypes       []string                `json:"accessTypes,omitempty"`
}
