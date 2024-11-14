/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:39 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfNon3GppAccessRegistrationModification struct {
	ImsVoPs       ImsVoPs         `json:"imsVoPs,omitempty"`
	BackupAmfInfo []BackupAmfInfo `json:"backupAmfInfo,omitempty"`
	Guami         Guami           `json:"guami"`
	PurgeFlag     *bool           `json:"purgeFlag,omitempty"`
	Pei           string          `json:"pei,omitempty"`
}
