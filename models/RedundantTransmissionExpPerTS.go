package models

type RedundantTransmissionExpPerTS struct {
	Confidence      *int                      `json:"confidence,omitempty"`
	TsStart         string                    `json:"tsStart"`
	TsDuration      int                       `json:"tsDuration"`
	ObsvRedTransExp ObservedRedundantTransExp `json:"obsvRedTransExp"`
	RedTransStatus  *bool                     `json:"redTransStatus,omitempty"`
	UeRatio         *int                      `json:"ueRatio,omitempty"`
}
