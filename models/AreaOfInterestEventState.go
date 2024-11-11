package models

type AreaOfInterestEventState struct {
	Presence            string   `json:"presence"`
	IndividualPraIdList []string `json:"individualPraIdList,omitempty"`
}
