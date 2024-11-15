/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:09:26 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UePolicyTransferFailureNotification struct {
	Ptis  []int                    `json:"ptis"`
	Cause N1N2MessageTransferCause `json:"cause"`
}
