package models

type Amf3GppAccessRegistration struct {
	DeregCallbackUri            string               `json:"deregCallbackUri"`
	BackupAmfInfo               []BackupAmfInfo      `json:"backupAmfInfo,omitempty"`
	EpsInterworkingInfo         *EpsInterworkingInfo `json:"epsInterworkingInfo,omitempty"`
	DataRestorationCallbackUri  string               `json:"dataRestorationCallbackUri,omitempty"`
	SorSnpnSiSupported          *bool                `json:"sorSnpnSiSupported,omitempty"`
	SupportedFeatures           string               `json:"supportedFeatures,omitempty"`
	Pei                         string               `json:"pei,omitempty"`
	InitialRegistrationInd      *bool                `json:"initialRegistrationInd,omitempty"`
	RegistrationTime            string               `json:"registrationTime,omitempty"`
	AdminDeregSubWithdrawn      *bool                `json:"adminDeregSubWithdrawn,omitempty"`
	AmfInstanceId               string               `json:"amfInstanceId"`
	RatType                     RatType              `json:"ratType"`
	ContextInfo                 *ContextInfo         `json:"contextInfo,omitempty"`
	Supi                        string               `json:"supi,omitempty"`
	ReRegistrationRequired      *bool                `json:"reRegistrationRequired,omitempty"`
	ResetIds                    []string             `json:"resetIds,omitempty"`
	UeMINTCapability            *bool                `json:"ueMINTCapability,omitempty"`
	UdrRestartInd               *bool                `json:"udrRestartInd,omitempty"`
	EmergencyRegistrationInd    *bool                `json:"emergencyRegistrationInd,omitempty"`
	Guami                       Guami                `json:"guami"`
	DisasterRoamingInd          *bool                `json:"disasterRoamingInd,omitempty"`
	DrFlag                      *bool                `json:"drFlag,omitempty"`
	UrrpIndicator               *bool                `json:"urrpIndicator,omitempty"`
	AmfEeSubscriptionId         string               `json:"amfEeSubscriptionId,omitempty"`
	UeReachableInd              UeReachableInd       `json:"ueReachableInd,omitempty"`
	LastSynchronizationTime     string               `json:"lastSynchronizationTime,omitempty"`
	ImsVoPs                     ImsVoPs              `json:"imsVoPs,omitempty"`
	AmfServiceNamePcscfRest     ServiceName          `json:"amfServiceNamePcscfRest,omitempty"`
	UeSrvccCapability           *bool                `json:"ueSrvccCapability,omitempty"`
	NoEeSubscriptionInd         *bool                `json:"noEeSubscriptionInd,omitempty"`
	PurgeFlag                   *bool                `json:"purgeFlag,omitempty"`
	AmfServiceNameDereg         ServiceName          `json:"amfServiceNameDereg,omitempty"`
	PcscfRestorationCallbackUri string               `json:"pcscfRestorationCallbackUri,omitempty"`
	VgmlcAddress                *VgmlcAddress        `json:"vgmlcAddress,omitempty"`
}
