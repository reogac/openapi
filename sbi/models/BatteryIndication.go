/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type BatteryIndication struct {
	RechargeableInd *bool `json:"rechargeableInd,omitempty"`
	BatteryInd      *bool `json:"batteryInd,omitempty"`
	ReplaceableInd  *bool `json:"replaceableInd,omitempty"`
}
