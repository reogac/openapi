package models
type AmfEventSubscriptionAddInfo struct {
	 AoiStateList	map[string]AreaOfInterestEventState	`json:"aoiStateList,omitempty"`
	 BindingInfo	[]string	`json:"bindingInfo,omitempty"`
	 SubscribingNfType	NFType	`json:"subscribingNfType,omitempty"`
	 EventSyncInd	*bool	`json:"eventSyncInd,omitempty"`
	 NfConsumerInfo	[]string	`json:"nfConsumerInfo,omitempty"`
}
