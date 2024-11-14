package models
type BatteryIndication struct {
	 ReplaceableInd	*bool	`json:"replaceableInd,omitempty"`
	 RechargeableInd	*bool	`json:"rechargeableInd,omitempty"`
	 BatteryInd	*bool	`json:"batteryInd,omitempty"`
}
