/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N1MessageNotification struct {
	N1MessageContainer        N1MessageContainer            `json:"n1MessageContainer"`
	LcsCorrelationId          string                        `json:"lcsCorrelationId,omitempty"`
	RegistrationCtxtContainer *RegistrationContextContainer `json:"registrationCtxtContainer,omitempty"`
	NewLmfIdentification      string                        `json:"newLmfIdentification,omitempty"`
	Ecgi                      *Ecgi                         `json:"ecgi,omitempty"`
	N1NotifySubscriptionId    string                        `json:"n1NotifySubscriptionId,omitempty"`
	Guami                     *Guami                        `json:"guami,omitempty"`
	CIoT5GSOptimisation       *bool                         `json:"cIoT5GSOptimisation,omitempty"`
	Ncgi                      *Ncgi                         `json:"ncgi,omitempty"`
}
