/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:14 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AfCoordinationInfo struct {
	SourceDnai           string             `json:"sourceDnai,omitempty"`
	SourceUeIpv4Addr     string             `json:"sourceUeIpv4Addr,omitempty"`
	SourceUeIpv6Prefix   string             `json:"sourceUeIpv6Prefix,omitempty"`
	NotificationInfoList []NotificationInfo `json:"notificationInfoList,omitempty"`
}