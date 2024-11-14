/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:22:55 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmfRegistration struct {
	SingleNssai                 Snssai             `json:"singleNssai"`
	PgwFqdn                     string             `json:"pgwFqdn,omitempty"`
	RegistrationReason          RegistrationReason `json:"registrationReason,omitempty"`
	PcfId                       string             `json:"pcfId,omitempty"`
	PduSessionId                int                `json:"pduSessionId"`
	Dnn                         string             `json:"dnn,omitempty"`
	PlmnId                      PlmnId             `json:"plmnId"`
	EpdgInd                     *bool              `json:"epdgInd,omitempty"`
	ContextInfo                 *ContextInfo       `json:"contextInfo,omitempty"`
	DataRestorationCallbackUri  string             `json:"dataRestorationCallbackUri,omitempty"`
	SmfInstanceId               string             `json:"smfInstanceId"`
	EmergencyServices           *bool              `json:"emergencyServices,omitempty"`
	PgwIpAddr                   *IpAddress         `json:"pgwIpAddr,omitempty"`
	UdrRestartInd               *bool              `json:"udrRestartInd,omitempty"`
	SmfSetId                    string             `json:"smfSetId,omitempty"`
	PcscfRestorationCallbackUri string             `json:"pcscfRestorationCallbackUri,omitempty"`
	DeregCallbackUri            string             `json:"deregCallbackUri,omitempty"`
	RegistrationTime            string             `json:"registrationTime,omitempty"`
	ResetIds                    []string           `json:"resetIds,omitempty"`
	LastSynchronizationTime     string             `json:"lastSynchronizationTime,omitempty"`
	SupportedFeatures           string             `json:"supportedFeatures,omitempty"`
}
