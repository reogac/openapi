package models

type SmfRegistration struct {
	SupportedFeatures           string             `json:"supportedFeatures,omitempty"`
	PgwIpAddr                   *IpAddress         `json:"pgwIpAddr,omitempty"`
	RegistrationReason          RegistrationReason `json:"registrationReason,omitempty"`
	ContextInfo                 *ContextInfo       `json:"contextInfo,omitempty"`
	LastSynchronizationTime     string             `json:"lastSynchronizationTime,omitempty"`
	Dnn                         string             `json:"dnn,omitempty"`
	PcscfRestorationCallbackUri string             `json:"pcscfRestorationCallbackUri,omitempty"`
	SmfInstanceId               string             `json:"smfInstanceId"`
	PduSessionId                int                `json:"pduSessionId"`
	SingleNssai                 Snssai             `json:"singleNssai"`
	RegistrationTime            string             `json:"registrationTime,omitempty"`
	PcfId                       string             `json:"pcfId,omitempty"`
	ResetIds                    []string           `json:"resetIds,omitempty"`
	SmfSetId                    string             `json:"smfSetId,omitempty"`
	EmergencyServices           *bool              `json:"emergencyServices,omitempty"`
	PlmnId                      PlmnId             `json:"plmnId"`
	PgwFqdn                     string             `json:"pgwFqdn,omitempty"`
	EpdgInd                     *bool              `json:"epdgInd,omitempty"`
	DeregCallbackUri            string             `json:"deregCallbackUri,omitempty"`
	DataRestorationCallbackUri  string             `json:"dataRestorationCallbackUri,omitempty"`
	UdrRestartInd               *bool              `json:"udrRestartInd,omitempty"`
}
