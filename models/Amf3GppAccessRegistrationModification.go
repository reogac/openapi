package models

type Amf3GppAccessRegistrationModification struct {
	ImsVoPs             ImsVoPs              `json:"imsVoPs,omitempty"`
	BackupAmfInfo       []BackupAmfInfo      `json:"backupAmfInfo,omitempty"`
	EpsInterworkingInfo *EpsInterworkingInfo `json:"epsInterworkingInfo,omitempty"`
	UeSrvccCapability   *bool                `json:"ueSrvccCapability,omitempty"`
	UeMINTCapability    *bool                `json:"ueMINTCapability,omitempty"`
	Guami               Guami                `json:"guami"`
	PurgeFlag           *bool                `json:"purgeFlag,omitempty"`
	Pei                 string               `json:"pei,omitempty"`
}
