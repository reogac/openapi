package models

type Amf3GppAccessRegistration struct {
	UrrpIndicator               *bool                `json:"urrpIndicator,omitempty"`
	Supi                        string               `json:"supi,omitempty"`
	ReRegistrationRequired      *bool                `json:"reRegistrationRequired,omitempty"`
	AmfInstanceId               string               `json:"amfInstanceId"`
	SupportedFeatures           string               `json:"supportedFeatures,omitempty"`
	VgmlcAddress                *VgmlcAddress        `json:"vgmlcAddress,omitempty"`
	NoEeSubscriptionInd         *bool                `json:"noEeSubscriptionInd,omitempty"`
	UeMINTCapability            *bool                `json:"ueMINTCapability,omitempty"`
	EmergencyRegistrationInd    *bool                `json:"emergencyRegistrationInd,omitempty"`
	DisasterRoamingInd          *bool                `json:"disasterRoamingInd,omitempty"`
	Pei                         string               `json:"pei,omitempty"`
	ImsVoPs                     ImsVoPs              `json:"imsVoPs,omitempty"`
	DeregCallbackUri            string               `json:"deregCallbackUri"`
	Guami                       Guami                `json:"guami"`
	EpsInterworkingInfo         *EpsInterworkingInfo `json:"epsInterworkingInfo,omitempty"`
	AdminDeregSubWithdrawn      *bool                `json:"adminDeregSubWithdrawn,omitempty"`
	ResetIds                    []string             `json:"resetIds,omitempty"`
	AmfServiceNameDereg         ServiceName          `json:"amfServiceNameDereg,omitempty"`
	PcscfRestorationCallbackUri string               `json:"pcscfRestorationCallbackUri,omitempty"`
	BackupAmfInfo               []BackupAmfInfo      `json:"backupAmfInfo,omitempty"`
	DrFlag                      *bool                `json:"drFlag,omitempty"`
	PurgeFlag                   *bool                `json:"purgeFlag,omitempty"`
	RatType                     RatType              `json:"ratType"`
	SorSnpnSiSupported          *bool                `json:"sorSnpnSiSupported,omitempty"`
	ContextInfo                 *ContextInfo         `json:"contextInfo,omitempty"`
	UeReachableInd              UeReachableInd       `json:"ueReachableInd,omitempty"`
	AmfServiceNamePcscfRest     ServiceName          `json:"amfServiceNamePcscfRest,omitempty"`
	InitialRegistrationInd      *bool                `json:"initialRegistrationInd,omitempty"`
	AmfEeSubscriptionId         string               `json:"amfEeSubscriptionId,omitempty"`
	UeSrvccCapability           *bool                `json:"ueSrvccCapability,omitempty"`
	RegistrationTime            string               `json:"registrationTime,omitempty"`
	DataRestorationCallbackUri  string               `json:"dataRestorationCallbackUri,omitempty"`
	UdrRestartInd               *bool                `json:"udrRestartInd,omitempty"`
	LastSynchronizationTime     string               `json:"lastSynchronizationTime,omitempty"`
}
