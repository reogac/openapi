/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:13 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UePolicyTransferFailureNotification struct {
	Cause N1N2MessageTransferCause `json:"cause"`
	Ptis  []int                    `json:"ptis"`
}
