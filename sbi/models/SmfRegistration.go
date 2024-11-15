/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:10 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmfRegistration struct {
	EmergencyServices           *bool              `json:"emergencyServices,omitempty"`
	PlmnId                      PlmnId             `json:"plmnId"`
	PgwFqdn                     string             `json:"pgwFqdn,omitempty"`
	PgwIpAddr                   *IpAddress         `json:"pgwIpAddr,omitempty"`
	ContextInfo                 *ContextInfo       `json:"contextInfo,omitempty"`
	SmfInstanceId               string             `json:"smfInstanceId"`
	PduSessionId                int                `json:"pduSessionId"`
	Dnn                         string             `json:"dnn,omitempty"`
	PcfId                       string             `json:"pcfId,omitempty"`
	DataRestorationCallbackUri  string             `json:"dataRestorationCallbackUri,omitempty"`
	UdrRestartInd               *bool              `json:"udrRestartInd,omitempty"`
	SingleNssai                 Snssai             `json:"singleNssai"`
	DeregCallbackUri            string             `json:"deregCallbackUri,omitempty"`
	SupportedFeatures           string             `json:"supportedFeatures,omitempty"`
	RegistrationTime            string             `json:"registrationTime,omitempty"`
	RegistrationReason          RegistrationReason `json:"registrationReason,omitempty"`
	ResetIds                    []string           `json:"resetIds,omitempty"`
	LastSynchronizationTime     string             `json:"lastSynchronizationTime,omitempty"`
	SmfSetId                    string             `json:"smfSetId,omitempty"`
	PcscfRestorationCallbackUri string             `json:"pcscfRestorationCallbackUri,omitempty"`
	EpdgInd                     *bool              `json:"epdgInd,omitempty"`
}
