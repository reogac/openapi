package models
type UeRequestedValueRep struct {
	 UserLoc	*UserLocation	`json:"userLoc,omitempty"`
	 PraStatuses	map[string]PresenceInfo	`json:"praStatuses,omitempty"`
	 PlmnId	*PlmnIdNid	`json:"plmnId,omitempty"`
	 ConnectState	CmState	`json:"connectState,omitempty"`
}
