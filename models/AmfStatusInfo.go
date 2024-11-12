package models

type AmfStatusInfo struct {
	GuamiList        []Guami      `json:"guamiList"`
	StatusChange     StatusChange `json:"statusChange"`
	TargetAmfRemoval string       `json:"targetAmfRemoval,omitempty"`
	TargetAmfFailure string       `json:"targetAmfFailure,omitempty"`
}
