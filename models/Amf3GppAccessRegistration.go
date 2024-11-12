package models

type Amf3GppAccessRegistration struct {
	DeregCallbackUri            string               `json:"deregCallbackUri"`
	RatType                     RatType              `json:"ratType"`
	VgmlcAddress                *VgmlcAddress        `json:"vgmlcAddress,omitempty"`
	SupportedFeatures           string               `json:"supportedFeatures,omitempty"`
	PurgeFlag                   *bool                `json:"purgeFlag,omitempty"`
	Supi                        string               `json:"supi,omitempty"`
	UeReachableInd              UeReachableInd       `json:"ueReachableInd,omitempty"`
	PcscfRestorationCallbackUri string               `json:"pcscfRestorationCallbackUri,omitempty"`
	DrFlag                      *bool                `json:"drFlag,omitempty"`
	NoEeSubscriptionInd         *bool                `json:"noEeSubscriptionInd,omitempty"`
	UeMINTCapability            *bool                `json:"ueMINTCapability,omitempty"`
	ContextInfo                 *ContextInfo         `json:"contextInfo,omitempty"`
	AdminDeregSubWithdrawn      *bool                `json:"adminDeregSubWithdrawn,omitempty"`
	DisasterRoamingInd          *bool                `json:"disasterRoamingInd,omitempty"`
	Guami                       Guami                `json:"guami"`
	EpsInterworkingInfo         *EpsInterworkingInfo `json:"epsInterworkingInfo,omitempty"`
	AmfServiceNameDereg         ServiceName          `json:"amfServiceNameDereg,omitempty"`
	LastSynchronizationTime     string               `json:"lastSynchronizationTime,omitempty"`
	Pei                         string               `json:"pei,omitempty"`
	ImsVoPs                     ImsVoPs              `json:"imsVoPs,omitempty"`
	AmfEeSubscriptionId         string               `json:"amfEeSubscriptionId,omitempty"`
	ReRegistrationRequired      *bool                `json:"reRegistrationRequired,omitempty"`
	DataRestorationCallbackUri  string               `json:"dataRestorationCallbackUri,omitempty"`
	SorSnpnSiSupported          *bool                `json:"sorSnpnSiSupported,omitempty"`
	UdrRestartInd               *bool                `json:"udrRestartInd,omitempty"`
	BackupAmfInfo               []BackupAmfInfo      `json:"backupAmfInfo,omitempty"`
	UrrpIndicator               *bool                `json:"urrpIndicator,omitempty"`
	ResetIds                    []string             `json:"resetIds,omitempty"`
	InitialRegistrationInd      *bool                `json:"initialRegistrationInd,omitempty"`
	UeSrvccCapability           *bool                `json:"ueSrvccCapability,omitempty"`
	EmergencyRegistrationInd    *bool                `json:"emergencyRegistrationInd,omitempty"`
	RegistrationTime            string               `json:"registrationTime,omitempty"`
	AmfInstanceId               string               `json:"amfInstanceId"`
	AmfServiceNamePcscfRest     ServiceName          `json:"amfServiceNamePcscfRest,omitempty"`
}
