package models

type N1MessageNotification struct {
	N1NotifySubscriptionId    string                        `json:"n1NotifySubscriptionId,omitempty"`
	LcsCorrelationId          string                        `json:"lcsCorrelationId,omitempty"`
	Ecgi                      *Ecgi                         `json:"ecgi,omitempty"`
	Ncgi                      *Ncgi                         `json:"ncgi,omitempty"`
	N1MessageContainer        N1MessageContainer            `json:"n1MessageContainer"`
	RegistrationCtxtContainer *RegistrationContextContainer `json:"registrationCtxtContainer,omitempty"`
	NewLmfIdentification      string                        `json:"newLmfIdentification,omitempty"`
	Guami                     *Guami                        `json:"guami,omitempty"`
	CIoT5GSOptimisation       *bool                         `json:"cIoT5GSOptimisation,omitempty"`
}
