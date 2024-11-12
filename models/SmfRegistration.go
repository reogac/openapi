package models

type SmfRegistration struct {
	EmergencyServices           *bool              `json:"emergencyServices,omitempty"`
	RegistrationReason          RegistrationReason `json:"registrationReason,omitempty"`
	RegistrationTime            string             `json:"registrationTime,omitempty"`
	ContextInfo                 *ContextInfo       `json:"contextInfo,omitempty"`
	SmfSetId                    string             `json:"smfSetId,omitempty"`
	Dnn                         string             `json:"dnn,omitempty"`
	PcscfRestorationCallbackUri string             `json:"pcscfRestorationCallbackUri,omitempty"`
	PlmnId                      PlmnId             `json:"plmnId"`
	DataRestorationCallbackUri  string             `json:"dataRestorationCallbackUri,omitempty"`
	ResetIds                    []string           `json:"resetIds,omitempty"`
	PduSessionId                int                `json:"pduSessionId"`
	SingleNssai                 Snssai             `json:"singleNssai"`
	PgwIpAddr                   *IpAddress         `json:"pgwIpAddr,omitempty"`
	EpdgInd                     *bool              `json:"epdgInd,omitempty"`
	DeregCallbackUri            string             `json:"deregCallbackUri,omitempty"`
	PcfId                       string             `json:"pcfId,omitempty"`
	UdrRestartInd               *bool              `json:"udrRestartInd,omitempty"`
	LastSynchronizationTime     string             `json:"lastSynchronizationTime,omitempty"`
	SmfInstanceId               string             `json:"smfInstanceId"`
	PgwFqdn                     string             `json:"pgwFqdn,omitempty"`
	SupportedFeatures           string             `json:"supportedFeatures,omitempty"`
}
