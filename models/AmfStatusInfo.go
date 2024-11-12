package models
type AmfStatusInfo struct {
	 TargetAmfRemoval	string	`json:"targetAmfRemoval,omitempty"`
	 TargetAmfFailure	string	`json:"targetAmfFailure,omitempty"`
	 GuamiList	[]Guami	`json:"guamiList"`
	 StatusChange	StatusChange	`json:"statusChange"`
}
