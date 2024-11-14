/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:44 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ReportingInformation struct {
	PartitionCriteria []string           `json:"partitionCriteria,omitempty"`
	GrpRepTime        *int               `json:"grpRepTime,omitempty"`
	NotifMethod       NotificationMethod `json:"notifMethod,omitempty"`
	MonDur            string             `json:"monDur,omitempty"`
	RepPeriod         *int               `json:"repPeriod,omitempty"`
	NotifFlag         NotificationFlag   `json:"notifFlag,omitempty"`
	ImmRep            *bool              `json:"immRep,omitempty"`
	MaxReportNbr      *int               `json:"maxReportNbr,omitempty"`
	SampRatio         *int               `json:"sampRatio,omitempty"`
}
