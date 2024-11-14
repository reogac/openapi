/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:44 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type BatteryIndication struct {
	BatteryInd      *bool `json:"batteryInd,omitempty"`
	ReplaceableInd  *bool `json:"replaceableInd,omitempty"`
	RechargeableInd *bool `json:"rechargeableInd,omitempty"`
}
