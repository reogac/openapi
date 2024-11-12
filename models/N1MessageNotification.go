package models

type N1MessageNotification struct {
	RegistrationCtxtContainer *RegistrationContextContainer `json:"registrationCtxtContainer,omitempty"`
	Guami                     *Guami                        `json:"guami,omitempty"`
	CIoT5GSOptimisation       *bool                         `json:"cIoT5GSOptimisation,omitempty"`
	NewLmfIdentification      string                        `json:"newLmfIdentification,omitempty"`
	Ecgi                      *Ecgi                         `json:"ecgi,omitempty"`
	Ncgi                      *Ncgi                         `json:"ncgi,omitempty"`
	N1NotifySubscriptionId    string                        `json:"n1NotifySubscriptionId,omitempty"`
	N1MessageContainer        N1MessageContainer            `json:"n1MessageContainer"`
	LcsCorrelationId          string                        `json:"lcsCorrelationId,omitempty"`
}
