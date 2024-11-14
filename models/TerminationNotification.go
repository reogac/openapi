package models
type TerminationNotification struct {
	 Cause	PolicyAssociationReleaseCause	`json:"cause"`
	 ResourceUri	string	`json:"resourceUri"`
}
