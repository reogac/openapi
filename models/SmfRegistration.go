package models

type SmfRegistration struct {
	ContextInfo                 *ContextInfo       `json:"contextInfo,omitempty"`
	SupportedFeatures           string             `json:"supportedFeatures,omitempty"`
	PduSessionId                int                `json:"pduSessionId"`
	Dnn                         string             `json:"dnn,omitempty"`
	DeregCallbackUri            string             `json:"deregCallbackUri,omitempty"`
	ResetIds                    []string           `json:"resetIds,omitempty"`
	UdrRestartInd               *bool              `json:"udrRestartInd,omitempty"`
	LastSynchronizationTime     string             `json:"lastSynchronizationTime,omitempty"`
	PcscfRestorationCallbackUri string             `json:"pcscfRestorationCallbackUri,omitempty"`
	PgwFqdn                     string             `json:"pgwFqdn,omitempty"`
	EpdgInd                     *bool              `json:"epdgInd,omitempty"`
	RegistrationReason          RegistrationReason `json:"registrationReason,omitempty"`
	EmergencyServices           *bool              `json:"emergencyServices,omitempty"`
	PlmnId                      PlmnId             `json:"plmnId"`
	PcfId                       string             `json:"pcfId,omitempty"`
	RegistrationTime            string             `json:"registrationTime,omitempty"`
	DataRestorationCallbackUri  string             `json:"dataRestorationCallbackUri,omitempty"`
	SmfInstanceId               string             `json:"smfInstanceId"`
	SmfSetId                    string             `json:"smfSetId,omitempty"`
	SingleNssai                 Snssai             `json:"singleNssai"`
	PgwIpAddr                   *IpAddress         `json:"pgwIpAddr,omitempty"`
}
