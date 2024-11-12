package models

type AmfStatusInfo struct {
	StatusChange     StatusChange `json:"statusChange"`
	TargetAmfRemoval string       `json:"targetAmfRemoval,omitempty"`
	TargetAmfFailure string       `json:"targetAmfFailure,omitempty"`
	GuamiList        []Guami      `json:"guamiList"`
}
