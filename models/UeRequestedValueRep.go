package models
type UeRequestedValueRep struct {
	 PlmnId	*PlmnIdNid	`json:"plmnId,omitempty"`
	 ConnectState	CmState	`json:"connectState,omitempty"`
	 UserLoc	*UserLocation	`json:"userLoc,omitempty"`
	 PraStatuses	map[string]PresenceInfo	`json:"praStatuses,omitempty"`
}
