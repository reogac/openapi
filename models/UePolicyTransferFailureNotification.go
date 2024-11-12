package models

type UePolicyTransferFailureNotification struct {
	Cause N1N2MessageTransferCause `json:"cause"`
	Ptis  []int                    `json:"ptis"`
}
