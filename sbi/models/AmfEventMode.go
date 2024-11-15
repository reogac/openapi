/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfEventMode struct {
	Trigger              AmfEventTrigger  `json:"trigger"`
	MaxReports           *int             `json:"maxReports,omitempty"`
	Expiry               string           `json:"expiry,omitempty"`
	RepPeriod            *int             `json:"repPeriod,omitempty"`
	SampRatio            *int             `json:"sampRatio,omitempty"`
	PartitioningCriteria []string         `json:"partitioningCriteria,omitempty"`
	NotifFlag            NotificationFlag `json:"notifFlag,omitempty"`
}
