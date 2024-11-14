package models
type AmfNon3GppAccessRegistrationModification struct {
	 Guami	Guami	`json:"guami"`
	 PurgeFlag	*bool	`json:"purgeFlag,omitempty"`
	 Pei	string	`json:"pei,omitempty"`
	 ImsVoPs	ImsVoPs	`json:"imsVoPs,omitempty"`
	 BackupAmfInfo	[]BackupAmfInfo	`json:"backupAmfInfo,omitempty"`
}
