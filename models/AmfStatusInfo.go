package models

type AmfStatusInfo struct {
	TargetAmfFailure string  `json:"targetAmfFailure,omitempty"`
	GuamiList        []Guami `json:"guamiList"`
	StatusChange     string  `json:"statusChange"`
	TargetAmfRemoval string  `json:"targetAmfRemoval,omitempty"`
}
