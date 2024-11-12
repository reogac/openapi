package models

type RoamingChargingProfile struct {
	PartialRecordMethod PartialRecordMethod `json:"partialRecordMethod,omitempty"`
	Triggers            []Trigger           `json:"triggers,omitempty"`
}
