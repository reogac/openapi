package models
type RoamingChargingProfile struct {
	 PartialRecordMethod	string	`json:"partialRecordMethod,omitempty"`
	 Triggers	[]Trigger	`json:"triggers,omitempty"`
}
