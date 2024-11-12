package models

type AreaOfInterestEventState struct {
	IndividualPraIdList []string      `json:"individualPraIdList,omitempty"`
	Presence            PresenceState `json:"presence"`
}
