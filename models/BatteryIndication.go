/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:23:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type BatteryIndication struct {
	ReplaceableInd  *bool `json:"replaceableInd,omitempty"`
	RechargeableInd *bool `json:"rechargeableInd,omitempty"`
	BatteryInd      *bool `json:"batteryInd,omitempty"`
}
