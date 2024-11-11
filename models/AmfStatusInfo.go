package models

type AmfStatusInfo struct {
	GuamiList        []Guami `json:"guamiList"`
	StatusChange     string  `json:"statusChange"`
	TargetAmfRemoval string  `json:"targetAmfRemoval,omitempty"`
	TargetAmfFailure string  `json:"targetAmfFailure,omitempty"`
}
