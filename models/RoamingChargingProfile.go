package models

type RoamingChargingProfile struct {
	Triggers            []Trigger           `json:"triggers,omitempty"`
	PartialRecordMethod PartialRecordMethod `json:"partialRecordMethod,omitempty"`
}
