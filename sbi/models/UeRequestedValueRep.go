/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:42 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeRequestedValueRep struct {
	PraStatuses  map[string]PresenceInfo `json:"praStatuses,omitempty"`
	PlmnId       *PlmnIdNid              `json:"plmnId,omitempty"`
	ConnectState CmState                 `json:"connectState,omitempty"`
	UserLoc      *UserLocation           `json:"userLoc,omitempty"`
}
