/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:39 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmfRegistration struct {
	SmfInstanceId               string             `json:"smfInstanceId"`
	SmfSetId                    string             `json:"smfSetId,omitempty"`
	Dnn                         string             `json:"dnn,omitempty"`
	PlmnId                      PlmnId             `json:"plmnId"`
	PgwFqdn                     string             `json:"pgwFqdn,omitempty"`
	DataRestorationCallbackUri  string             `json:"dataRestorationCallbackUri,omitempty"`
	UdrRestartInd               *bool              `json:"udrRestartInd,omitempty"`
	PgwIpAddr                   *IpAddress         `json:"pgwIpAddr,omitempty"`
	RegistrationTime            string             `json:"registrationTime,omitempty"`
	PcfId                       string             `json:"pcfId,omitempty"`
	ResetIds                    []string           `json:"resetIds,omitempty"`
	SupportedFeatures           string             `json:"supportedFeatures,omitempty"`
	PcscfRestorationCallbackUri string             `json:"pcscfRestorationCallbackUri,omitempty"`
	RegistrationReason          RegistrationReason `json:"registrationReason,omitempty"`
	LastSynchronizationTime     string             `json:"lastSynchronizationTime,omitempty"`
	PduSessionId                int                `json:"pduSessionId"`
	SingleNssai                 Snssai             `json:"singleNssai"`
	EmergencyServices           *bool              `json:"emergencyServices,omitempty"`
	EpdgInd                     *bool              `json:"epdgInd,omitempty"`
	DeregCallbackUri            string             `json:"deregCallbackUri,omitempty"`
	ContextInfo                 *ContextInfo       `json:"contextInfo,omitempty"`
}
