package models

type Amf3GppAccessRegistration struct {
	DrFlag                      *bool                `json:"drFlag,omitempty"`
	UrrpIndicator               *bool                `json:"urrpIndicator,omitempty"`
	LastSynchronizationTime     string               `json:"lastSynchronizationTime,omitempty"`
	SupportedFeatures           string               `json:"supportedFeatures,omitempty"`
	ImsVoPs                     ImsVoPs              `json:"imsVoPs,omitempty"`
	ReRegistrationRequired      *bool                `json:"reRegistrationRequired,omitempty"`
	AmfEeSubscriptionId         string               `json:"amfEeSubscriptionId,omitempty"`
	NoEeSubscriptionInd         *bool                `json:"noEeSubscriptionInd,omitempty"`
	EpsInterworkingInfo         *EpsInterworkingInfo `json:"epsInterworkingInfo,omitempty"`
	RegistrationTime            string               `json:"registrationTime,omitempty"`
	ResetIds                    []string             `json:"resetIds,omitempty"`
	SorSnpnSiSupported          *bool                `json:"sorSnpnSiSupported,omitempty"`
	UdrRestartInd               *bool                `json:"udrRestartInd,omitempty"`
	PcscfRestorationCallbackUri string               `json:"pcscfRestorationCallbackUri,omitempty"`
	Guami                       Guami                `json:"guami"`
	EmergencyRegistrationInd    *bool                `json:"emergencyRegistrationInd,omitempty"`
	UeReachableInd              UeReachableInd       `json:"ueReachableInd,omitempty"`
	UeMINTCapability            *bool                `json:"ueMINTCapability,omitempty"`
	PurgeFlag                   *bool                `json:"purgeFlag,omitempty"`
	AmfServiceNameDereg         ServiceName          `json:"amfServiceNameDereg,omitempty"`
	Pei                         string               `json:"pei,omitempty"`
	AmfServiceNamePcscfRest     ServiceName          `json:"amfServiceNamePcscfRest,omitempty"`
	VgmlcAddress                *VgmlcAddress        `json:"vgmlcAddress,omitempty"`
	AmfInstanceId               string               `json:"amfInstanceId"`
	DeregCallbackUri            string               `json:"deregCallbackUri"`
	ContextInfo                 *ContextInfo         `json:"contextInfo,omitempty"`
	AdminDeregSubWithdrawn      *bool                `json:"adminDeregSubWithdrawn,omitempty"`
	InitialRegistrationInd      *bool                `json:"initialRegistrationInd,omitempty"`
	BackupAmfInfo               []BackupAmfInfo      `json:"backupAmfInfo,omitempty"`
	Supi                        string               `json:"supi,omitempty"`
	DataRestorationCallbackUri  string               `json:"dataRestorationCallbackUri,omitempty"`
	DisasterRoamingInd          *bool                `json:"disasterRoamingInd,omitempty"`
	RatType                     RatType              `json:"ratType"`
	UeSrvccCapability           *bool                `json:"ueSrvccCapability,omitempty"`
}
