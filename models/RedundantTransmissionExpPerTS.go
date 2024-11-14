/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:44 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RedundantTransmissionExpPerTS struct {
	TsStart         string                    `json:"tsStart"`
	TsDuration      int                       `json:"tsDuration"`
	ObsvRedTransExp ObservedRedundantTransExp `json:"obsvRedTransExp"`
	RedTransStatus  *bool                     `json:"redTransStatus,omitempty"`
	UeRatio         *int                      `json:"ueRatio,omitempty"`
	Confidence      *int                      `json:"confidence,omitempty"`
}
