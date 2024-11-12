package models

type IpSmGwInfo struct {
	IpSmGwRegistration *IpSmGwRegistration `json:"ipSmGwRegistration,omitempty"`
	IpSmGwGuidance     *IpSmGwGuidance     `json:"ipSmGwGuidance,omitempty"`
}
