package models

type AfCoordinationInfo struct {
	SourceDnai           string             `json:"sourceDnai,omitempty"`
	SourceUeIpv4Addr     string             `json:"sourceUeIpv4Addr,omitempty"`
	SourceUeIpv6Prefix   string             `json:"sourceUeIpv6Prefix,omitempty"`
	NotificationInfoList []NotificationInfo `json:"notificationInfoList,omitempty"`
}
