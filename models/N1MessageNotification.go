package models

type N1MessageNotification struct {
	N1NotifySubscriptionId    string                        `json:"n1NotifySubscriptionId,omitempty"`
	LcsCorrelationId          string                        `json:"lcsCorrelationId,omitempty"`
	RegistrationCtxtContainer *RegistrationContextContainer `json:"registrationCtxtContainer,omitempty"`
	Guami                     *Guami                        `json:"guami,omitempty"`
	CIoT5GSOptimisation       *bool                         `json:"cIoT5GSOptimisation,omitempty"`
	Ecgi                      *Ecgi                         `json:"ecgi,omitempty"`
	N1MessageContainer        N1MessageContainer            `json:"n1MessageContainer"`
	NewLmfIdentification      string                        `json:"newLmfIdentification,omitempty"`
	Ncgi                      *Ncgi                         `json:"ncgi,omitempty"`
}
