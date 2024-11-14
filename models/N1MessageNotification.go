package models
type N1MessageNotification struct {
	 Ncgi	*Ncgi	`json:"ncgi,omitempty"`
	 N1NotifySubscriptionId	string	`json:"n1NotifySubscriptionId,omitempty"`
	 NewLmfIdentification	string	`json:"newLmfIdentification,omitempty"`
	 Guami	*Guami	`json:"guami,omitempty"`
	 Ecgi	*Ecgi	`json:"ecgi,omitempty"`
	 N1MessageContainer	N1MessageContainer	`json:"n1MessageContainer"`
	 LcsCorrelationId	string	`json:"lcsCorrelationId,omitempty"`
	 RegistrationCtxtContainer	*RegistrationContextContainer	`json:"registrationCtxtContainer,omitempty"`
	 CIoT5GSOptimisation	*bool	`json:"cIoT5GSOptimisation,omitempty"`
}
