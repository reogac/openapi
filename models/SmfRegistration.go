package models

type SmfRegistration struct {
	UdrRestartInd               *bool              `json:"udrRestartInd,omitempty"`
	SupportedFeatures           string             `json:"supportedFeatures,omitempty"`
	PduSessionId                int                `json:"pduSessionId"`
	PcscfRestorationCallbackUri string             `json:"pcscfRestorationCallbackUri,omitempty"`
	PgwFqdn                     string             `json:"pgwFqdn,omitempty"`
	ResetIds                    []string           `json:"resetIds,omitempty"`
	Dnn                         string             `json:"dnn,omitempty"`
	EpdgInd                     *bool              `json:"epdgInd,omitempty"`
	PcfId                       string             `json:"pcfId,omitempty"`
	LastSynchronizationTime     string             `json:"lastSynchronizationTime,omitempty"`
	SmfSetId                    string             `json:"smfSetId,omitempty"`
	PlmnId                      PlmnId             `json:"plmnId"`
	ContextInfo                 *ContextInfo       `json:"contextInfo,omitempty"`
	RegistrationReason          RegistrationReason `json:"registrationReason,omitempty"`
	RegistrationTime            string             `json:"registrationTime,omitempty"`
	DataRestorationCallbackUri  string             `json:"dataRestorationCallbackUri,omitempty"`
	SmfInstanceId               string             `json:"smfInstanceId"`
	SingleNssai                 Snssai             `json:"singleNssai"`
	EmergencyServices           *bool              `json:"emergencyServices,omitempty"`
	PgwIpAddr                   *IpAddress         `json:"pgwIpAddr,omitempty"`
	DeregCallbackUri            string             `json:"deregCallbackUri,omitempty"`
}
