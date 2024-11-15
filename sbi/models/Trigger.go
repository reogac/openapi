/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:43 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type Trigger struct {
	TimeLimit        *int            `json:"timeLimit,omitempty"`
	VolumeLimit      *int            `json:"volumeLimit,omitempty"`
	VolumeLimit64    *int            `json:"volumeLimit64,omitempty"`
	EventLimit       *int            `json:"eventLimit,omitempty"`
	MaxNumberOfccc   *int            `json:"maxNumberOfccc,omitempty"`
	TariffTimeChange string          `json:"tariffTimeChange,omitempty"`
	TriggerType      TriggerType     `json:"triggerType,omitempty"`
	TriggerCategory  TriggerCategory `json:"triggerCategory"`
}
