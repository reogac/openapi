package models

type SmfRegistration struct {
	SmfInstanceId               string             `json:"smfInstanceId"`
	SupportedFeatures           string             `json:"supportedFeatures,omitempty"`
	PduSessionId                int                `json:"pduSessionId"`
	Dnn                         string             `json:"dnn,omitempty"`
	EpdgInd                     *bool              `json:"epdgInd,omitempty"`
	PcfId                       string             `json:"pcfId,omitempty"`
	LastSynchronizationTime     string             `json:"lastSynchronizationTime,omitempty"`
	SmfSetId                    string             `json:"smfSetId,omitempty"`
	PgwFqdn                     string             `json:"pgwFqdn,omitempty"`
	DeregCallbackUri            string             `json:"deregCallbackUri,omitempty"`
	RegistrationReason          RegistrationReason `json:"registrationReason,omitempty"`
	RegistrationTime            string             `json:"registrationTime,omitempty"`
	ContextInfo                 *ContextInfo       `json:"contextInfo,omitempty"`
	DataRestorationCallbackUri  string             `json:"dataRestorationCallbackUri,omitempty"`
	EmergencyServices           *bool              `json:"emergencyServices,omitempty"`
	PcscfRestorationCallbackUri string             `json:"pcscfRestorationCallbackUri,omitempty"`
	PgwIpAddr                   *IpAddress         `json:"pgwIpAddr,omitempty"`
	SingleNssai                 Snssai             `json:"singleNssai"`
	PlmnId                      PlmnId             `json:"plmnId"`
	ResetIds                    []string           `json:"resetIds,omitempty"`
	UdrRestartInd               *bool              `json:"udrRestartInd,omitempty"`
}
