/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:39 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmfRegistration struct {
	EmergencyServices           *bool              `json:"emergencyServices,omitempty"`
	RegistrationReason          RegistrationReason `json:"registrationReason,omitempty"`
	RegistrationTime            string             `json:"registrationTime,omitempty"`
	SmfSetId                    string             `json:"smfSetId,omitempty"`
	SupportedFeatures           string             `json:"supportedFeatures,omitempty"`
	PduSessionId                int                `json:"pduSessionId"`
	Dnn                         string             `json:"dnn,omitempty"`
	PlmnId                      PlmnId             `json:"plmnId"`
	PgwFqdn                     string             `json:"pgwFqdn,omitempty"`
	ContextInfo                 *ContextInfo       `json:"contextInfo,omitempty"`
	PcfId                       string             `json:"pcfId,omitempty"`
	SmfInstanceId               string             `json:"smfInstanceId"`
	DataRestorationCallbackUri  string             `json:"dataRestorationCallbackUri,omitempty"`
	ResetIds                    []string           `json:"resetIds,omitempty"`
	UdrRestartInd               *bool              `json:"udrRestartInd,omitempty"`
	SingleNssai                 Snssai             `json:"singleNssai"`
	PgwIpAddr                   *IpAddress         `json:"pgwIpAddr,omitempty"`
	EpdgInd                     *bool              `json:"epdgInd,omitempty"`
	DeregCallbackUri            string             `json:"deregCallbackUri,omitempty"`
	LastSynchronizationTime     string             `json:"lastSynchronizationTime,omitempty"`
	PcscfRestorationCallbackUri string             `json:"pcscfRestorationCallbackUri,omitempty"`
}
