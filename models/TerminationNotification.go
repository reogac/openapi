package models

type TerminationNotification struct {
	ResourceUri string                        `json:"resourceUri"`
	Cause       PolicyAssociationReleaseCause `json:"cause"`
}
