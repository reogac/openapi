package models
type AreaOfInterestEventState struct {
	 Presence	PresenceState	`json:"presence"`
	 IndividualPraIdList	[]string	`json:"individualPraIdList,omitempty"`
}
