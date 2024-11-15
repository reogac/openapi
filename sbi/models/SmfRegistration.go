/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:11:55 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmfRegistration struct {
	LastSynchronizationTime     string             `json:"lastSynchronizationTime,omitempty"`
	DeregCallbackUri            string             `json:"deregCallbackUri,omitempty"`
	UdrRestartInd               *bool              `json:"udrRestartInd,omitempty"`
	SmfSetId                    string             `json:"smfSetId,omitempty"`
	SupportedFeatures           string             `json:"supportedFeatures,omitempty"`
	PduSessionId                int                `json:"pduSessionId"`
	EmergencyServices           *bool              `json:"emergencyServices,omitempty"`
	EpdgInd                     *bool              `json:"epdgInd,omitempty"`
	RegistrationTime            string             `json:"registrationTime,omitempty"`
	ContextInfo                 *ContextInfo       `json:"contextInfo,omitempty"`
	DataRestorationCallbackUri  string             `json:"dataRestorationCallbackUri,omitempty"`
	Dnn                         string             `json:"dnn,omitempty"`
	PcscfRestorationCallbackUri string             `json:"pcscfRestorationCallbackUri,omitempty"`
	PgwFqdn                     string             `json:"pgwFqdn,omitempty"`
	PgwIpAddr                   *IpAddress         `json:"pgwIpAddr,omitempty"`
	RegistrationReason          RegistrationReason `json:"registrationReason,omitempty"`
	ResetIds                    []string           `json:"resetIds,omitempty"`
	SmfInstanceId               string             `json:"smfInstanceId"`
	SingleNssai                 Snssai             `json:"singleNssai"`
	PlmnId                      PlmnId             `json:"plmnId"`
	PcfId                       string             `json:"pcfId,omitempty"`
}
