/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:10 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfNon3GppAccessRegistrationModification struct {
	BackupAmfInfo []BackupAmfInfo `json:"backupAmfInfo,omitempty"`
	Guami         Guami           `json:"guami"`
	PurgeFlag     *bool           `json:"purgeFlag,omitempty"`
	Pei           string          `json:"pei,omitempty"`
	ImsVoPs       ImsVoPs         `json:"imsVoPs,omitempty"`
}
