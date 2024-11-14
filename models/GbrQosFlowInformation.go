package models
type GbrQosFlowInformation struct {
	 GuaFbrDl	string	`json:"guaFbrDl"`
	 GuaFbrUl	string	`json:"guaFbrUl"`
	 NotifControl	NotificationControl	`json:"notifControl,omitempty"`
	 MaxPacketLossRateDl	*int	`json:"maxPacketLossRateDl,omitempty"`
	 MaxPacketLossRateUl	*int	`json:"maxPacketLossRateUl,omitempty"`
	 AlternativeQosProfileList	[]AlternativeQosProfile	`json:"alternativeQosProfileList,omitempty"`
	 MaxFbrDl	string	`json:"maxFbrDl"`
	 MaxFbrUl	string	`json:"maxFbrUl"`
}
