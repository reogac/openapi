package models

type AmfNon3GppAccessRegistrationModification struct {
	BackupAmfInfo []BackupAmfInfo `json:"backupAmfInfo,omitempty"`
	Guami         Guami           `json:"guami"`
	PurgeFlag     *bool           `json:"purgeFlag,omitempty"`
	Pei           string          `json:"pei,omitempty"`
	ImsVoPs       ImsVoPs         `json:"imsVoPs,omitempty"`
}
