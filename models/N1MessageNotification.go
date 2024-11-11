package models
type N1MessageNotification struct {
	 CIoT5GSOptimisation	*bool	`json:"cIoT5GSOptimisation,omitempty"`
	 Ecgi	*Ecgi	`json:"ecgi,omitempty"`
	 N1MessageContainer	N1MessageContainer	`json:"n1MessageContainer"`
	 LcsCorrelationId	string	`json:"lcsCorrelationId,omitempty"`
	 Guami	*Guami	`json:"guami,omitempty"`
	 Ncgi	*Ncgi	`json:"ncgi,omitempty"`
	 N1NotifySubscriptionId	string	`json:"n1NotifySubscriptionId,omitempty"`
	 RegistrationCtxtContainer	*RegistrationContextContainer	`json:"registrationCtxtContainer,omitempty"`
	 NewLmfIdentification	string	`json:"newLmfIdentification,omitempty"`
}
