package models

type SmfRegistration struct {
	PcscfRestorationCallbackUri string             `json:"pcscfRestorationCallbackUri,omitempty"`
	EpdgInd                     *bool              `json:"epdgInd,omitempty"`
	RegistrationReason          RegistrationReason `json:"registrationReason,omitempty"`
	RegistrationTime            string             `json:"registrationTime,omitempty"`
	ContextInfo                 *ContextInfo       `json:"contextInfo,omitempty"`
	SupportedFeatures           string             `json:"supportedFeatures,omitempty"`
	SingleNssai                 Snssai             `json:"singleNssai"`
	Dnn                         string             `json:"dnn,omitempty"`
	PcfId                       string             `json:"pcfId,omitempty"`
	EmergencyServices           *bool              `json:"emergencyServices,omitempty"`
	PlmnId                      PlmnId             `json:"plmnId"`
	PgwFqdn                     string             `json:"pgwFqdn,omitempty"`
	DeregCallbackUri            string             `json:"deregCallbackUri,omitempty"`
	DataRestorationCallbackUri  string             `json:"dataRestorationCallbackUri,omitempty"`
	SmfInstanceId               string             `json:"smfInstanceId"`
	SmfSetId                    string             `json:"smfSetId,omitempty"`
	PduSessionId                int                `json:"pduSessionId"`
	ResetIds                    []string           `json:"resetIds,omitempty"`
	UdrRestartInd               *bool              `json:"udrRestartInd,omitempty"`
	LastSynchronizationTime     string             `json:"lastSynchronizationTime,omitempty"`
	PgwIpAddr                   *IpAddress         `json:"pgwIpAddr,omitempty"`
}
