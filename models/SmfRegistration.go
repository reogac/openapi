package models

type SmfRegistration struct {
	SmfSetId                    string             `json:"smfSetId,omitempty"`
	PcscfRestorationCallbackUri string             `json:"pcscfRestorationCallbackUri,omitempty"`
	DataRestorationCallbackUri  string             `json:"dataRestorationCallbackUri,omitempty"`
	SmfInstanceId               string             `json:"smfInstanceId"`
	SingleNssai                 Snssai             `json:"singleNssai"`
	PlmnId                      PlmnId             `json:"plmnId"`
	EpdgInd                     *bool              `json:"epdgInd,omitempty"`
	RegistrationTime            string             `json:"registrationTime,omitempty"`
	UdrRestartInd               *bool              `json:"udrRestartInd,omitempty"`
	PgwFqdn                     string             `json:"pgwFqdn,omitempty"`
	RegistrationReason          RegistrationReason `json:"registrationReason,omitempty"`
	LastSynchronizationTime     string             `json:"lastSynchronizationTime,omitempty"`
	SupportedFeatures           string             `json:"supportedFeatures,omitempty"`
	PduSessionId                int                `json:"pduSessionId"`
	Dnn                         string             `json:"dnn,omitempty"`
	EmergencyServices           *bool              `json:"emergencyServices,omitempty"`
	PgwIpAddr                   *IpAddress         `json:"pgwIpAddr,omitempty"`
	DeregCallbackUri            string             `json:"deregCallbackUri,omitempty"`
	ContextInfo                 *ContextInfo       `json:"contextInfo,omitempty"`
	PcfId                       string             `json:"pcfId,omitempty"`
	ResetIds                    []string           `json:"resetIds,omitempty"`
}
