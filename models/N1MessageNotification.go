package models

type N1MessageNotification struct {
	Ecgi                      *Ecgi                         `json:"ecgi,omitempty"`
	Ncgi                      *Ncgi                         `json:"ncgi,omitempty"`
	N1NotifySubscriptionId    string                        `json:"n1NotifySubscriptionId,omitempty"`
	N1MessageContainer        N1MessageContainer            `json:"n1MessageContainer"`
	NewLmfIdentification      string                        `json:"newLmfIdentification,omitempty"`
	Guami                     *Guami                        `json:"guami,omitempty"`
	LcsCorrelationId          string                        `json:"lcsCorrelationId,omitempty"`
	RegistrationCtxtContainer *RegistrationContextContainer `json:"registrationCtxtContainer,omitempty"`
	CIoT5GSOptimisation       *bool                         `json:"cIoT5GSOptimisation,omitempty"`
}
