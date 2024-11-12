package models

type BatteryIndication struct {
	RechargeableInd *bool `json:"rechargeableInd,omitempty"`
	BatteryInd      *bool `json:"batteryInd,omitempty"`
	ReplaceableInd  *bool `json:"replaceableInd,omitempty"`
}
