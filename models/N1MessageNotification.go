/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:23:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N1MessageNotification struct {
	Ncgi                      *Ncgi                         `json:"ncgi,omitempty"`
	LcsCorrelationId          string                        `json:"lcsCorrelationId,omitempty"`
	RegistrationCtxtContainer *RegistrationContextContainer `json:"registrationCtxtContainer,omitempty"`
	NewLmfIdentification      string                        `json:"newLmfIdentification,omitempty"`
	CIoT5GSOptimisation       *bool                         `json:"cIoT5GSOptimisation,omitempty"`
	Ecgi                      *Ecgi                         `json:"ecgi,omitempty"`
	N1NotifySubscriptionId    string                        `json:"n1NotifySubscriptionId,omitempty"`
	N1MessageContainer        N1MessageContainer            `json:"n1MessageContainer"`
	Guami                     *Guami                        `json:"guami,omitempty"`
}
