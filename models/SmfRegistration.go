package models

type SmfRegistration struct {
	EpdgInd                     *bool              `json:"epdgInd,omitempty"`
	ContextInfo                 *ContextInfo       `json:"contextInfo,omitempty"`
	DataRestorationCallbackUri  string             `json:"dataRestorationCallbackUri,omitempty"`
	LastSynchronizationTime     string             `json:"lastSynchronizationTime,omitempty"`
	SmfInstanceId               string             `json:"smfInstanceId"`
	Dnn                         string             `json:"dnn,omitempty"`
	DeregCallbackUri            string             `json:"deregCallbackUri,omitempty"`
	RegistrationReason          RegistrationReason `json:"registrationReason,omitempty"`
	ResetIds                    []string           `json:"resetIds,omitempty"`
	SmfSetId                    string             `json:"smfSetId,omitempty"`
	SingleNssai                 Snssai             `json:"singleNssai"`
	PgwFqdn                     string             `json:"pgwFqdn,omitempty"`
	RegistrationTime            string             `json:"registrationTime,omitempty"`
	PcfId                       string             `json:"pcfId,omitempty"`
	PgwIpAddr                   *IpAddress         `json:"pgwIpAddr,omitempty"`
	UdrRestartInd               *bool              `json:"udrRestartInd,omitempty"`
	SupportedFeatures           string             `json:"supportedFeatures,omitempty"`
	PduSessionId                int                `json:"pduSessionId"`
	EmergencyServices           *bool              `json:"emergencyServices,omitempty"`
	PcscfRestorationCallbackUri string             `json:"pcscfRestorationCallbackUri,omitempty"`
	PlmnId                      PlmnId             `json:"plmnId"`
}
